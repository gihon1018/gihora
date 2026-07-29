package mod

import (
	"fmt"
	"gihora/internal/dst/mod/model"

	"github.com/gin-gonic/gin"

	"gihora/pkg/apperr"
	"gihora/pkg/resp"
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

func (h *Handler) ListMods(c *gin.Context) {
	mods := h.service.ListMods()
	resp.Success(c, mods)
}

func (h *Handler) SubMod(c *gin.Context) {
	var mod model.Mod
	if err := c.ShouldBindJSON(&mod); err != nil {
		resp.Fail(c, apperr.ModParamInvalid)
		return
	}

	if err := h.service.SubMod(mod.Id, mod.Remark); err != nil {
		resp.Fail(c, err)
		return
	}

	resp.Success(c, fmt.Sprintf("模组 %s 订阅成功", mod.Id))
}

func (h *Handler) UnsubMod(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.UnsubMod(id); err != nil {
		resp.Fail(c, err)
		return
	}

	resp.Success(c, fmt.Sprintf("模组 %s 取消订阅成功", id))
}

func (h *Handler) UpdateModRemark(c *gin.Context) {
	id := c.Param("id")

	var mod model.Mod
	if err := c.ShouldBindJSON(&mod); err != nil {
		resp.Fail(c, apperr.ModParamInvalid)
		return
	}

	if err := h.service.UpdateModRemark(id, mod.Remark); err != nil {
		resp.Fail(c, err)
		return
	}

	resp.Success(c, fmt.Sprintf("模组 %s 更新备注成功", id))
}
