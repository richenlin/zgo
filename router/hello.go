package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloAPI 根路由路由
type HelloAPI gin.IRoutes

// NewHelloAPI 新建根路由
func NewHelloAPI(r RootPath) HelloAPI {
	r.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello, world",
		})
	})
	return nil
}
