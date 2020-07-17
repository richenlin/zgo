// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

// 注意,该文件不会参与编译, 需要通过[make wire]命令生成wire_gen.go文件编译
// 文件第一行是通知golang编译器忽略该文件
/*
	系统使用google/wire作为框架的依赖注入组件,需要注意一下细节.

	1.系统没有自发现依赖组件,需要人工配置
	2.当依赖发生变更后,需要执行[make wire]命令更新
	3.当前可能会注入无用属性,比如[Swagger, Healthz], 该内容主要通知wire,执行构造方法
	4.Injector中的顺序,决定了整体的执行顺序
	5.wire_gen.go文件不要编辑,为自动生成
*/
import (
	"github.com/suisrc/zgo/app/api"
	"github.com/suisrc/zgo/middlewire"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(
	wire.Struct(new(Injector), "*"),
	middlewire.NewSwagger,
	middlewire.NewHealthz,
)

//======================================
// 注入控制器
//======================================

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine    *gin.Engine
	Endpoints *api.Endpoints
	Swagger   middlewire.Swagger
	Healthz   middlewire.Healthz

	//Enforcer *casbin.SyncedEnforcer
}

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InjectorSet,              // wire索引
		middlewire.DefaultGinSet, // gin引擎
		api.EndpointSet,          // 服务接口
		//casbin内容
		//casbinsqlx.CasbinAdapterSet,
	)
	return new(Injector), nil, nil
}

//======================================
// END
//======================================
