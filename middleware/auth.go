package middleware

import (
	"zgo/engine"
	"zgo/modules/auth"
	"zgo/modules/config"
	"zgo/modules/result"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) engine.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return EmptyMiddleware()
	}

	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		user, err := a.GetUserInfo(c)
		if err != nil {
			if err == auth.ErrInvalidToken || err == auth.ErrNoneToken {
				result.ResError(c, result.Err401Unauthorized)
				return
			}
			result.ResError(c, result.Err400BadRequest)
			return
		}
		c.SetUserInfo(user)

		c.Next()
	}
}
