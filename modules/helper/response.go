package helper

import (
	"net/http"
	"zgo/modules/language"

	"github.com/gin-gonic/gin"
)

// ResError 包装响应错误
func ResError(ctx *gin.Context, em *ErrorModel) error {
	res := &ErrorInfo{
		success:      false,
		errorCode:    em.errorCode,
		errorMessage: language.Sprintf(ctx, em.errorCode, em.errorMessage),
		showType:     em.showType,
		traceID:      GetTraceID(ctx),
	}
	//ctx.JSON(http.StatusOK, res)
	//ctx.Abort()
	ResJSON(ctx, em.status, res)
	return res
}

// ResErrorResBody 包装响应错误
func ResErrorResBody(ctx *gin.Context, em *ErrorModel) error {
	res := &ErrorInfo{
		success:      false,
		errorCode:    em.errorCode,
		errorMessage: language.Sprintf(ctx, em.errorCode, em.errorMessage),
		showType:     em.showType,
		traceID:      GetTraceID(ctx),
	}
	ResJSONResBody(ctx, em.status, res)
	return res
}

// ResJSONResBody 响应JSON数据
func ResJSONResBody(ctx *gin.Context, status int, v interface{}) {
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
func ResJSON(ctx *gin.Context, status int, v interface{}) {
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
