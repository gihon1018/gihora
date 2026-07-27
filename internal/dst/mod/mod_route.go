package mod

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) error {
	manager, err := NewManager()
	if err != nil {
		return err
	}
	service := NewService(manager)
	handler := NewHandler(service)

	modGroup := group.Group("/mods")
	{
		modGroup.GET("/", handler.ListMods)
		modGroup.POST("/", handler.SubMod)
		modGroup.DELETE("/:id", handler.UnsubMod)
		modGroup.PATCH("/:id", handler.UpdateModRemark)
	}

	return nil
}
