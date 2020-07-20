package auth

import (
	"context"
	"errors"
)

// 定义错误
var (
	// ErrInvalidToken 无效令牌
	ErrInvalidToken = errors.New("invalid token")
	// ErrNoneToken 没有令牌
	ErrNoneToken = errors.New("none token")
)

// TokenInfo 令牌信息
type TokenInfo interface {
	// 获取访问令牌
	GetAccessToken() string
	// 获取令牌类型
	GetTokenType() string
	// 获取令牌到期时间戳
	GetExpiresAt() int64
	// JSON
	EncodeToJSON() ([]byte, error)
}

// UserInfo user
type UserInfo interface {
	//helper.UserInfo

	// GetUserName 用户名
	GetUserName() string

	// GetUserID 用户ID
	GetUserID() string
	// GetRoleID 角色ID
	GetRoleID() string
	// GetTokenID 令牌ID, 主要用于验证或者销毁令牌等关于令牌的操作
	GetTokenID() string

	// GetProps() 获取私有属性,该内容会被加密, 注意:内容敏感,不要存储太多的内容
	GetProps() (interface{}, bool)

	// 令牌签发者
	GetIssuer() string
	// 令牌接收者
	GetAudience() string
}

// Auther 认证接口
type Auther interface {
	// GetUserInfo 获取用户
	GetUserInfo(c context.Context) (UserInfo, error)

	// GenerateToken 生成令牌
	GenerateToken(c context.Context, u UserInfo) (TokenInfo, error)

	// DestroyToken 销毁令牌
	DestroyToken(c context.Context, u UserInfo) error
}
