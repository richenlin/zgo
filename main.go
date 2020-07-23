package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/middlewire"
	cmd "github.com/suisrc/zgo/modules/app"
	"github.com/suisrc/zgo/modules/config"
	// // 引入swagger
	// _ "github.com/suisrc/zgo/app/swagger"
)

func main() {
	a := newApp()
	middlewire.NewHealthz(a)

	r := newRouter(a)
	c, err := router(r)

	if err != nil {
		return
	}
	// middlewire.NewSwagger(a)
	runApp(a, c)
}

func init() {
	//log.Println("=================init=================")
	config.MustLoad("zgo.toml")
}

func router(router gin.IRouter) (func(), error) {

	return func() {}, nil
}

func newRouter(app *gin.Engine) gin.IRouter {
	var r gin.IRouter
	if v := config.C.HTTP.ContextPath; v != "" {
		r = app.Group(v)
	} else {
		r = app
	}
	return r
}

func newApp() *gin.Engine {
	gin.SetMode(config.C.RunMode)
	//gin.SetMode(gin.DebugMode)

	app := gin.New()
	//app := gin.Default()

	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	app.Use(gin.Logger())
	//app.Use(middleware.LoggerMiddleware())

	app.Use(gin.Recovery())
	app.Use(middleware.RecoveryMiddleware())

	return app
}

func runApp(a *gin.Engine, clean func()) {
	ctx := context.Background()
	err := cmd.RunWithShutdown(ctx, func() (func(), error) {
		//a.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
		shutdownServerFunc := cmd.RunHTTPServer(ctx, a)
		return func() {
			shutdownServerFunc()
			clean()
		}, nil
	})
	if err != nil {
		panic(err)
	}
}
