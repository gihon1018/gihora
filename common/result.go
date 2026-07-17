package common

import "net/http"

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) R {
	return R{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
}

func SuccessMsg(msg string, data any) R {
	return R{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	}
}

func Fail(code int, msg string) R {
	return R{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
