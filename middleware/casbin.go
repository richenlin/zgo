package middleware

import (
	"github.com/suisrc/zgo/modules/config"
	"github.com/suisrc/zgo/modules/helper"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	conf := config.C.Casbin
	if !conf.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		u, ok := helper.GetUserInfo(c)
		if !ok {
			helper.ResError(c, &helper.Err401Unauthorized)
			return
		}

		r := u.GetRoleID()
		p := c.Request.URL.Path
		m := c.Request.Method
		if b, err := enforcer.Enforce(r, p, m); err != nil {
			helper.ResError(c, &helper.Err401Unauthorized)
			return
		} else if !b {
			helper.ResError(c, &helper.Err401Unauthorized)
			return
		}
		c.Next()
	}
}
