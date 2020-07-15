package router

import (
	"zgo/engine"
	"zgo/middleware"
	"zgo/modules/config"
)

// NewContextPath 新建根路由
func NewContextPath(engine engine.IEngine) RootPath {

	// 默认中间件
	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NoRouteHandler())

	root := engine.Group(GetContextPath())
	return root
}

// RootPath 根路由路由
type RootPath engine.IRouter

// GetContextPath 服务器根目录
func GetContextPath() string {
	return config.C.HTTP.ContextPath
}
