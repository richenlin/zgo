package middleware

import (
	"github.com/suisrc/zgo/modules/config"

	"github.com/LyricTian/gzip"
	"github.com/gin-gonic/gin"
)

// GizMiddleware Giz, 主要部署前端时候(www中间件)对静态资源进行压缩
func GizMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	conf := config.C.GZIP
	return gzip.Gzip(gzip.BestCompression,
		gzip.WithExcludedExtensions(conf.ExcludedExtentions),
		gzip.WithExcludedPaths(conf.ExcludedPaths),
	)
}
