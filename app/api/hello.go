package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello 接口
type Hello struct {
}

// Hello godoc
// @Summary Hello
// @Description Hello world
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /test/hello [get]
func (a *Hello) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, world",
	})
}
