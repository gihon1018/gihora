package apperr

type AppError struct {
	Code int
	Msg  string
	Err  error
}

func New(code int, msg string) *AppError {
	return &AppError{
		Code: code,
		Msg:  msg,
	}
}

var (
	ModParamInvalid = New(400101, "模组参数无效")
	ModSubFail      = New(400102, "模组订阅失败")
	ModUnsubFail    = New(400103, "模组取消订阅失败")
)
