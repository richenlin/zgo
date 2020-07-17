package middleware

import (
	"time"
	"github.com/suisrc/zgo/modules/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware 跨域
func CORSMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	conf := config.C.CORS
	return cors.New(cors.Config{
		AllowOrigins:     conf.AllowOrigins,
		AllowMethods:     conf.AllowMethods,
		AllowHeaders:     conf.AllowHeaders,
		AllowCredentials: conf.AllowCredentials,
		MaxAge:           time.Second * time.Duration(conf.MaxAge),
	})
}
