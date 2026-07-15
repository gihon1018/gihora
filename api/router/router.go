package router

import (
	"gihora/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	appGroup := r.Group("/api/v1")
	{
		appGroup.GET("/ping", handler.PingHandler)
	}

	return r
}
