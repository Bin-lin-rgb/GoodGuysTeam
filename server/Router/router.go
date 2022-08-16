package Router

import (
	"server/Controller"
	middlewares "server/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	r.Static("/static", "./static") // 加载静态文件

	mwCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			return true
		},

		MaxAge: 24 * time.Hour,
	})

	r.Use(mwCORS)

	v1 := r.Group("/v1")
	v1.POST("/user/login", Controller.Login)
	v1.POST("/user/register", Controller.Register)
	//r.GET("/user", Controller.ListUser)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		r.GET("/user", Controller.ListUser)
		v1.GET("/refresh_token", Controller.RefreshTokenHandler)
	}

}
