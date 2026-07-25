package dst

import (
	"gihora/internal/dst/mod"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) error {
	dstGroup := r.Group("/dst")

	if err := mod.RegisterRoutes(dstGroup); err != nil {
		log.Printf("[ERROR] 饥荒联机版路由注册失败\n")
		return err
	}

	return nil
}
