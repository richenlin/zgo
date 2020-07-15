package middleware

import (
	"os"
	"path/filepath"
	"zgo/engine"
	"zgo/modules/config"
)

// WWWMiddleware 静态站点中间件
func WWWMiddleware(root string, skippers ...SkipperFunc) engine.HandlerFunc {
	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		if root == "" {
			root = config.C.WWW.Dir
		}

		p := c.RequestURLPath()
		fpath := filepath.Join(root, filepath.FromSlash(p))
		_, err := os.Stat(fpath)
		if err != nil && os.IsNotExist(err) {
			fpath = filepath.Join(root, config.C.WWW.Index)
		}

		c.File(fpath)
		c.Abort()
	}
}
