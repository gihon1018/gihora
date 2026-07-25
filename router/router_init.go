package router

import (
	"fmt"
	"gihora/internal/dst"

	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("设置信任代理失败: %w", err)
	}

	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	{
		if err := dst.RegisterRoutes(api); err != nil {
			return nil, fmt.Errorf("注册模组路由失败: %w", err)
		}
	}

	return r, nil
}
