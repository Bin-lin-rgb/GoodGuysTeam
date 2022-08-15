package Controller

import (
	"go.uber.org/zap"
	"log"
	"net/http"
	"server/Dao"
	"server/Model"
	"server/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginResponse struct {
	response.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	zap.L().Debug("----register----",
		zap.String("username", username), 
		zap.String("password", password),
	)

	token, _ := GenToken(username, password)

	// 查找用户是否存在
	user, err := Dao.Mgr.IsExist(username)
	if err != nil {
		//log.Println(err)
		zap.L().Error("Dao.mgr.IsExist(username) Failed", zap.Error(err))
	}

	if user.Username != "" {
		log.Println("User already exist")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: response.Response{Code: 1, Message: "User already exist"},
		})
		return
	}

	// encrypted : 已加密的密码
	encrypted, _ := GetPwd(password)
	userinfo := Model.User{
		Username: username,
		Password: string(encrypted),
	}
	// 将加密的密码写入数据库
	err = Dao.Mgr.Register(userinfo)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: response.Response{Code: 1, Message: "请更换一个账户名称。"},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: response.Response{Code: 0},
		UserId:   userinfo.Id,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println(username, "----**login**----", password)

	token, _ := GenToken(username, password)

	// 查找用户是否存在
	user, err := Dao.Mgr.IsExist(username)
	if err != nil {
		log.Println(err)
	}

	if user.Username == "" {
		log.Println("User doesn't exist")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: response.Response{Code: 1, Message: "User doesn't exist"},
		})
		return
	}

	if ComparePwd(user.Password, password) {
		log.Println("Login successful!")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: response.Response{Code: 0, Message: "Login successful!"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: response.Response{Code: 1, Message: "Password is not correct."},
		})
	}
}

func ListUser(c *gin.Context) {
	users, err := Dao.Mgr.GetUser()
	if err != nil {
		response.Failed("返回失败！", c)
	}
	response.Success("数据返回成功！", users, c)
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
