package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zgo/modules/config"
	"zgo/modules/logger"

	"github.com/gin-gonic/gin"
)

// RunHTTPServer 初始化http服务
func RunHTTPServer(ctx context.Context, handler http.Handler) func() {
	conf := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Printf(ctx, "HTTP Server is running at %s.", addr)
		var err error
		if conf.CertFile != "" && conf.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(conf.CertFile, conf.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(conf.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
}

// RunWithShutdown 运行服务
func RunWithShutdown(ctx context.Context, runServer func() (func(), error)) error {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	shutdownServer, err := runServer()
	if err != nil {
		return err
	}

	sig := <-sc // 等待服务器中断
	logger.Printf(ctx, "Received a signal [%s]", sig.String())
	// 结束服务
	logger.Printf(ctx, "HTTP Server shutdown ...")
	shutdownServer()
	logger.Printf(ctx, "HTTP Server exiting")
	time.Sleep(time.Second) // 延迟1s, 用于日志等信息保存
	return nil
}

//====================================
// 默认启动程序, 可以直接重新替换
//====================================

// Run 运行服务, 注意,必须对BuildInjector进行初始化
func Run(ctx context.Context, opts ...Option) error {
	return RunWithShutdown(ctx, func() (func(), error) {
		return RunServer(ctx, opts...)
	})
}

// Options options
type Options struct {
	ConfigFile    string
	Version       string
	BuildInjector BuildInjector
}

// Option 定义配置项
type Option func(*Options)

// BuildInjector 构建注入器的方法
type BuildInjector func() (engine *gin.Engine, clean func(), err error)

// RunServer 启动服务
func RunServer(ctx context.Context, opts ...Option) (func(), error) {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}
	SetVersion(o.Version)
	// 加载配置文件
	config.MustLoad(o.ConfigFile)
	config.PrintWithJSON()

	// 启动日志
	logger.Printf(ctx, "HTTP Server startup, M[%s]-V[%s]-P[%d]", config.C.RunMode, o.Version, os.Getpid())

	// 初始化日志模块
	loggerCleanFunc, err := logger.InitLogger(ctx)
	if err != nil {
		return nil, err
	}

	// 初始化依赖注入器
	engine, injectorCleanFunc, err := o.BuildInjector()
	if err != nil {
		return nil, err
	}

	// 初始化HTTP服务
	shutdownServerFunc := RunHTTPServer(ctx, engine)

	return func() {
		shutdownServerFunc()
		injectorCleanFunc()
		loggerCleanFunc()
	}, nil
}
