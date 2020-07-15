// Copyright 2020 Kratos Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

import (
	"net/http"
)

// var WebFrameworkSet = wire.NewSet(InitWebFramework, wire.Bind(new(engine.IEngine), new(*WebFramework)))

// IEngine web框架的接口
type IEngine interface {
	IRouter

	// NoMethod 未匹配到方法
	NoMethod(...HandlerFunc)

	// NoRoute 未匹配到路由
	NoRoute(...HandlerFunc)

	// Target target
	Target() interface{}

	// Name web框架的名称
	Name() string

	// Handler Web Handler
	RunHandler() http.Handler

	// Run 运行服务器
	Run(addr ...string) error
}
