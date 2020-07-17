package middlewire

// 对gin进行初始化
import (
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/modules/config"

	"github.com/gin-gonic/gin"
)

// InitGinEngine engine
func InitGinEngine() *gin.Engine {
	gin.SetMode(config.C.RunMode)
	//gin.SetMode(gin.DebugMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	app.Use(gin.Logger())
	//app.Use(middleware.LoggerMiddleware())

	app.Use(gin.Recovery())
	app.Use(middleware.RecoveryMiddleware())

	return app
}
