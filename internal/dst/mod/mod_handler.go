package mod

import (
	"fmt"

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
	id := c.Param("id")

	var mod Mod
	if err := c.ShouldBindJSON(&mod); err != nil {
		resp.Fail(c, apperr.ModParamInvalid.Code, fmt.Sprintf("模组 %s 参数无效", id))
		return
	}

	if err := h.service.SubMod(id, mod.Remark); err != nil {
		resp.Fail(c, apperr.ModSubFail.Code, fmt.Sprintf("模组 %s 订阅失败", id))
		return
	}

	resp.Success(c, fmt.Sprintf("模组 %s 订阅成功", id))
}

func (h *Handler) UnsubMod(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.UnsubMod(id); err != nil {
		resp.Fail(c, apperr.ModUnsubFail.Code, fmt.Sprintf("模组 %s 取消订阅失败", id))
		return
	}

	resp.Success(c, fmt.Sprintf("模组 %s 取消订阅成功", id))
}

func (h *Handler) UpdateModRemark(c *gin.Context) {
	id := c.Param("id")

	var mod Mod
	if err := c.ShouldBindJSON(&mod); err != nil {
		resp.Fail(c, apperr.ModParamInvalid.Code, fmt.Sprintf("模组 %s 参数无效", id))
		return
	}

	if err := h.service.UpdateModRemark(id, mod.Remark); err != nil {
		resp.Fail(c, apperr.ModSubFail.Code, fmt.Sprintf("模组 %s 更新备注失败", id))
		return
	}

	resp.Success(c, fmt.Sprintf("模组 %s 更新备注成功", id))
}
