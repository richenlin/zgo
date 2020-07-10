// Copyright 2020 Kratos Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"zgo/engine"
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// WebFrameWorkSet 注入wire
var WebFrameWorkSet = wire.NewSet(InitWebFrameWork, wire.Bind(new(engine.IEngine), new(*WebFrameWork)))

// WebFrameWork web框架
type WebFrameWork struct {
	engine *gin.Engine
}

// InitWebFrameWork 初始化web框架
func InitWebFrameWork() (*WebFrameWork, error) {
	gin.SetMode(config.C.RunMode)
	//gin.SetMode(gin.DebugMode)

	app := gin.New()
	return &WebFrameWork{
		engine: app,
	}, nil
}

// Name web框架的名称
func (wf *WebFrameWork) Name() string {
	return "gin"
}

// Use web框架中间件
func (wf *WebFrameWork) Use(middleware ...gin.HandlerFunc) (*WebFrameWork, error) {
	wf.engine.Use(middleware...)
	return wf, nil
}

// Run 运行服务器
func (wf *WebFrameWork) Run(addr ...string) error {
	return wf.engine.Run(addr...)
}

// RunHandler 获取服务器句柄
func (wf *WebFrameWork) RunHandler() http.Handler {
	return wf.engine
}
