package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/suisrc/zgo/app/service"
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/middlewire"
	"github.com/suisrc/zgo/modules/auth"
	casbinjson "github.com/suisrc/zgo/modules/casbin/adapter/json"
	"github.com/suisrc/zgo/modules/config"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// EndpointSet wire注入声明
var EndpointSet = wire.NewSet(
	service.ServiceSet,             // 系统提供的服务列表
	wire.Struct(new(Options), "*"), // 初始化接口参数
	InitEndpoints,                  // 初始化接口方法
	casbinjson.CasbinAdapterSet,    // Casbin依赖
	NewAuther,                      // Auther注册

	// 接口注册
	wire.Struct(new(Auth), "*"),
	wire.Struct(new(Signin), "*"),
	wire.Struct(new(User), "*"),
)

//=====================================
// Endpoint
//=====================================

// Options options
type Options struct {
	Engine   *gin.Engine
	Enforcer *casbin.SyncedEnforcer
	Auther   auth.Auther
	Router   middlewire.Router

	// 接口注入
	Auth   *Auth
	Signin *Signin
	User   *User
}

// Endpoints result
type Endpoints struct {
}

// InitEndpoints init
func InitEndpoints(o *Options) *Endpoints {
	// 在根路由注册通用授权接口, (没有ContextPath限定,一般是给nginx使用)
	// 在nginx注册认证接口时候,请放行zgo服务器接口,防止重复认证
	o.Auth.Register(o.Engine)

	// ContextPath路由
	r := o.Router
	// 服务器授权控制器
	uac := middleware.UserAuthCasbinMiddleware(
		o.Auther,
		o.Enforcer,
		// sign 登陆接口需要排除
		middleware.AllowPathPrefixSkipper(middleware.JoinPath(config.C.HTTP.ContextPath, "sign")),
	)
	// 增加权限认证
	r.Use(uac)
	// 注册登陆接口
	o.Signin.Register(r)
	o.User.Register(r)

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
