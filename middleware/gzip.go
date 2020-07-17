package middleware

import (
	"zgo/modules/config"

	"github.com/LyricTian/gzip"
	"github.com/gin-gonic/gin"
)

// GizMiddleware 跨域
func GizMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	conf := config.C.GZIP
	return gzip.Gzip(gzip.BestCompression,
		gzip.WithExcludedExtensions(conf.ExcludedExtentions),
		gzip.WithExcludedPaths(conf.ExcludedPaths),
	)
}
