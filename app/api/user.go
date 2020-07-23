package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/app/ser"
)

// User 接口
type User struct {
	SerUser *ser.User
}

// Hello godoc
// @Summary Hello
// @Description Hello world
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /test/hello [get]
func (a *User) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, world",
	})
}
