package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/modules/auth"
	"github.com/suisrc/zgo/modules/auth/jwt"
	"github.com/suisrc/zgo/modules/auth/jwt/store/buntdb"
	"github.com/suisrc/zgo/modules/helper"
)

// Auth auth
type Auth struct {
	Enforcer *casbin.SyncedEnforcer
	Auther   auth.Auther
}

// Register 注册路由,认证接口特殊,需要独立注册
func (a *Auth) Register(r gin.IRouter) {
	uac := middleware.UserAuthCasbinMiddleware(a.Auther, a.Enforcer)
	r.GET("authz", uac, a.authorize)
}

// Authorize godoc
// @Tags auth
// @Summary Authorize
// @Description 授权接口
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.Success
// @Router /auth [get]
func (a *Auth) authorize(c *gin.Context) {
	// 返回正常结果即可
	helper.ResSuccess(c, "ok")
}

// NewAuther of auth.Auther
// 注册认证使用的auther内容
func NewAuther() auth.Auther {
	store, err := buntdb.NewStore(":memory:") // 使用内存缓存
	if err != nil {
		panic(err)
	}
	auther := jwt.New(store,
		jwt.SetSigningSecret("12345678"), // 注册令牌签名密钥
	)

	return auther
}
