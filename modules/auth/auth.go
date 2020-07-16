package auth

import (
	"errors"
	"strings"
	"zgo/modules/helper"

	"github.com/gin-gonic/gin"
)

// 定义错误
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrNoneToken    = errors.New("none token")
)

// TokenInfo 令牌信息
type TokenInfo interface {
	// 获取访问令牌
	GetToken() string
	// 获取令牌类型
	GetTokenType() string
	// 获取令牌到期时间戳
	GetExpiresAt() int64
	// JSON编码
	EncodeToJSON() ([]byte, error)
}

// UserInfo user
type UserInfo interface {
	helper.UserInfo
}

// Auther 认证接口
type Auther interface {
	// GetUserInfo 获取用户ID
	GetUserInfo(c *gin.Context) (UserInfo, error)

	// GetToken 获取用户令牌
	GetToken(c *gin.Context) string

	// GenerateToken 生成令牌
	GenerateToken(c *gin.Context, user UserInfo) (TokenInfo, error)

	// DestroyToken 销毁令牌
	DestroyToken(c *gin.Context, token string) error

	// ParseUserInfo 解析用户信息
	ParseUserInfo(c *gin.Context, token string) (UserInfo, error)

	// Release 释放资源
	Release() error
}

// GetBearerToken 获取用户令牌
func GetBearerToken(c *gin.Context) string {
	var token string

	prefix := "Bearer "
	if auth := c.GetHeader("Authorization"); auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetQueryToken 获取用户令牌
func GetQueryToken(c *gin.Context) string {
	var token string

	if auth, ok := c.GetQuery("token"); ok {
		token = auth
	}
	return token
}

// GetCookieToken 获取用户令牌
func GetCookieToken(c *gin.Context) string {
	var token string

	if auth, err := c.Cookie("authorization"); err != nil && auth != "" {
		token = auth
	}
	return token
}
