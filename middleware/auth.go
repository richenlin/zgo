package middleware

import (
	"zgo/modules/auth"
	"zgo/modules/config"
	"zgo/modules/helper"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		user, err := a.GetUserInfo(c)
		if err != nil {
			if err == auth.ErrInvalidToken || err == auth.ErrNoneToken {
				helper.ResError(c, &helper.Err401Unauthorized)
				return
			}
			helper.ResError(c, &helper.Err400BadRequest)
			return
		}
		helper.SetUserInfo(c, user)

		c.Next()
	}
}
