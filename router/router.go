package router

import (
	"zgo/middleware"
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
)

// NewContextPath 新建根路由
func NewContextPath(app *gin.Engine) RootPath {

	// 默认中间件
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	root := app.Group(GetContextPath())
	return root
}

// RootPath 根路由路由
type RootPath gin.IRouter

// GetContextPath 服务器根目录
func GetContextPath() string {
	return config.C.HTTP.ContextPath
}
