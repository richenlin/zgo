package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/modules/auth"
	"github.com/suisrc/zgo/modules/auth/jwt"
	"github.com/suisrc/zgo/modules/auth/jwt/store/buntdb"
	"github.com/suisrc/zgo/modules/config"
	"github.com/suisrc/zgo/modules/helper"
	"github.com/suisrc/zgo/modules/logger"
)

// Auth auth
type Auth struct {
	Enforcer *casbin.SyncedEnforcer
	Auther   auth.Auther
}

// RegisterWithUAC 注册路由,认证接口特殊,需要独立注册
func (a *Auth) RegisterWithUAC(r gin.IRouter) {
	uac := middleware.UserAuthCasbinMiddleware(a.Auther, a.Enforcer)

	r.GET("authz", uac, a.authorize)
	// r.GET(middleware.JoinPath(config.C.HTTP.ContextPath, "authz"), uac, a.authorize)
}

// Register 主路由必须包含UAC内容
func (a *Auth) Register(r gin.IRouter) {
	r.GET("authz", a.authorize)
}

// @Param Authorization header string true "Bearer token"

// Authorize godoc
// @Tags auth
// @Summary Authorize
// @Description 授权接口
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} helper.Success
// @Router /authz [get]
func (a *Auth) authorize(c *gin.Context) {
	// 权限判断有UserAuthCasbinMiddleware完成
	// 仅仅返回正常结果即可
	helper.ResSuccess(c, "ok")
}

// NewAuther of auth.Auther
// 注册认证使用的auther内容
func NewAuther() auth.Auther {
	store, err := buntdb.NewStore(":memory:") // 使用内存缓存
	if err != nil {
		panic(err)
	}
	secret := config.C.JWTAuth.SigningSecret
	if secret == "" {
		secret = auth.UUID(128)
		logger.Infof(nil, "jwt secret: %s", secret)
	}
	auther := jwt.New(store,
		jwt.SetSigningSecret(secret), // 注册令牌签名密钥
	)

	return auther
}
