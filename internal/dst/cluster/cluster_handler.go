package cluster

import (
	"gihora/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	h := &Handler{
		service: service,
	}
	return h
}

func (h *Handler) ListClusters(c *gin.Context) {
	clusters := h.service.ListClusters()
	resp.Success(c, clusters)
}
