package helper

import (
	"net/http"

	"github.com/suisrc/zgo/modules/language"

	"github.com/gin-gonic/gin"
)

// ResSuccess 包装响应错误
func ResSuccess(ctx *gin.Context, v interface{}) error {
	res := NewOK(ctx, v)
	//ctx.JSON(http.StatusOK, res)
	//ctx.Abort()
	ResJSON(ctx, http.StatusOK, res)
	return res
}

// ResError 包装响应错误
func ResError(ctx *gin.Context, em *ErrorModel) error {
	res := ErrorInfo{
		Success:      false,
		ErrorCode:    em.ErrorCode,
		ErrorMessage: language.Sprintf(ctx, em.ErrorCode, em.ErrorMessage),
		ShowType:     em.ShowType,
		TraceID:      GetTraceID(ctx),
	}
	//ctx.JSON(http.StatusOK, res)
	//ctx.Abort()
	ResJSON(ctx, em.Status, res)
	return &res
}

// ResErrorResBody 包装响应错误
func ResErrorResBody(ctx *gin.Context, em *ErrorModel) error {
	res := ErrorInfo{
		Success:      false,
		ErrorCode:    em.ErrorCode,
		ErrorMessage: language.Sprintf(ctx, em.ErrorCode, em.ErrorMessage),
		ShowType:     em.ShowType,
		TraceID:      GetTraceID(ctx),
	}
	ResJSONResBody(ctx, em.Status, res)
	return &res
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
