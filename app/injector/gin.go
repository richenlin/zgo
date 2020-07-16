package injector

import (
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
)

// InitGinEngine engine
func InitGinEngine() *gin.Engine {
	gin.SetMode(config.C.RunMode)
	//gin.SetMode(gin.DebugMode)

	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	return app
}
