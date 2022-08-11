package Router

import (
	"log"
	"server/Controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

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

	e.Use(mwCORS)
	e.POST("/user/login", Controller.Login)
	e.POST("/user/register", Controller.Register)
	e.GET("/user", Controller.ListUser)
	err := e.Run()
	if err != nil {
		log.Println("e.Run() Failed!")
	}
}
