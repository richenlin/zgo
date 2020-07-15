package router

import (
	"net/http"
	"zgo/engine"
)

// DemoAPI 根路由路由
type DemoAPI engine.IRoutes

// NewDemoAPI 新建根路由
func NewDemoAPI(r RootPath) DemoAPI {
	r.GET("demo", func(ctx engine.Context) {
		ctx.JSON(http.StatusOK, engine.H{
			"message": "demo, world",
		})
	})
	return nil
}
