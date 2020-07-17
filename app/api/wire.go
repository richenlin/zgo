package api

import (
	ser "github.com/suisrc/zgo/app/service"
	"github.com/suisrc/zgo/middlewire"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// EndpointSet wire注入声明
var EndpointSet = wire.NewSet(
	ser.ServiceSet,                 // 系统提供的服务列表
	wire.Struct(new(Options), "*"), // 初始化接口参数
	InitEndpoints,                  // 初始化接口方法

	// 接口列表
	HelloSet,
)

//=====================================
// Endpoint
//=====================================

// Options options
type Options struct {
	Engine *gin.Engine
	Router middlewire.Router
	Hello  *Hello
}

// Endpoints result
type Endpoints struct {
}

// InitEndpoints init
func InitEndpoints(o *Options) *Endpoints {

	r := o.Router
	test := r.Group("test")
	{
		test.GET("hello", o.Hello.Hello)
		test.GET("hello2", o.Hello.Hello)
	}

	return &Endpoints{}
}

//=====================================
// v1 router
//=====================================

// V1Router v1 path
type V1Router gin.IRouter

// V1RouterFunc  v1 path
func V1RouterFunc(r middlewire.Router) V1Router {
	return r.Group("v1")
}
