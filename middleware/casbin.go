package middleware

import (
	"zgo/engine"
	"zgo/modules/config"
	"zgo/modules/result"

	"github.com/casbin/casbin/v2"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) engine.HandlerFunc {
	conf := config.C.Casbin
	if !conf.Enable {
		return EmptyMiddleware()
	}

	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		u, ok := c.GetUserInfo()
		if !ok {
			result.ResError(c, result.Err401Unauthorized)
			return
		}

		r := u.GetRoleID()
		p := c.RequestURLPath()
		m := c.RequestMethod()
		if b, err := enforcer.Enforce(r, p, m); err != nil {
			result.ResError(c, result.Err401Unauthorized)
			return
		} else if !b {
			result.ResError(c, result.Err401Unauthorized)
			return
		}
		c.Next()
	}
}
