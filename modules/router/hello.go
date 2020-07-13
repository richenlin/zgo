package router

import (
	"net/http"
	"zgo/engine"
)

// HelloAPI 根路由路由
type HelloAPI engine.IRoutes

// NewHelloAPI 新建根路由
func NewHelloAPI(root RootRouter) HelloAPI {
	root.GET("hello", func(ctx engine.IContext) {
		ctx.JSON(http.StatusOK, engine.H{
			"message": "hello, world",
		})
	})
	return nil
}
