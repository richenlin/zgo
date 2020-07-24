package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suisrc/zgo/demo/model/ent"
	sqlxm "github.com/suisrc/zgo/demo/model/sqlx"
	"github.com/suisrc/zgo/demo/model/sqlxc"
	"github.com/suisrc/zgo/demo/schema"
	"github.com/suisrc/zgo/demo/service"
	"github.com/suisrc/zgo/modules/helper"
)

// Demo 接口
type Demo struct {
	GPA         service.GPA
	DemoService *service.Demo
}

// Register 注册路由
func (a *Demo) Register(r gin.IRouter) {
	r.GET("hello", a.Hello)
	r.GET("get", a.Get)
	r.POST("set", a.Set)
	r.GET("get1", a.Get1)
}

// Hello godoc
// @Summary Hello
// @Description Hello world
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /demo/hello [get]
func (a *Demo) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, world",
	})
}

// Set godoc
// @Summary Set
// @Description Set
// @Accept  json
// @Produce  json
// @Param item body schema.DemoSet false "Demo Info"
// @Success 200 {object} helper.ErrorInfo
// @@Failure 500 {object} helper.ErrorInfo
// @Router /demo/set [post]
func (a *Demo) Set(c *gin.Context) {
	item := &schema.DemoSet{}
	err := helper.ParseJSON(c, item)
	if err != nil {
		return
	}
	res, err := a.GPA.DBE.Demo.Create().
		SetCode(item.Code).
		SetName(item.Name).
		SetMemo(item.Memo).
		SetStatus(1).
		Save(c)

	if err != nil {
		panic(err)
	}
	helper.ResSuccess(c, res)
}

// Get godoc
// @Summary Get
// @Description Get
// @Accept  json
// @Produce  json
// @Param id query string true "Demo id"
// @Success 200 {object} helper.ErrorInfo
// @@Failure 500 {object} helper.ErrorInfo
// @Router /demo/get [get]
func (a *Demo) Get(c *gin.Context) {
	idstr := c.Query("id")
	if idstr == "" {
		helper.ResError(c, &helper.Err406NotAcceptable)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		helper.ResError(c, &helper.Err406NotAcceptable)
		return
	}
	res, err := a.GPA.DBE.Demo.Get(c, id)
	if err != nil && !ent.IsNotFound(err) {
		panic(err)
	}

	helper.ResSuccess(c, res)
}

// Get1 godoc
// @Summary Get
// @Description Get
// @Accept  json
// @Produce  json
// @Param id query string true "Demo id"
// @Success 200 {object} helper.ErrorInfo
// @@Failure 500 {object} helper.ErrorInfo
// @Router /demo/get1 [get]
func (a *Demo) Get1(c *gin.Context) {
	idstr := c.Query("id")
	if idstr == "" {
		helper.ResError(c, &helper.Err406NotAcceptable)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		helper.ResError(c, &helper.Err406NotAcceptable)
		return
	}
	res := &sqlxm.Demo{}
	err = a.GPA.DBS.Get(res, "SELECT id, code, name, memo FROM demo WHERE id=$1", id)
	if err != nil {
		if !sqlxc.IsNotFound(err) {
			panic(err)
		}
		res = nil // 没有找到数据
	}

	helper.ResSuccess(c, res)
}
