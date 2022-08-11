package Controller

import (
	"errors"
	"log"
	"server/Dao"
	"server/Model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MySecret 定义Secret
var MySecret = []byte("GoodGuysTeam")

// GenToken 生成JWT
func GenToken(username string, password string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间
			Issuer:    "GoodGuysTeam",                        // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return MySecret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func TokenIsValid(token string) (Model.User, bool) {
	var userinfo Model.User

	// 是否过期
	claims, err := ParseToken(token)
	if err != nil {
		return userinfo, false
	}

	user, err := Dao.Mgr.IsExist(claims.Username)
	if err != nil {
		log.Println(err)
		return userinfo, false
	}

	if user.Username != claims.Username {
		return userinfo, false
	}

	return user, true
}
