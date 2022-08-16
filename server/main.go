package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"server/Config"
	"server/Logger"
	"server/Router"
	"server/pkg/snowflake"
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

	// 雪花算法生成分布式ID
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	gin.SetMode(Config.Conf.Mode)

	Router.Start(r)
	// 注册zap相关中间件
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))

	err := r.Run()
	if err != nil {
		log.Println("r.Run() Failed!")
	}
}
