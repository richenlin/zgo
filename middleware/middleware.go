package middleware

import (
	"zgo/engine"
	"zgo/modules/result"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() engine.HandlerFunc {
	return func(c engine.Context) {
		result.ResError(c, result.Err405MethodNotAllowed)
		// Abort, 终止
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() engine.HandlerFunc {
	return func(c engine.Context) {
		result.ResError(c, result.Err404NotFound)
		// Abort, 终止
	}
}

// EmptyMiddleware 不执行业务处理的中间件
func EmptyMiddleware() engine.HandlerFunc {
	return func(c engine.Context) {
		c.Next() // Pass, 跳过
	}
}

// TraceMiddleware 跟踪ID中间件
func TraceMiddleware(skippers ...SkipperFunc) engine.HandlerFunc {
	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		// 内容交由engine.Context中的方法处理
		// 这里不在具有任何意义
		c.GetTraceID()
		c.Next()
	}
}
