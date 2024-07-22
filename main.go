package main

import (
	"gin-vue-admin/api"
	"gin-vue-admin/utils"
)

func main() {
	// 加载配置
	utils.LoadConfig()
	// 连接数据库
	utils.InitDb()

	route := api.NewRoute()

	route.Run(utils.GetEnv("ADDRESS"))
}
