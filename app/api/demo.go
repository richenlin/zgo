package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/app/ser"
)

// Demo 接口
type Demo struct {
	SerDemo *ser.Demo
}

// Register 注册路由
func (a *Demo) Register(r gin.IRouter) {
	r.GET("hello", a.Hello)
}

// Hello godoc
// @Summary Hello
// @Description Hello world
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /test/hello [get]
func (a *Demo) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, world",
	})
}
