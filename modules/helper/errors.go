package helper

import (
	"github.com/suisrc/zgo/modules/language"

	"github.com/gin-gonic/gin"
)

// H h -> map
type H map[string]interface{}

// ErrorModel 异常模型
type ErrorModel struct {
	Status       int
	ShowType     int
	ErrorCode    string
	ErrorMessage string
}

// 定义错误
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
var (
	Err400BadRequest       = ErrorModel{Status: 400, ShowType: ShowWarn, ErrorCode: "ERR-BAD-REQUEST", ErrorMessage: "请求发生错误"}
	Err401Unauthorized     = ErrorModel{Status: 401, ShowType: ShowWarn, ErrorCode: "ERR-UNAUTHORIZED", ErrorMessage: "用户没有权限（令牌、用户名、密码错误）"}
	Err403Forbidden        = ErrorModel{Status: 403, ShowType: ShowWarn, ErrorCode: "ERR-FORBIDDEN", ErrorMessage: "用户得到授权，但是访问是被禁止的"}
	Err404NotFound         = ErrorModel{Status: 404, ShowType: ShowWarn, ErrorCode: "ERR-NOT-FOUND", ErrorMessage: "发出的请求针对的是不存在的记录，服务器没有进行操作"}
	Err405MethodNotAllowed = ErrorModel{Status: 405, ShowType: ShowWarn, ErrorCode: "ERR-METHOD-NOT-ALLOWED", ErrorMessage: "请求的方法不允许"}
	Err406NotAcceptable    = ErrorModel{Status: 406, ShowType: ShowWarn, ErrorCode: "ERR-NOT-ACCEPTABLE", ErrorMessage: "请求的格式不可得"}
	Err429TooManyRequests  = ErrorModel{Status: 429, ShowType: ShowWarn, ErrorCode: "ERR-TOO-MANY-REQUESTS", ErrorMessage: "请求次数过多"}
	Err500InternalServer   = ErrorModel{Status: 500, ShowType: ShowWarn, ErrorCode: "ERR-INTERNAL-SERVER", ErrorMessage: "服务器发生错误"}
)

// NewError 包装响应错误
func NewError(ctx *gin.Context, showType int, code string, msg string, args ...interface{}) *ErrorInfo {
	res := &ErrorInfo{
		Success:      false,
		ErrorCode:    code,
		ErrorMessage: language.Sprintf(ctx, code, msg, args...),
		ShowType:     showType,
		TraceID:      GetTraceID(ctx),
	}
	return res
}

// NewOK 包装响应结果
func NewOK(ctx *gin.Context, data interface{}) *Success {
	res := &Success{
		Success: true,
		Data:    data,
		TraceID: GetTraceID(ctx),
	}
	return res
}

// Wrap400Response 无法解析异常
func Wrap400Response(ctx *gin.Context, err error) *ErrorModel {
	return &ErrorModel{
		Status:       400,
		ShowType:     ShowWarn,
		ErrorCode:    "ERR-BAD-REQUEST-X",
		ErrorMessage: language.Sprintf(ctx, "ERR-BAD-REQUEST-X", "解析请求参数发生错误 - %s", err.Error()),
	}
}
