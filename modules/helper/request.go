package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ParseJSON 解析请求JSON, 注意,解析失败后需要直接返回
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return ResError(c, Wrap400Response(c, err))
	}
	return nil
}

// ParseQuery 解析Query参数, 注意,解析失败后需要直接返回
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return ResError(c, Wrap400Response(c, err))
	}
	return nil
}

// ParseForm 解析Form请求, 注意,解析失败后需要直接返回
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return ResError(c, Wrap400Response(c, err))
	}
	return nil
}
