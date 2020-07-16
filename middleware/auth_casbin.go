package middleware

import (
	"zgo/modules/auth"
	"zgo/modules/config"
	"zgo/modules/helper"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// UserAuthCasbinMiddleware 用户授权中间件
func UserAuthCasbinMiddleware(auther auth.Auther, enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return EmptyMiddleware()
	}
	conf := config.C.Casbin

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		user, err := auther.GetUserInfo(c)
		if err != nil {
			if err == auth.ErrInvalidToken || err == auth.ErrNoneToken {
				helper.ResError(c, &helper.Err401Unauthorized)
				return
			}
			helper.ResError(c, &helper.Err400BadRequest)
			return
		}

		if conf.Enable {
			r := user.GetRoleID()
			p := c.Request.URL.Path
			m := c.Request.Method
			if b, err := enforcer.Enforce(r, p, m); err != nil {
				helper.ResError(c, &helper.Err401Unauthorized)
				return
			} else if !b {
				helper.ResError(c, &helper.Err401Unauthorized)
				return
			}
		}

		helper.SetUserInfo(c, user)
		c.Next()
	}
}
