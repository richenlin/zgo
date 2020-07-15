// Copyright 2020 Kratos Team. All rights reserved.
// Use of that source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

import (
	"io"
	"net/http"
	"time"
	"zgo/engine"
	"zgo/modules/config"
	"zgo/modules/result"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	app := gin.New()

	// 默认
	// UseMiddlewares(app)

	result := new(WebFramework)
	result.target = app
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

// Target target
func (that *WebFramework) Target() interface{} {
	return that.target
}

// Run 运行服务器
func (that *WebFramework) Run(addr ...string) error {
	return that.target.(*gin.Engine).Run(addr...)
}

// RunHandler 获取服务器句柄
func (that *WebFramework) RunHandler() http.Handler {
	return that.target.(*gin.Engine)
}

// NoMethod 未匹配到方法
func (that *WebFramework) NoMethod(handlers ...engine.HandlerFunc) {
	that.target.(*gin.Engine).NoMethod(ConvertHandlerFunc(handlers)...)
}

// NoRoute 未匹配到路由
func (that *WebFramework) NoRoute(handlers ...engine.HandlerFunc) {
	that.target.(*gin.Engine).NoMethod(ConvertHandlerFunc(handlers)...)
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

//========================================================= 分割线

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

//========================================================= 分割线

// Set set
func (that *WebContext) Set(key string, value interface{}) {
	that.ctx.Set(key, value)
}

// Get set
func (that *WebContext) Get(key string) (value interface{}, exists bool) {
	return that.ctx.Get(key)
}

//========================================================= 分割线

// JSON json
func (that *WebContext) JSON(code int, obj interface{}) {
	that.ctx.JSON(code, obj)
}

// JSONP jsonp
func (that *WebContext) JSONP(code int, obj interface{}) {
	that.ctx.JSONP(code, obj)
}

// Data data
func (that *WebContext) Data(code int, contentType string, data []byte) {
	that.ctx.Data(code, contentType, data)
}

//========================================================= 分割线

// File file
func (that *WebContext) File(file string) {
	that.ctx.File(file)
}

// Next next
func (that *WebContext) Next() {
	that.ctx.Next()
}

// Abort abort
func (that *WebContext) Abort() {
	that.ctx.Abort()
}

//========================================================= 分割线

// RequestMethod method
func (that *WebContext) RequestMethod() string {
	return that.ctx.Request.Method
}

// RequestHeader header
func (that *WebContext) RequestHeader() http.Header {
	return that.ctx.Request.Header
}

// RequestURLPath path
func (that *WebContext) RequestURLPath() string {
	return that.ctx.Request.URL.Path
}

// RequestURLString path
func (that *WebContext) RequestURLString() string {
	return that.ctx.Request.URL.String()
}

// RequestContentLength content length
func (that *WebContext) RequestContentLength() int64 {
	return that.ctx.Request.ContentLength
}

// RequestGetBody request body
func (that *WebContext) RequestGetBody() (io.ReadCloser, error) {
	return that.ctx.Request.GetBody()
}

// RequestSetBody request body
func (that *WebContext) RequestSetBody(body io.ReadCloser) {
	that.ctx.Request.Body = body
}

// ResponseWriter Writer
func (that *WebContext) ResponseWriter() http.ResponseWriter {
	return that.ctx.Writer
}

// ResponseStatus Writer
func (that *WebContext) ResponseStatus() int {
	return that.ctx.Writer.Status()
}

// ResponseSize Writer
func (that *WebContext) ResponseSize() int {
	return that.ctx.Writer.Size()
}

//========================================================= 分割线

// ClientIP client ip
func (that *WebContext) ClientIP() string {
	return that.ctx.ClientIP()
}

// GetTraceID 追踪ID
func (that *WebContext) GetTraceID() string {
	if v, ok := that.ctx.Get(result.TraceIDKey); ok && v != "" {
		return v.(string)
	}

	// 优先从请求头中获取请求ID
	traceID := that.ctx.GetHeader("X-Request-Id")
	if traceID == "" {
		// 没有自建
		v, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}
		traceID = v.String()
	}
	that.ctx.Set(result.TraceIDKey, traceID)
	return traceID
}

// GetUserInfo 用户ID
func (that *WebContext) GetUserInfo() (engine.UserInfo, bool) {
	v, ok := that.ctx.Get(result.UserInfoKey)
	return v.(engine.UserInfo), ok
}

// SetUserInfo 用户ID
func (that *WebContext) SetUserInfo(user engine.UserInfo) {
	that.ctx.Set(result.UserInfoKey, user)
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

// Deadline always returns that there is no deadline (ok==false)
func (that *WebContext) Deadline() (deadline time.Time, ok bool) {
	return that.Deadline()
}

// Done always returns nil (chan which will wait forever)
func (that *WebContext) Done() <-chan struct{} {
	return that.Done()
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (that *WebContext) Err() error {
	return that.Err()
}

// Value returns the value associated with this context for key, or nil
func (that *WebContext) Value(key interface{}) interface{} {
	return that.Value(key)
}
