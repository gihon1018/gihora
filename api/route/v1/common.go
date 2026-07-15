package v1

import (
	"gihora/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCommonRoutes(group *gin.RouterGroup) {
	common := group.Group("/common")
	{
		common.GET("/ping", handler.PingHandler)
	}
}
