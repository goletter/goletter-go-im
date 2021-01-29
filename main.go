package main

import (
	"fmt"
	"goletter-go-im/bootstrap"
	"goletter-go-im/config"
	conf "goletter-go-im/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {

	// 定时任务
	// commands.CleanInit()

	app := bootstrap.Start()
	addr := fmt.Sprintf(":%s", conf.GetString("app.port"))
	app.Run(addr)
}