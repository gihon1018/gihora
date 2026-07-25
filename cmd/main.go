package main

import (
	"gihora/router"
	"log"
)

func main() {
	r, err := router.InitRouter()

	if err != nil {
		log.Fatalf("[FATAL] 路由初始化失败: %v", err)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("[FATAL] 服务启动失败: %v", err)
	}
}
