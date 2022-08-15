package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"server/Config"
	"server/Logger"
	"server/Router"
)

func main() {
	r := gin.Default()

	// load config from config.json
	if len(os.Args) < 1 {
		return
	}

	if err := Config.Init(os.Args[1]); err != nil {
		panic(err)
	}

	// init logger
	if err := Logger.InitLogger(Config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	gin.SetMode(Config.Conf.Mode)

	Router.Start(r)
	// 注册zap相关中间件
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))

	//r.GET("/hello", func(c *gin.Context) {
	//	// 假设有一些数据需要记录到日志中
	//	var (
	//		name = "hello"
	//		age  = 20
	//	)
	//	// 记录日志并使用zap.Xxx(key, val)记录相关字段
	//	zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))
	//
	//	c.String(http.StatusOK, "hello world.com")
	//})

	err := r.Run()
	if err != nil {
		log.Println("r.Run() Failed!")
	}
}
