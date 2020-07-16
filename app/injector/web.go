package injector

import (
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InitWebSet 注入到wire中
var InitWebSet = wire.NewSet(wire.Struct(new(InitWebResultOptions), "*"), InitGinEngine)

// InitGinEngine engine
func InitGinEngine() *gin.Engine {
	gin.SetMode(config.C.RunMode)
	//gin.SetMode(gin.DebugMode)

	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	return app
}
