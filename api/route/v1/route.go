package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(r *gin.Engine) {
	v1Group := r.Group("/api/v1")

	RegisterCommonRoutes(v1Group)
}
