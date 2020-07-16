package result

import (
	"net/http"
	"zgo/engine"
	"zgo/modules/language"
)

// ResError 包装响应错误
func ResError(ctx engine.Context, em *ErrorModel) error {
	res := &ErrorInfo{
		success:      false,
		errorCode:    em.errorCode,
		errorMessage: language.Sprintf(ctx, em.errorCode, em.errorMessage),
		showType:     em.showType,
		traceID:      ctx.GetTraceID(),
	}
	//ctx.JSON(http.StatusOK, res)
	//ctx.Abort()
	ResJSON(ctx, em.status, res)
	return res
}

// ResErrorResBody 包装响应错误
func ResErrorResBody(ctx engine.Context, em *ErrorModel) error {
	res := &ErrorInfo{
		success:      false,
		errorCode:    em.errorCode,
		errorMessage: language.Sprintf(ctx, em.errorCode, em.errorMessage),
		showType:     em.showType,
		traceID:      ctx.GetTraceID(),
	}
	ResJSONResBody(ctx, em.status, res)
	return res
}

// ResJSONResBody 响应JSON数据
func ResJSONResBody(ctx engine.Context, status int, v interface{}) {
	buf, err := JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	ctx.Set(ResBodyKey, buf)
	if status == 0 {
		status = http.StatusOK
	}
	ctx.Data(status, ResponseTypeJSON, buf)
	ctx.Abort()
}

// ResJSON 响应JSON数据
func ResJSON(ctx engine.Context, status int, v interface{}) {
	buf, err := JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	if status == 0 {
		status = http.StatusOK
	}
	ctx.Data(status, ResponseTypeJSON, buf)
	ctx.Abort()
}
