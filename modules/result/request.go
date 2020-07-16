package result

import (
	"zgo/engine"

	"github.com/gin-gonic/gin/binding"
)

// ParseJSON 解析请求JSON, 注意,解析失败后需要直接返回
func ParseJSON(c engine.Context, obj interface{}) error {
	if err := c.ShouldBind(obj, binding.JSON); err != nil {
		return ResError(c, Wrap400Response(c, err))
	}
	return nil
}

// ParseQuery 解析Query参数, 注意,解析失败后需要直接返回
func ParseQuery(c engine.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return ResError(c, Wrap400Response(c, err))
	}
	return nil
}

// ParseForm 解析Form请求, 注意,解析失败后需要直接返回
func ParseForm(c engine.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return ResError(c, Wrap400Response(c, err))
	}
	return nil
}
