package resp

import (
	"gihora/pkg/apperr"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

func Fail(c *gin.Context, err error) {
	ae, ok := apperr.IsAppError(err)
	if ok {
		c.JSON(ae.HTTPStatus, Resp{
			Code: ae.Code,
			Msg:  ae.Msg,
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, Resp{
		Code: 500000,
		Msg:  "服务器内部异常",
		Data: nil,
	})
}

func FailMsg(c *gin.Context, httpStatus, code int, msg string) {
	c.JSON(httpStatus, Resp{
		Code: code,
		Msg:  msg,
	})
}
