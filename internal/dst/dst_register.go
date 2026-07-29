package dst

import (
	"gihora/internal/dst/cluster"
	"gihora/internal/dst/mod"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) error {
	dstGroup := r.Group("/dst")

	if err := mod.RegisterRoutes(dstGroup); err != nil {
		log.Printf("[ERROR] 饥荒联机版模组路由注册失败\n")
		return err
	}

	if err := cluster.RegisterRoutes(dstGroup); err != nil {
		log.Printf("[ERROR] 饥荒联机版集群路由注册失败\n")
		return err
	}

	return nil
}
