package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	HTTPStatus int
	Code       int
	Msg        string
	cause      error
}

func New(httpStatus, code int, msg string) *AppError {
	return &AppError{
		HTTPStatus: httpStatus,
		Code:       code,
		Msg:        msg,
	}
}

var (
	ModParamInvalid     = New(http.StatusBadRequest, 400101, "模组参数无效")
	ModAlreadySubbed    = New(http.StatusConflict, 400102, "模组已订阅")
	ModSubFail          = New(http.StatusInternalServerError, 400103, "模组订阅失败")
	ModNotSubbed        = New(http.StatusNotFound, 400104, "模组未订阅")
	ModUnsubFail        = New(http.StatusInternalServerError, 400105, "模组取消订阅失败")
	ModUpdateRemarkFail = New(http.StatusInternalServerError, 400106, "模组更新备注失败")
)

func (e *AppError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Msg, e.cause)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

func (e *AppError) WithMsg(msg string) *AppError {
	return &AppError{
		HTTPStatus: e.HTTPStatus,
		Code:       e.Code,
		Msg:        msg,
		cause:      e.cause,
	}
}

func (e *AppError) WithCause(err error) *AppError {
	return &AppError{
		HTTPStatus: e.HTTPStatus,
		Code:       e.Code,
		Msg:        e.Msg,
		cause:      err,
	}
}

func (e *AppError) Unwrap() error {
	return e.cause
}

func (e *AppError) Is(target error) bool {
	var t *AppError
	ok := errors.As(target, &t)
	return ok && e.Code == t.Code
}

func IsAppError(err error) (*AppError, bool) {
	if ae, ok := errors.AsType[*AppError](err); ok {
		return ae, true
	}
	return nil, false
}
