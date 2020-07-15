package gin

import (
	"time"
	"zgo/modules/config"

	"github.com/LyricTian/gzip"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// MiddlewareCORS 跨域请求中间件
func MiddlewareCORS() gin.HandlerFunc {
	cfg := config.C.CORS
	return cors.New(cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     cfg.AllowMethods,
		AllowHeaders:     cfg.AllowHeaders,
		AllowCredentials: cfg.AllowCredentials,
		MaxAge:           time.Second * time.Duration(cfg.MaxAge),
	})
}

// RoutesSwagger 跨域请求中间件
func RoutesSwagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}

// Middlewares swagger
type Middlewares struct {
}

// UseMiddlewares UseMiddlewares
func UseMiddlewares(app *gin.Engine) {

	// DefaultMiddle
	if config.C.Middle.Logger {
		app.Use(gin.Logger())
	}
	if config.C.Middle.Recover {
		app.Use(gin.Recovery())
	}

	// Swagger
	if config.C.Swagger {
		app.GET("/swagger/*any", RoutesSwagger())
	}

	// CORS
	if config.C.CORS.Enable {
		app.Use(MiddlewareCORS())
	}

	// GZIP
	if config.C.GZIP.Enable {
		app.Use(gzip.Gzip(gzip.BestCompression,
			gzip.WithExcludedExtensions(config.C.GZIP.ExcludedExtentions),
			gzip.WithExcludedPaths(config.C.GZIP.ExcludedPaths),
		))
	}
}
