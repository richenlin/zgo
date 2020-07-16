package auth

import (
	"errors"
	"strings"
	"zgo/engine"
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

// UserInfo 用户信息
type UserInfo interface {
	engine.UserInfo
}

// Auther 认证接口
type Auther interface {
	// GetUserInfo 获取用户ID
	GetUserInfo(ctx engine.Context) (UserInfo, error)

	// GetToken 获取用户令牌
	GetToken(ctx engine.Context) string

	// GenerateToken 生成令牌
	GenerateToken(ctx engine.Context, user UserInfo) (TokenInfo, error)

	// DestroyToken 销毁令牌
	DestroyToken(ctx engine.Context, token string) error

	// ParseUserInfo 解析用户信息
	ParseUserInfo(ctx engine.Context, token string) (UserInfo, error)

	// Release 释放资源
	Release() error
}

// GetBearerToken 获取用户令牌
func GetBearerToken(c engine.Context) string {
	var token string

	prefix := "Bearer "
	if auth := c.GetHeader("Authorization"); auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetQueryToke 获取用户令牌
func GetQueryToke(c engine.Context) string {
	var token string

	if auth := c.GetHeader("Authorization"); auth != "" {
		token = auth
	}
	return token
}

// GetCookieToke 获取用户令牌
func GetCookieToke(c engine.Context) string {
	var token string

	if auth := c.GetHeader("Authorization"); auth != "" {
		token = auth
	}
	return token
}
