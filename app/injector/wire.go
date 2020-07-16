// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine *gin.Engine
	Routes *InitRoutesResult
	//Enforcer *casbin.SyncedEnforcer
}

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InjectorSet,
		InitGinEngine,
		InitRoutesSet,

		//casbin内容
		//casbinsqlx.CasbinAdapterSet,
	)
	return new(Injector), nil, nil
}
