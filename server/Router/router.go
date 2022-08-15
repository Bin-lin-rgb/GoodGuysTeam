package Router

import (
	"server/Controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {

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
	r.POST("/user/login", Controller.Login)
	r.POST("/user/register", Controller.Register)
	r.GET("/user", Controller.ListUser)

}
