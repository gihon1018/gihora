package cluster

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	clusterGroup := group.Group("/clusters")
	{
		clusterGroup.GET("/", ListClusters)
	}
}
