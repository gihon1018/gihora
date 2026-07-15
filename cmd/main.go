package main

import (
	"gihora/api/router"
	"log"
)

func main() {
	r := router.InitRouter()

	log.Println("Server starting on 8080...")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
