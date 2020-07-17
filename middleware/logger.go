package middleware

import (
	"time"
	"github.com/suisrc/zgo/modules/helper"
	"github.com/suisrc/zgo/modules/logger"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()

		fields := make(map[string]interface{})
		fields["ip"] = helper.GetClientIP(c)
		fields["method"] = c.Request.Method
		fields["path"] = c.Request.URL.Path
		fields["url"] = c.Request.URL.String()
		if v, ok := c.Get(helper.ReqBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["body"] = string(b)
			}
		}
		fields["res_status"] = c.Writer.Status()
		fields["res_time"] = time.Since(start).Nanoseconds() / 1e6
		if v, ok := c.Get(helper.ResBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["res_body"] = string(b)
			}
		}

		logger.StartTrace(c).WithFields(fields).Infof("[access]")
	}
}
