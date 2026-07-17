package cluster

import (
	"gihora/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListClusters(c *gin.Context) {
	c.JSON(http.StatusOK, common.Success([]Cluster{
		{
			ClusterId:   100,
			ClusterName: "Dreamisland",
		},
		{
			ClusterId:   101,
			ClusterName: "Destiny",
		},
	}))
}
