package api

import (
	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/modules/helper"
)

// Signin signin
type Signin struct {
	//Enforcer *casbin.SyncedEnforcer
	//Auther   auth.Auther
}

// Register 注册路由,认证接口特殊,需要独立注册
func (a *Signin) Register(r gin.IRouter) {
	// sign 开头的路由会被全局casbin放行
	r.POST("signin", a.signin) // 登陆必须是POST请求
	r.GET("signout", a.signout)
}

// Signin godoc
// @Tags sign
// @Summary Signin
// @Description 登陆
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.Success
// @Router /signin [post]
func (a *Signin) signin(c *gin.Context) {
	// 返回正常结果即可
	helper.ResSuccess(c, "ok")
}

// Signout godoc
// @Tags sign
// @Summary Signin
// @Description 登出
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.Success
// @Router /signout [get]
func (a *Signin) signout(c *gin.Context) {
	// 返回正常结果即可
	helper.ResSuccess(c, "ok")
}
