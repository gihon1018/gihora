package router

import (
	"gihora/internal/dst"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	if err := r.SetTrustedProxies(nil); err != nil {
		return nil
	}

	r.Use(gin.Logger(), gin.Recovery())

	dst.RegisterRoutes(r)

	return r
}
