package engine

import "net/http"

// Context web框架的上下文
type Context interface {
	GetHeader(key string) string
	Header(key, value string)

	Set(key string, value interface{})

	JSON(code int, obj interface{})
	JSONP(code int, obj interface{})

	Next()
}

// HandlerFunc 上下文操作句柄
type HandlerFunc func(*Context)

// IRouter web路由组
type IRouter interface {
	IRoutes
	Group(string, HandlerFunc) IRouter
}

// IRoutes web路由
type IRoutes interface {
	Use(HandlerFunc) IRoutes

	Handle(string, string, HandlerFunc) IRoutes
	Any(string, HandlerFunc) IRoutes
	GET(string, HandlerFunc) IRoutes
	POST(string, HandlerFunc) IRoutes
	DELETE(string, HandlerFunc) IRoutes
	PATCH(string, HandlerFunc) IRoutes
	PUT(string, HandlerFunc) IRoutes
	OPTIONS(string, HandlerFunc) IRoutes
	HEAD(string, HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}
