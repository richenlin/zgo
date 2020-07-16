package middleware

import (
	"os"
	"path/filepath"
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
)

// WWWMiddleware 静态站点中间件
func WWWMiddleware(root string, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		if root == "" {
			root = config.C.WWW.Dir
		}

		p := c.Request.URL.Path
		fpath := filepath.Join(root, filepath.FromSlash(p))
		_, err := os.Stat(fpath)
		if err != nil && os.IsNotExist(err) {
			fpath = filepath.Join(root, config.C.WWW.Index)
		}

		c.File(fpath)
		c.Abort()
	}
}
