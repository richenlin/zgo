package jwt

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/suisrc/zgo/modules/auth"
)

var _ auth.TokenInfo = &TokenInfo{}

// TokenInfo 令牌信息
type TokenInfo struct {
	AccessToken string `json:"token,omitempty"`   // 访问令牌
	TokenStatus string `json:"status,omitempty"`  // 令牌状态, ok | error
	TokenType   string `json:"type,omitempty"`    // 令牌类型
	ExpiresAt   int64  `json:"expired,omitempty"` // 令牌到期时间
	ErrMessage  string `json:"message,omitempty"` // 令牌异常原因
}

// GetAccessToken access token
func (t *TokenInfo) GetAccessToken() string {
	return t.AccessToken
}

// GetTokenStatus token status
func (t *TokenInfo) GetTokenStatus() string {
	return t.TokenStatus
}

// GetTokenType token type
func (t *TokenInfo) GetTokenType() string {
	return t.TokenType
}

// GetExpiresAt expires at
func (t *TokenInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

// EncodeToJSON to json
func (t *TokenInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}

//=================================================
// 分割线
//=================================================

// GetBearerToken 获取用户令牌
func GetBearerToken(ctx context.Context) (string, error) {
	if c, ok := ctx.(*gin.Context); ok {
		prefix := "Bearer "
		if auth := c.GetHeader("Authorization"); auth != "" && strings.HasPrefix(auth, prefix) {
			return auth[len(prefix):], nil
		}
	}

	return "", auth.ErrNoneToken
}

// GetQueryToken 获取用户令牌
func GetQueryToken(ctx context.Context) (string, error) {
	if c, ok := ctx.(*gin.Context); ok {
		if auth, ok := c.GetQuery("token"); ok && auth != "" {
			return auth, nil
		}
	}

	return "", auth.ErrNoneToken
}

// GetCookieToken 获取用户令牌
func GetCookieToken(ctx context.Context) (string, error) {
	if c, ok := ctx.(*gin.Context); ok {
		if auth, err := c.Cookie("authorization"); err == nil && auth != "" {
			return auth, nil
		}
	}

	return "", auth.ErrNoneToken
}
