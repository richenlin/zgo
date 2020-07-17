package middlewire

// 系统默认根路由
import (
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/modules/config"

	"github.com/gin-gonic/gin"
)

// Router 根路由 Register???
type Router gin.IRouter

// NewRouter 初始化根路由
func NewRouter(app *gin.Engine) Router {

	var router Router
	if v := config.C.HTTP.ContextPath; v != "" {
		router = app.Group(v)
	} else {
		router = app
	}

	// CORS
	if config.C.CORS.Enable {
		router.Use(middleware.CORSMiddleware())
	}
	// GZIP
	if config.C.GZIP.Enable {
		router.Use(middleware.GizMiddleware())
	}

	return router
}
