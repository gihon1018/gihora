package main

import (
	"gihora/api/route"
	"log"
)

func main() {
	r := route.InitRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
