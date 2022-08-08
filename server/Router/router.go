package Router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"server/Controller"
	"time"
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
	e.POST("/add", Controller.AddUser)
	e.GET("/list", Controller.ListUser)
	err := e.Run(":9090")
	if err != nil {
		log.Println("e.Run() Failed!")
	}
}
