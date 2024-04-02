package main

import (
	"backend/router"
)

func main() {
	// 注册路由
	r := router.SetupRouter()
	// 启动服务
	r.Run(":8080")
}
