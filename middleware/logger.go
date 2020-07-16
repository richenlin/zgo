package middleware

import (
	"time"
	"zgo/modules/helper"
	"zgo/modules/logger"

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
		timeConsuming := time.Since(start).Nanoseconds() / 1e6

		trace := logger.StartTrace(c)
		path := c.Request.URL.Path
		method := c.Request.Method
		client := c.ClientIP()
		status := c.Writer.Status
		url := c.Request.URL.String()

		fields := make(map[string]interface{})
		fields["ip"] = client
		fields["method"] = method
		fields["url"] = url
		if v, ok := c.Get(helper.ReqBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["body"] = string(b)
			}
		}
		fields["res_status"] = status
		if v, ok := c.Get(helper.ResBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["res_body"] = string(b)
			}
		}

		trace.WithFields(fields).Infof("[http] %s-%s-%s-%d(%dms)", path, method, client, status, timeConsuming)
	}
}
