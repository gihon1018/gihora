package main

import (
	"gihora/router"
	"log"
)

func main() {
	r := router.InitRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
