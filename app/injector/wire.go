// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"zgo/engine/gin"

	"github.com/google/wire"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		gin.WebFrameWorkSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
