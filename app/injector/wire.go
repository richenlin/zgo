// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"zgo/engine"
	"zgo/engine/gin"
	"zgo/router"

	"github.com/google/wire"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine   engine.IEngine
	Result   *InitWebResult
	HelloAPI router.HelloAPI
	DemoAPI  router.DemoAPI
	//Enforcer *casbin.SyncedEnforcer
}

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InjectorSet, InitWebSet, // injector

		gin.WebFrameworkSet,
		router.NewContextPath,
		router.NewDemoAPI,
		router.NewHelloAPI,
		//casbin内容
		//casbinsqlx.CasbinAdapterSet,
	)
	return new(Injector), nil, nil
}
