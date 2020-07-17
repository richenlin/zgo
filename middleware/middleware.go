package middleware

import (
	"github.com/suisrc/zgo/modules/helper"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		helper.ResError(c, &helper.Err405MethodNotAllowed)
		// Abort, 终止
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		helper.ResError(c, &helper.Err404NotFound)
		// Abort, 终止
	}
}

// EmptyMiddleware 不执行业务处理的中间件
func EmptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Pass, 跳过
	}
}

// TraceMiddleware 跟踪ID中间件
func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		// 内容交由gin.Context中的方法处理
		// 这里不在具有任何意义
		helper.GetTraceID(c)
		c.Next()
	}
}
