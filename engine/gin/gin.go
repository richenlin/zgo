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

var _ engine.IContext = &WebContext{}
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

	result := new(WebFramework)
	result.target = engine
	return result, nil
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebFramework web框架
type WebFramework struct {
	WebRouter
}

// Name web框架的名称
func (that *WebFramework) Name() string {
	return "gin"
}

// Run 运行服务器
func (that *WebFramework) Run(addr ...string) error {
	return that.target.(*gin.Engine).Run(addr...)
}

// RunHandler 获取服务器句柄
func (that *WebFramework) RunHandler() http.Handler {
	return that.target.(*gin.Engine)
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebRouter router
type WebRouter struct {
	WebRoutes
}

// Group 路由分组
func (that *WebRouter) Group(relativePath string, handlers ...engine.HandlerFunc) engine.IRouter {
	router := that.target.(gin.IRouter).Group(relativePath, ConvertHandlerFunc(handlers)...)
	if that.target == router {
		return that
	}
	result := new(WebRouter)
	result.target = router
	return result
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebRoutes routes
type WebRoutes struct {
	//target gin.IRoutes
	target interface{}
}

// newRoutes routes
func (that *WebRoutes) newRoutes(routes gin.IRoutes) engine.IRoutes {
	if that.target == routes {
		return that
	}
	result := new(WebRouter)
	result.target = routes
	return result
}

// Use use
func (that *WebRoutes) Use(middleware ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).Use(ConvertHandlerFunc(middleware)...)
	return that.newRoutes(routes)
}

// Handle handle
func (that *WebRoutes) Handle(httpMethod string, relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).Handle(httpMethod, relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// Any any
func (that *WebRoutes) Any(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).Any(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// GET get
func (that *WebRoutes) GET(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).GET(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// POST post
func (that *WebRoutes) POST(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).POST(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// DELETE delete
func (that *WebRoutes) DELETE(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).DELETE(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// PATCH patch
func (that *WebRoutes) PATCH(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).PATCH(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// PUT put
func (that *WebRoutes) PUT(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).PUT(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// OPTIONS options
func (that *WebRoutes) OPTIONS(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).OPTIONS(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// HEAD head
func (that *WebRoutes) HEAD(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := that.target.(gin.IRoutes).HEAD(relativePath, ConvertHandlerFunc(handlers)...)
	return that.newRoutes(routes)
}

// StaticFile static file
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (that *WebRoutes) StaticFile(relativePath, filepath string) engine.IRoutes {
	routes := that.target.(gin.IRoutes).StaticFile(relativePath, filepath)
	return that.newRoutes(routes)
}

// Static static
// router.Static("/static", "/var/www")
func (that *WebRoutes) Static(relativePath, root string) engine.IRoutes {
	routes := that.target.(gin.IRoutes).Static(relativePath, root)
	return that.newRoutes(routes)
}

// StaticFS static fs
// Gin by default user: gin.Dir()
func (that *WebRoutes) StaticFS(relativePath string, fs http.FileSystem) engine.IRoutes {
	routes := that.target.(gin.IRoutes).StaticFS(relativePath, fs)
	return that.newRoutes(routes)
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// ConvertHandlerFunc 转换 engine.HandlerFunc => gin.HandlerFunc
func ConvertHandlerFunc(handlers []engine.HandlerFunc) []gin.HandlerFunc {
	//return interface{}(handlers).([]gin.HandlerFunc)
	var hs []gin.HandlerFunc
	for _, handler := range handlers {
		hs = append(hs, func(ctx *gin.Context) {
			handler(&WebContext{
				ctx: ctx,
			})
		})
	}
	return hs
}

// WebContext context
type WebContext struct {
	ctx *gin.Context
}

// Header header
func (that *WebContext) Header(key, value string) {
	that.ctx.Header(key, value)
}

// GetHeader header
func (that *WebContext) GetHeader(key string) string {
	return that.ctx.GetHeader(key)
}

// Set set
func (that *WebContext) Set(key string, value interface{}) {
	that.ctx.Set(key, value)
}

// JSON json
func (that *WebContext) JSON(code int, obj interface{}) {
	that.ctx.JSON(code, obj)
}

// JSONP jsonp
func (that *WebContext) JSONP(code int, obj interface{}) {
	that.ctx.JSONP(code, obj)
}

// Next next
func (that *WebContext) Next() {
	that.ctx.Next()
}

// RequestMethod method
func (that *WebContext) RequestMethod() string {
	return that.ctx.Request.Method
}

// RequestHeader header
func (that *WebContext) RequestHeader() http.Header {
	return that.ctx.Request.Header
}
