package mod

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	modGroup := group.Group("/mods")
	{
		modGroup.GET("/", ListMods)
		modGroup.POST("/:modId", SubMod)
		modGroup.DELETE("/:modId", UnsubMod)
		modGroup.PUT("/:modId", UpdateModRemark)
	}
}
