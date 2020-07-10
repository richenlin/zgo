// Copyright 2020 Kratos Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

import "net/http"

//var WebFrameWorkSet = wire.NewSet(wire.Bind(new(engine.IEngine), new(*Engine)))

// IEngine web框架的接口
type IEngine interface {

	// Name web框架的名称
	Name() string

	// Use 绑定web框架中间件
	//Use(middleware ...func(*interface{})) (*IEngine, error)

	// Handler Web Handler
	RunHandler() http.Handler

	// Run 运行服务器
	Run(addr ...string) error
}
