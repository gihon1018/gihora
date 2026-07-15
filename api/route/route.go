package route

import (
	v1 "gihora/api/route/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.SetTrustedProxies(nil)

	r.Use(gin.Logger(), gin.Recovery())

	v1.RegisterV1Routes(r)

	return r
}
