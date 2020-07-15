package language

import (
	"fmt"
	"zgo/engine"
)

// Sprintf 格式化异常, 为国际化提供后端支持
func Sprintf(ctx engine.Context, code string, format string, args ...interface{}) string {
	if format == "" {
		return ""
	}
	//logger.Infof(ctx, format, args...)
	return fmt.Sprintf(format, args...)
}
