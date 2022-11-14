package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/Logger"
	"server/Router"
	"server/Utils"
	"server/pkg/snowflake"
)

func main() {
	r := gin.Default()

	//// load config from config.json
	//if len(os.Args) < 1 {
	//	return
	//}
	//
	//if err := Config.Init(os.Args[1]); err != nil {
	//	panic(err)
	//}

	// init logger
	var logconfig Logger.LogConfig
	logconfig.Level = Utils.Level
	logconfig.Filename = Utils.Filename
	logconfig.MaxSize = Utils.MaxSize
	logconfig.MaxAge = Utils.MaxAge
	logconfig.MaxBackups = Utils.MaxBackups

	if err := Logger.InitLogger(&logconfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 雪花算法生成分布式ID
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	gin.SetMode(Utils.AppMode)

	Router.Start(r)
	// 注册zap相关中间件
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))

	err := r.Run(Utils.HttpPort)
	if err != nil {
		log.Println("r.Run() Failed!")
	}
}
