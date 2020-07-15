package middleware

import (
	"zgo/engine"
	"zgo/modules/auth"
	"zgo/modules/config"
	"zgo/modules/result"

	"github.com/casbin/casbin/v2"
)

// UserAuthCasbinMiddleware 用户授权中间件
func UserAuthCasbinMiddleware(auther auth.Auther, enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) engine.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return EmptyMiddleware()
	}
	conf := config.C.Casbin

	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		u, err := auther.GetUserInfo(c)
		if err != nil {
			if err == auth.ErrInvalidToken || err == auth.ErrNoneToken {
				result.ResError(c, result.Err401Unauthorized)
				return
			}
			result.ResError(c, result.Err400BadRequest)
			return
		}

		if conf.Enable {
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
		}

		c.SetUserInfo(u)
		c.Next()
	}
}
