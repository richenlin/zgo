// Copyright 2020 Kratos Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

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
)

// var WebFrameworkSet = wire.NewSet(InitWebFramework, wire.Bind(new(engine.IEngine), new(*WebFramework)))

// IEngine web框架的接口
type IEngine interface {
	IRouter

	// Name web框架的名称
	Name() string

	// Handler Web Handler
	RunHandler() http.Handler

	// Run 运行服务器
	Run(addr ...string) error
}

//====================================================分割线
//====================================================分割线
//====================================================分割线

// RunHTTPServer 初始化http服务
func RunHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

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
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
}

// Run 运行服务
func Run(ctx context.Context, runServer func() (func(), error)) error {
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
