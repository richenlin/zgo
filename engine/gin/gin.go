// Copyright 2020 Kratos Team. All rights reserved.
// Use of that source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"time"
	"zgo/engine"
	"zgo/modules/config"
	"zgo/modules/result"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
func (f *WebFramework) Name() string {
	return "gin"
}

// Target target
func (f *WebFramework) Target() interface{} {
	return f.target
}

// Run 运行服务器
func (f *WebFramework) Run(addr ...string) error {
	return f.target.(*gin.Engine).Run(addr...)
}

// RunHandler 获取服务器句柄
func (f *WebFramework) RunHandler() http.Handler {
	return f.target.(*gin.Engine)
}

// NoMethod 未匹配到方法
func (f *WebFramework) NoMethod(handlers ...engine.HandlerFunc) {
	f.target.(*gin.Engine).NoMethod(ConvertHandlerFunc(handlers)...)
}

// NoRoute 未匹配到路由
func (f *WebFramework) NoRoute(handlers ...engine.HandlerFunc) {
	f.target.(*gin.Engine).NoMethod(ConvertHandlerFunc(handlers)...)
}

//========================================================= 分割线
//========================================================= 分割线
//========================================================= 分割线

// WebRouter router
type WebRouter struct {
	WebRoutes
}

// Group 路由分组
func (r *WebRouter) Group(relativePath string, handlers ...engine.HandlerFunc) engine.IRouter {
	router := r.target.(gin.IRouter).Group(relativePath, ConvertHandlerFunc(handlers)...)
	if r.target == router {
		return r
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
func (r *WebRoutes) newRoutes(routes gin.IRoutes) engine.IRoutes {
	if r.target == routes {
		return r
	}
	result := new(WebRouter)
	result.target = routes
	return result
}

// Use use
func (r *WebRoutes) Use(middleware ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).Use(ConvertHandlerFunc(middleware)...)
	return r.newRoutes(routes)
}

// Handle handle
func (r *WebRoutes) Handle(httpMethod string, relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).Handle(httpMethod, relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// Any any
func (r *WebRoutes) Any(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).Any(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// GET get
func (r *WebRoutes) GET(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).GET(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// POST post
func (r *WebRoutes) POST(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).POST(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// DELETE delete
func (r *WebRoutes) DELETE(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).DELETE(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// PATCH patch
func (r *WebRoutes) PATCH(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).PATCH(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// PUT put
func (r *WebRoutes) PUT(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).PUT(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// OPTIONS options
func (r *WebRoutes) OPTIONS(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).OPTIONS(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// HEAD head
func (r *WebRoutes) HEAD(relativePath string, handlers ...engine.HandlerFunc) engine.IRoutes {
	routes := r.target.(gin.IRoutes).HEAD(relativePath, ConvertHandlerFunc(handlers)...)
	return r.newRoutes(routes)
}

// StaticFile static file
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (r *WebRoutes) StaticFile(relativePath, filepath string) engine.IRoutes {
	routes := r.target.(gin.IRoutes).StaticFile(relativePath, filepath)
	return r.newRoutes(routes)
}

// Static static
// router.Static("/static", "/var/www")
func (r *WebRoutes) Static(relativePath, root string) engine.IRoutes {
	routes := r.target.(gin.IRoutes).Static(relativePath, root)
	return r.newRoutes(routes)
}

// StaticFS static fs
// Gin by default user: gin.Dir()
func (r *WebRoutes) StaticFS(relativePath string, fs http.FileSystem) engine.IRoutes {
	routes := r.target.(gin.IRoutes).StaticFS(relativePath, fs)
	return r.newRoutes(routes)
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

// Request request
func (c *WebContext) Request() *http.Request {
	return c.ctx.Request
}

// ResponseWriter Writer
func (c *WebContext) ResponseWriter() http.ResponseWriter {
	return c.ctx.Writer
}

//========================================================= 分割线

// Set set
func (c *WebContext) Set(key string, value interface{}) {
	c.ctx.Set(key, value)
}

// Get set
func (c *WebContext) Get(key string) (value interface{}, exists bool) {
	return c.ctx.Get(key)
}

//========================================================= 分割线

// Header header
func (c *WebContext) Header(key, value string) {
	c.ctx.Header(key, value)
}

// GetHeader header
func (c *WebContext) GetHeader(key string) string {
	return c.ctx.GetHeader(key)
}

// GetCookie 获取cookie
func (c *WebContext) GetCookie(name string) (string, error) {
	return c.ctx.Cookie(name)
}

// GetQuery query
func (c *WebContext) GetQuery(key string) (string, bool) {
	return c.ctx.GetQuery(key)
}

// GetQueryArray query
func (c *WebContext) GetQueryArray(key string) ([]string, bool) {
	return c.ctx.GetQueryArray(key)
}

// GetQueryMap query
func (c *WebContext) GetQueryMap(key string) (map[string]string, bool) {
	return c.ctx.GetQueryMap(key)
}

// DefaultQuery query
func (c *WebContext) DefaultQuery(key, defaultValue string) string {
	if value, ok := c.GetQuery(key); ok {
		return value
	}
	return defaultValue
}

// GetParam param
func (c *WebContext) GetParam(key string) string {
	return c.ctx.Param(key)
}

// ResponseStatus Writer
func (c *WebContext) ResponseStatus() int {
	return c.ctx.Writer.Status()
}

// ResponseSize Writer
func (c *WebContext) ResponseSize() int {
	return c.ctx.Writer.Size()
}

//========================================================= 分割线

// JSON json
func (c *WebContext) JSON(code int, obj interface{}) {
	c.ctx.JSON(code, obj)
}

// JSONP jsonp
func (c *WebContext) JSONP(code int, obj interface{}) {
	c.ctx.JSONP(code, obj)
}

// Data data
func (c *WebContext) Data(code int, contentType string, data []byte) {
	c.ctx.Data(code, contentType, data)
}

//========================================================= 分割线

// File file
func (c *WebContext) File(file string) {
	c.ctx.File(file)
}

// Next next
func (c *WebContext) Next() {
	c.ctx.Next()
}

// Abort abort
func (c *WebContext) Abort() {
	c.ctx.Abort()
}

//========================================================= 分割线

// ClientIP client ip
func (c *WebContext) ClientIP() string {
	return c.ctx.ClientIP()
}

// GetTraceID 追踪ID
func (c *WebContext) GetTraceID() string {
	if v, ok := c.ctx.Get(result.TraceIDKey); ok && v != "" {
		return v.(string)
	}

	// 优先从请求头中获取请求ID
	traceID := c.ctx.GetHeader("X-Request-Id")
	if traceID == "" {
		// 没有自建
		v, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}
		traceID = v.String()
	}
	c.ctx.Set(result.TraceIDKey, traceID)
	return traceID
}

// GetUserInfo 用户ID
func (c *WebContext) GetUserInfo() (engine.UserInfo, bool) {
	v, ok := c.ctx.Get(result.UserInfoKey)
	return v.(engine.UserInfo), ok
}

// SetUserInfo 用户ID
func (c *WebContext) SetUserInfo(user engine.UserInfo) {
	c.ctx.Set(result.UserInfoKey, user)
}

/*****************************************/
/***** GOLANG.ORG/X/NET/HTTP/BINDING *****/
/*****************************************/

// ShouldBind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//     "application/json" --> JSON binding
//     "application/xml"  --> XML binding
// otherwise --> returns an error
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.Bind() but this method does not set the response status code to 400 and abort if the json is not valid.
func (c *WebContext) ShouldBind(obj interface{}) error {
	return c.ctx.ShouldBind(obj)
}

// ShouldBindURI binds the passed struct pointer using the specified binding engine.
func (c *WebContext) ShouldBindURI(obj interface{}) error {
	return c.ctx.ShouldBindUri(obj)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func (c *WebContext) ShouldBindWith(obj interface{}, b binding.Binding) error {
	return c.ctx.ShouldBindWith(obj)
}

// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before binding. So you should use
// ShouldBindWith for better performance if you need to call only once.
func (c *WebContext) ShouldBindBodyWith(obj interface{}, bb binding.BindingBody) (err error) {
	return c.ctx.ShouldBindBodyWith(obj)
}

// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
func (c *WebContext) ShouldBindJSON(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.JSON)
}

// Bind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//     "application/json" --> JSON binding
//     "application/xml"  --> XML binding
// otherwise --> returns an error.
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
func (c *WebContext) Bind(obj interface{}) error {
	return c.ctx.Bind(obj)
}

// BindURI binds the passed struct pointer using binding.Uri.
// It will abort the request with HTTP 400 if any error occurs.
func (c *WebContext) BindURI(obj interface{}) error {
	return c.ctx.BindUri(obj)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func (c *WebContext) MustBindWith(obj interface{}, b binding.Binding) error {
	return c.ctx.MustBindWith(obj)
}

// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
func (c *WebContext) BindJSON(obj interface{}) error {
	return c.MustBindWith(obj, binding.JSON)
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

// Deadline always returns c there is no deadline (ok==false)
func (c *WebContext) Deadline() (deadline time.Time, ok bool) {
	return c.Deadline()
}

// Done always returns nil (chan which will wait forever)
func (c *WebContext) Done() <-chan struct{} {
	return c.Done()
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (c *WebContext) Err() error {
	return c.Err()
}

// Value returns the value associated with this context for key, or nil
func (c *WebContext) Value(key interface{}) interface{} {
	return c.Value(key)
}
