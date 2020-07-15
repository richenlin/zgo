package injector

import (
	"zgo/engine"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InitWebSet 注入到wire中
var InitWebSet = wire.NewSet(wire.Struct(new(InitWebResultOptions), "*"), InitWebFunc)

// InitWebResult engine
type InitWebResult struct {
}

// InitWebResultOptions options
type InitWebResultOptions struct {
	eng engine.IEngine
}

// InitWebFunc engine
func InitWebFunc(opts *InitWebResultOptions) *InitWebResult {
	app := opts.eng.Target().(*gin.Engine) // 如果确认使用了gin引擎,可以直接使用
	app.Use(gin.Logger(), gin.Recovery())

	return &InitWebResult{}
}
