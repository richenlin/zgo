package router

import (
	"net/http"
	"zgo/engine"
)

// HelloAPI 根路由路由
type HelloAPI engine.IRoutes

// NewHelloAPI 新建根路由
func NewHelloAPI(r RootPath) HelloAPI {
	r.GET("hello", func(ctx engine.Context) {
		ctx.JSON(http.StatusOK, engine.H{
			"message": "hello, world",
		})
	})
	return nil
}
