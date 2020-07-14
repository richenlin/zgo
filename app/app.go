package app

import (
	"context"
	"os"
	"zgo/engine"

	// 引入swagger
	_ "zgo/app/swagger"
	// 引入app
	"zgo/app/injector"
	// 引入modules
	"zgo/modules/config"
	"zgo/modules/logger"
)

type options struct {
	ConfigFile string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// RunServer 启动服务
func RunServer(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	// 加载配置文件
	config.MustLoad(o.ConfigFile)
	config.PrintWithJSON()

	// 启动日志
	logger.Printf(ctx, "服务启动，运行模式：%s，版本号：%s，进程号：%d", config.C.RunMode, o.Version, os.Getpid())

	// 初始化日志模块
	loggerCleanFunc, err := InitLogger(ctx)
	if err != nil {
		return nil, err
	}

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := injector.BuildInjector()
	if err != nil {
		return nil, err
	}

	// 初始化HTTP服务
	shutdownServerFunc := engine.RunHTTPServer(ctx, injector.Engine.RunHandler())

	return func() {
		shutdownServerFunc()
		injectorCleanFunc()
		loggerCleanFunc()
	}, nil
}

// Run 运行服务
func Run(ctx context.Context, opts ...Option) error {
	return engine.Run(ctx, func() (func(), error) {
		return RunServer(ctx, opts...)
	})
}
