package api

import (
	"strconv"
	"time"

	"github.com/suisrc/zgo/modules/logger"

	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/app/schema"
	"github.com/suisrc/zgo/modules/auth"
	"github.com/suisrc/zgo/modules/helper"
)

// Signin signin
type Signin struct {
	Auther auth.Auther
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
// @Param item body schema.SigninBody true "SigninBody Info"
// @Success 200 {object} helper.Success
// @Router /signin [post]
func (a *Signin) signin(c *gin.Context) {
	body := schema.SigninBody{}

	if err := helper.ParseJSON(c, &body); err != nil {
		logger.Infof(c, err.Error())
		return
	}

	user := &schema.SigninUser{}

	user.UserName = body.Username
	user.UserID = strconv.Itoa(1)
	user.RoleID = "basic"
	user.Issuer = "t.icgear.cn"
	user.Audience = "go.t.icgear.cn"

	token, err := a.Auther.GenerateToken(c, user)
	if err != nil {
		helper.ResError(c, &helper.Err401Unauthorized)
		return
	}

	result := schema.SigninResult{
		Status:  "ok",
		Token:   token.GetAccessToken(),
		Expired: token.GetExpiresAt() - time.Now().Unix(),
	}
	// 返回正常结果即可
	helper.ResSuccess(c, &result)
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
