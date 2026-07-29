package cluster

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) error {
	manager, err := NewManager()
	if err != nil {
		return err
	}

	service := NewService(manager)
	handler := NewHandler(service)

	clusterGroup := group.Group("/clusters")
	{
		clusterGroup.GET("/", handler.ListClusters)
	}

	return nil
}
