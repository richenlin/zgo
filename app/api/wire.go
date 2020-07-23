package api

import (
	"github.com/casbin/casbin/v2"
	service "github.com/suisrc/zgo/app/ser"
	"github.com/suisrc/zgo/middlewire"
	casbinjson "github.com/suisrc/zgo/modules/casbin/adapter/json"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// EndpointSet wire注入声明
var EndpointSet = wire.NewSet(
	service.ServiceSet,             // 系统提供的服务列表
	wire.Struct(new(Options), "*"), // 初始化接口参数
	InitEndpoints,                  // 初始化接口方法
	casbinjson.CasbinAdapterSet,    // Casbin依赖

	// 接口注册
	wire.Struct(new(Demo), "*"),
)

//=====================================
// Endpoint
//=====================================

// Options options
type Options struct {
	Engine   *gin.Engine
	Enforcer *casbin.SyncedEnforcer
	Router   middlewire.Router
	Demo     *Demo
}

// Endpoints result
type Endpoints struct {
}

// InitEndpoints init
func InitEndpoints(o *Options) *Endpoints {
	r := o.Router

	demo := r.Group("demo")
	{
		//demo.Use(middleware.CasbinMiddleware(o.Enforcer))
		o.Demo.Register(demo)
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
