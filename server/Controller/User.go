package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"server/Dao"
	"server/Model"
	"server/pkg/jwt"
	"server/pkg/snowflake"
	"strings"
)

type UserLoginResponse struct {
	Response
	UserId       uint64 `json:"user_id,omitempty"`
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	id, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
	}

	zap.L().Debug("----register----",
		zap.String("username", username),
		zap.String("password", password),
		zap.Uint64("id", id),
	)

	// 查找用户是否存在
	user, err := Dao.Mgr.IsExist(username)
	if err != nil {
		zap.L().Error("Dao.mgr.IsExist(username) Failed", zap.Error(err))
	}

	if user.Username != "" {
		ResponseError(c, CodeUserExist)
		return
	}

	// encrypted : 已加密的密码
	encrypted, _ := GetPwd(password)
	userinfo := Model.User{
		UserId:   id,
		Username: username,
		Password: string(encrypted),
	}
	// 将加密的密码写入数据库
	err = Dao.Mgr.Register(userinfo)
	if err != nil {
		ResponseError(c, CodeSQLError)
		return
	}

	ResponseSuccess(c, gin.H{
		"user_id":   id,
		"user_name": username,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 查找用户是否存在
	user, err := Dao.Mgr.IsExist(username)
	if err != nil {
		zap.L().Error("Dao.Mgr.IsExist(username) failed", zap.String("username", username), zap.Error(err))
	}

	AccessToken, RefreshToken, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		zap.L().Error("GenToken failed", zap.String("username", username), zap.Error(err))
	}

	if user.Username == "" {
		ResponseError(c, CodeUserNotExist)
		return
	}

	if ComparePwd(user.Password, password) {
		ResponseSuccess(c, UserLoginResponse{
			UserId:       user.UserId,
			Username:     username,
			AccessToken:  AccessToken,
			RefreshToken: RefreshToken,
		})
	} else {
		ResponseError(c, CodeInvalidPassword)
	}
}

func ListUser(c *gin.Context) {
	users, err := Dao.Mgr.GetUser()
	if err != nil {
		Failed("返回失败！", c)
	}
	Success("数据返回成功！", users, c)
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		zap.L().Error("RefreshToken Failed", zap.Error(err))
	}
	id, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID Failed", zap.Error(err))
	}
	log.Println("--------", id)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
