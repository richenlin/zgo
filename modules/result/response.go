package result

import (
	"net/http"
	"zgo/engine"
	"zgo/modules/language"
)

// H h -> map
type H map[string]interface{}

// ErrorModel 异常模型
type ErrorModel struct {
	showType     int
	errorCode    string
	errorMessage string
}

// NewError 包装响应错误
func NewError(ctx engine.Context, showType int, code string, msg string, args ...interface{}) error {
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
func NewSuccess(ctx engine.Context, data interface{}) error {
	res := &Result{
		success: true,
		data:    data,
		traceID: ctx.GetTraceID(),
	}
	return res
}

// ResError 包装响应错误
func ResError(ctx engine.Context, em ErrorModel) error {
	res := &ErrorInfo{
		success:      false,
		errorCode:    em.errorCode,
		errorMessage: language.Sprintf(ctx, em.errorCode, em.errorMessage),
		showType:     em.showType,
		traceID:      ctx.GetTraceID(),
	}
	//ctx.JSON(http.StatusOK, res)
	//ctx.Abort()
	ResJSON(ctx, res)
	return res
}

// ResErrorResBody 包装响应错误
func ResErrorResBody(ctx engine.Context, em ErrorModel) error {
	res := &ErrorInfo{
		success:      false,
		errorCode:    em.errorCode,
		errorMessage: language.Sprintf(ctx, em.errorCode, em.errorMessage),
		showType:     em.showType,
		traceID:      ctx.GetTraceID(),
	}
	ResJSONResBody(ctx, res)
	return res
}

// ResJSONResBody 响应JSON数据
func ResJSONResBody(ctx engine.Context, v interface{}) {
	buf, err := JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	ctx.Set(ResBodyKey, buf)
	ctx.Data(http.StatusOK, ResponseTypeJSON, buf)
	ctx.Abort()
}

// ResJSON 响应JSON数据
func ResJSON(ctx engine.Context, v interface{}) {
	buf, err := JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	ctx.Data(http.StatusOK, ResponseTypeJSON, buf)
	ctx.Abort()
}
