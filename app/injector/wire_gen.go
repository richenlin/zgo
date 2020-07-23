// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/suisrc/zgo/app/api"
	"github.com/suisrc/zgo/app/model/entc"
	"github.com/suisrc/zgo/app/model/sqlxc"
	"github.com/suisrc/zgo/app/service"
	"github.com/suisrc/zgo/middlewire"
	"github.com/suisrc/zgo/modules/casbin"
	"github.com/suisrc/zgo/modules/casbin/adapter/json"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	engine := middlewire.InitGinEngine()
	adapter, err := casbinjson.NewCasbinAdapter()
	if err != nil {
		return nil, nil, err
	}
	syncedEnforcer, cleanup, err := casbin.NewCasbinEnforcer(adapter)
	if err != nil {
		return nil, nil, err
	}
	router := middlewire.NewRouter(engine)
	client, cleanup2, err := entc.NewClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	db, cleanup3, err := sqlxc.NewClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gpa := service.GPA{
		DBE: client,
		DBS: db,
	}
	demo := &service.Demo{
		GPA: gpa,
	}
	apiDemo := &api.Demo{
		DemoService: demo,
	}
	options := &api.Options{
		Engine:   engine,
		Enforcer: syncedEnforcer,
		Router:   router,
		Demo:     apiDemo,
	}
	endpoints := api.InitEndpoints(options)
	swagger := middlewire.NewSwagger(engine)
	healthz := middlewire.NewHealthz(engine)
	injector := &Injector{
		Engine:    engine,
		Endpoints: endpoints,
		Swagger:   swagger,
		Healthz:   healthz,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"), middlewire.NewSwagger, middlewire.NewHealthz)

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine    *gin.Engine
	Endpoints *api.Endpoints
	Swagger   middlewire.Swagger
	Healthz   middlewire.Healthz
}
