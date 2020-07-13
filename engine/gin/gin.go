// Copyright 2020 Kratos Team. All rights reserved.
// Use of that source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"zgo/engine"
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ engine.IEngine = (*WebFramework)(nil)

var _ engine.Context = &WebContext{}
var _ engine.IRouter = &WebRouter{}
var _ engine.IRoutes = &WebRoutes{}

// WebFrameworkSet 注入WebFramework
var WebFrameworkSet = wire.NewSet(InitWebFramework, wire.Bind(new(engine.IEngine), new(*WebFramework)))

// InitWebFramework 初始化web框架
func InitWebFramework() (*WebFramework, error) {
	gin.SetMode(config.C.RunMode)
	//gin.SetMode(gin.DebugMode)

	engine := gin.New()

	// 默认
	engine.Use(gin.Logger(), gin.Recovery())

	return &WebFramework{
		target: engine,
	}, nil
}

// ConvertHandlerFunc 转换 engine.HandlerFunc => gin.HandlerFunc
func ConvertHandlerFunc(handler engine.HandlerFunc) func(*gin.Context) {
	return func(ctx *gin.Context) {
		handler(interface{}(&WebContext{
			target: ctx,
		}).(*engine.Context))
	}
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebFramework web框架
type WebFramework struct {
	WebRouter
	target *gin.Engine
}

// getTarget target
func (that *WebFramework) getTarget() *gin.Engine {
	return that.target
}

// Name web框架的名称
func (that *WebFramework) Name() string {
	return "gin"
}

// Run 运行服务器
func (that *WebFramework) Run(addr ...string) error {
	return that.target.Run(addr...)
}

// RunHandler 获取服务器句柄
func (that *WebFramework) RunHandler() http.Handler {
	return that.target
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebContext context
type WebContext struct {
	target *gin.Context
}

// Header header
func (that *WebContext) Header(key, value string) {
	that.target.Header(key, value)
}

// GetHeader header
func (that *WebContext) GetHeader(key string) string {
	return that.target.GetHeader(key)
}

// Set set
func (that *WebContext) Set(key string, value interface{}) {
	that.target.Set(key, value)
}

// JSON json
func (that *WebContext) JSON(code int, obj interface{}) {
	that.target.JSON(code, obj)
}

// JSONP jsonp
func (that *WebContext) JSONP(code int, obj interface{}) {
	that.target.JSONP(code, obj)
}

// Next next
func (that *WebContext) Next() {
	that.target.Next()
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebRouter router
type WebRouter struct {
	WebRoutes
	target gin.IRouter
}

// getTarget target
func (that *WebRouter) getTarget() gin.IRouter {
	return that.target
}

// Group 路由分组
func (that *WebRouter) Group(relativePath string, handler engine.HandlerFunc) engine.IRouter {
	router := that.getTarget().Group(relativePath, ConvertHandlerFunc(handler))
	if that.getTarget() == router {
		return that
	}
	return &WebRouter{target: router}
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebRoutes routes
type WebRoutes struct {
	target gin.IRoutes
}

// getTarget target
func (that *WebRoutes) getTarget() gin.IRoutes {
	return that.target
}

// newRoutes routes
func (that *WebRoutes) newRoutes(routes gin.IRoutes) engine.IRoutes {
	if that.getTarget() == routes {
		return that
	}
	return &WebRoutes{target: routes}
}

// Use use
func (that *WebRoutes) Use(middleware engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().Use(ConvertHandlerFunc(middleware))
	return that.newRoutes(routes)
}

// Handle handle
func (that *WebRoutes) Handle(httpMethod string, relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().Handle(httpMethod, relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// Any any
func (that *WebRoutes) Any(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().Any(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// GET get
func (that *WebRoutes) GET(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().GET(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// POST post
func (that *WebRoutes) POST(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().POST(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// DELETE delete
func (that *WebRoutes) DELETE(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().DELETE(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// PATCH patch
func (that *WebRoutes) PATCH(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().PATCH(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// PUT put
func (that *WebRoutes) PUT(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().PUT(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// OPTIONS options
func (that *WebRoutes) OPTIONS(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().OPTIONS(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// HEAD head
func (that *WebRoutes) HEAD(relativePath string, handler engine.HandlerFunc) engine.IRoutes {
	routes := that.getTarget().HEAD(relativePath, ConvertHandlerFunc(handler))
	return that.newRoutes(routes)
}

// StaticFile static file
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (that *WebRoutes) StaticFile(relativePath, filepath string) engine.IRoutes {
	routes := that.getTarget().StaticFile(relativePath, filepath)
	return that.newRoutes(routes)
}

// Static static
// router.Static("/static", "/var/www")
func (that *WebRoutes) Static(relativePath, root string) engine.IRoutes {
	routes := that.getTarget().Static(relativePath, root)
	return that.newRoutes(routes)
}

// StaticFS static fs
// Gin by default user: gin.Dir()
func (that *WebRoutes) StaticFS(relativePath string, fs http.FileSystem) engine.IRoutes {
	routes := that.getTarget().StaticFS(relativePath, fs)
	return that.newRoutes(routes)
}
