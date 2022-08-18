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
	v1.GET("/community", Controller.GetCommunity) // 获取分类社区列表
	v1.GET("/refreshToken", Controller.RefreshTokenHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/user", Controller.ListUser)
		v1.POST("/post", Controller.PostArticle) // 发布文章
	}

}
