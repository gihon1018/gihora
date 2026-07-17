package dst

import (
	"gihora/internal/dst/cluster"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	dstGroup := r.Group("/dst")

	cluster.RegisterRoutes(dstGroup)
}
