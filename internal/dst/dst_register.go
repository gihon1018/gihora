package dst

import (
	"gihora/internal/dst/mod"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	dstGroup := r.Group("/dst")

	mod.RegisterRoutes(dstGroup)
}
