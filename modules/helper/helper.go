package helper

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserInfo 用户
func GetUserInfo(c *gin.Context) (UserInfo, bool) {
	if v, ok := c.Get(UserInfoKey); ok {
		return v.(UserInfo), true
	}
	return nil, false
}

// SetUserInfo 用户
func SetUserInfo(c *gin.Context, user UserInfo) {
	c.Set(UserInfoKey, user)
}

// GetTraceID 根据山下问,获取追踪ID
func GetTraceID(c *gin.Context) string {
	if v, ok := c.Get(TraceIDKey); ok && v != "" {
		return v.(string)
	}

	// 优先从请求头中获取请求ID
	traceID := c.GetHeader("X-Request-Id")
	if traceID == "" {
		// 没有自建
		v, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}
		traceID = v.String()
	}
	c.Set(TraceIDKey, traceID)
	return traceID
}

// GetClientIP 获取客户端IP
func GetClientIP(c *gin.Context) string {
	if v, err := c.Cookie("X-Forwarded-For"); err != nil && v != "" {
		log.Println(v)
		len := strings.Index(v, ",")
		if len < 0 {
			return v
		}
		return v[:len]
	}
	return c.ClientIP()
}
