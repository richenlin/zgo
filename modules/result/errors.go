package result

import (
	"zgo/engine"
	"zgo/modules/language"
)

// H h -> map
type H map[string]interface{}

// ErrorModel 异常模型
type ErrorModel struct {
	status       int
	showType     int
	errorCode    string
	errorMessage string
}

// 定义错误
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
var (
	Err400BadRequest       = ErrorModel{status: 400, showType: ShowWarn, errorCode: "ERR-BAD-REQUEST", errorMessage: "请求发生错误"}
	Err401Unauthorized     = ErrorModel{status: 401, showType: ShowWarn, errorCode: "ERR-UNAUTHORIZED", errorMessage: "用户没有权限（令牌、用户名、密码错误）"}
	Err403Forbidden        = ErrorModel{status: 403, showType: ShowWarn, errorCode: "ERR-FORBIDDEN", errorMessage: "用户得到授权，但是访问是被禁止的"}
	Err404NotFound         = ErrorModel{status: 404, showType: ShowWarn, errorCode: "ERR-NOT-FOUND", errorMessage: "发出的请求针对的是不存在的记录，服务器没有进行操作"}
	Err405MethodNotAllowed = ErrorModel{status: 405, showType: ShowWarn, errorCode: "ERR-METHOD-NOT-ALLOWED", errorMessage: "请求的方法不允许"}
	Err406NotAcceptable    = ErrorModel{status: 406, showType: ShowWarn, errorCode: "ERR-NOT-ACCEPTABLE", errorMessage: "请求的格式不可得"}
	Err429TooManyRequests  = ErrorModel{status: 429, showType: ShowWarn, errorCode: "ERR-TOO-MANY-REQUESTS", errorMessage: "请求次数过多"}
	Err500InternalServer   = ErrorModel{status: 500, showType: ShowWarn, errorCode: "ERR-INTERNAL-SERVER", errorMessage: "服务器发生错误"}
)

// NewError 包装响应错误
func NewError(ctx engine.Context, showType int, code string, msg string, args ...interface{}) *ErrorInfo {
	res := &ErrorInfo{
		success:      false,
		errorCode:    code,
		errorMessage: language.Sprintf(ctx, code, msg, args...),
		showType:     showType,
		traceID:      ctx.GetTraceID(),
	}
	return res
}

// NewSuccess 包装响应结果
func NewSuccess(ctx engine.Context, data interface{}) *Result {
	res := &Result{
		success: true,
		data:    data,
		traceID: ctx.GetTraceID(),
	}
	return res
}

// Wrap400Response 无法解析异常
func Wrap400Response(ctx engine.Context, err error) *ErrorModel {
	return &ErrorModel{
		status:       400,
		showType:     ShowWarn,
		errorCode:    "ERR-BAD-REQUEST-X",
		errorMessage: language.Sprintf(ctx, "ERR-BAD-REQUEST-X", "解析请求参数发生错误 - %s", err.Error()),
	}
}
