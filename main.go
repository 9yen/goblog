package main

import (
	"goblog/bootstrap"
	btsConig "goblog/config"
	"goblog/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConig.Initialize()
}
func main() {
	config.InitConfig("local")
	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
}
