package middleware

import (
	"time"
	"zgo/engine"
	"zgo/modules/logger"
	"zgo/modules/result"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(skippers ...SkipperFunc) engine.HandlerFunc {
	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()
		timeConsuming := time.Since(start).Nanoseconds() / 1e6

		trace := logger.StartTrace(c)
		path := c.RequestURLPath()
		method := c.RequestMethod()
		client := c.ClientIP()
		status := c.ResponseStatus()
		url := c.RequestURLString()

		fields := make(map[string]interface{})
		fields["ip"] = client
		fields["method"] = method
		fields["url"] = url
		if v, ok := c.Get(result.ReqBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["body"] = string(b)
			}
		}
		fields["res_status"] = status
		if v, ok := c.Get(result.ResBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["res_body"] = string(b)
			}
		}

		trace.WithFields(fields).Infof("[http] %s-%s-%s-%d(%dms)", path, method, client, status, timeConsuming)
	}
}
