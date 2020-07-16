package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DemoAPI 根路由路由
type DemoAPI gin.IRoutes

// NewDemoAPI 新建根路由
func NewDemoAPI(r RootPath) DemoAPI {
	r.GET("demo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "demo, world",
		})
	})
	return nil
}
