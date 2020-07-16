package language

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Sprintf 格式化异常, 为国际化提供后端支持
func Sprintf(ctx *gin.Context, code string, format string, args ...interface{}) string {
	if format == "" {
		return ""
	}
	//logger.Infof(ctx, format, args...)
	return fmt.Sprintf(format, args...)
}
