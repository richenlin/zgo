package engine

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin/binding"
)

// Context web框架的上下文
type Context interface {
	GetHeader(key string) string
	Header(key, value string)
	Cookie(name string) (string, error)

	Request() *http.Request
	ResponseWriter() http.ResponseWriter
	ResponseStatus() int
	ResponseSize() int

	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)

	JSON(code int, obj interface{})
	JSONP(code int, obj interface{})
	Data(code int, contentType string, data []byte)

	File(file string)
	Next()
	Abort()

	ClientIP() string
	GetTraceID() string
	GetUserInfo() (UserInfo, bool)
	SetUserInfo(UserInfo)

	/*****************************************/
	/***** GOLANG.ORG/X/NET/HTTP/BINDING *****/
	/*****************************************/
	ShouldBind(obj interface{}) error
	ShouldBindURI(obj interface{}) error
	ShouldBindWith(obj interface{}, b binding.Binding)
	ShouldBindBodyWith(obj interface{}, bb binding.BindingBody)

	ShouldBindJSON(obj interface{}) error

	Bind(obj interface{}) error
	BindURI(obj interface{}) error
	MustBindWith(obj interface{}, b binding.Binding) error

	BindJSON(obj interface{}) error

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
