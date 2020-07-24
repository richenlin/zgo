package api

import (
	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/modules/helper"
)

// User 用户管理器
type User struct {
}

// Register 注册接口
func (a *User) Register(r gin.IRouter) {
	user := r.Group("user")

	user.GET("hello", a.hello)
}

// hello godoc
// @Tags user
// @Summary hello
// @Description 用户接口测试
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.Success
// @Router /user/hello [get]
func (a *User) hello(c *gin.Context) {
	// 返回正常结果即可
	helper.ResSuccess(c, "ok")
}
