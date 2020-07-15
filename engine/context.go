package engine

import (
	"context"
	"io"
	"net/http"
)

// Context web框架的上下文
type Context interface {
	GetHeader(key string) string
	Header(key, value string)

	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)

	JSON(code int, obj interface{})
	JSONP(code int, obj interface{})
	Data(code int, contentType string, data []byte)

	File(file string)
	Next()
	Abort()

	RequestMethod() string
	RequestHeader() http.Header
	RequestURLPath() string
	RequestURLString() string
	RequestContentLength() int64
	RequestGetBody() (io.ReadCloser, error)
	RequestSetBody(io.ReadCloser)

	ResponseWriter() http.ResponseWriter
	ResponseStatus() int
	ResponseSize() int

	ClientIP() string
	GetTraceID() string
	GetUserInfo() (UserInfo, bool)
	SetUserInfo(UserInfo)

	/************************************/
	/***** GOLANG.ORG/X/NET/CONTEXT *****/
	/************************************/
	context.Context
}

// UserInfo 用户信息
type UserInfo interface {
	// GetUserID 用户ID
	GetUserID() string
	// GetRoleID 用户角色
	GetRoleID() string
}

// HandlerFunc 上下文操作句柄
type HandlerFunc func(Context)

// H h -> map
type H map[string]interface{}

// IRouter web路由组
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) IRouter
}

// IRoutes web路由
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}
