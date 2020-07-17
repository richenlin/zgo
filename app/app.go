package app

import (
	"context"
	"zgo/app/injector"
	cmd "zgo/modules/app"

	"github.com/gin-gonic/gin"

	// 引入swagger
	_ "zgo/app/swagger"
)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) cmd.Option {
	return func(o *cmd.Options) {
		o.ConfigFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) cmd.Option {
	return func(o *cmd.Options) {
		o.Version = s
	}
}

// SetBuildInjector 设定版本号
func SetBuildInjector(f cmd.BuildInjector) cmd.Option {
	return func(o *cmd.Options) {
		o.BuildInjector = f
	}
}

// Run 运行服务
func Run(ctx context.Context, opts ...cmd.Option) error {
	injectorOption := SetBuildInjector(func() (*gin.Engine, func(), error) {
		injector, clean, err := injector.BuildInjector()
		if err != nil {
			return nil, nil, err
		}
		return injector.Engine, clean, err
	})
	return cmd.Run(ctx, append(opts, injectorOption)...)
}
