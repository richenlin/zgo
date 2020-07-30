package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/suisrc/zgo/modules/auth"
)

// NewUserInfo 获取用户信息
func NewUserInfo(user auth.UserInfo) *UserClaims {
	claims := UserClaims{}

	tokenID := user.GetTokenID()
	if tokenID == "" {
		tokenID = NewRandomID()
	}

	claims.Id = tokenID
	claims.Subject = user.GetUserID()
	claims.Name = user.GetUserName()
	claims.Role = user.GetRoleID()

	claims.Issuer = user.GetIssuer()
	claims.Audience = user.GetAudience()

	return &claims
}

var _ auth.UserInfo = &UserClaims{}

// UserClaims 用户信息声明
type UserClaims struct {
	jwt.StandardClaims
	Name       string      `json:"nam,omitempty"` // 用户名
	Role       string      `json:"rol,omitempty"` // 角色ID, role id
	Properties interface{} `json:"pps,omitempty"` // 用户的额外属性
}

// GetUserName name
func (u *UserClaims) GetUserName() string {
	return u.Name
}

// GetUserID user
func (u *UserClaims) GetUserID() string {
	return u.Subject
}

// GetRoleID role
func (u *UserClaims) GetRoleID() string {
	return u.Role
}

// GetTokenID token
func (u *UserClaims) GetTokenID() string {
	return u.Id
}

// GetIssuer issuer
func (u *UserClaims) GetIssuer() string {
	return u.Issuer
}

// GetAudience audience
func (u *UserClaims) GetAudience() string {
	return u.Audience
}

// GetProps props
func (u *UserClaims) GetProps() (interface{}, bool) {
	return nil, false
}
