package dst

import (
	"gihora/internal/dst/cluster"
	"gihora/internal/dst/mod"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	dstGroup := r.Group("/dst")

	cluster.RegisterRoutes(dstGroup)
	mod.RegisterRoutes(dstGroup)
}
