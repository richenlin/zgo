package schema

import "github.com/suisrc/zgo/modules/auth"

// SigninBody 登陆参数
type SigninBody struct {
	Username string `json:"usrname" binding:"required"` // 账户
	Password string `json:"password"`                   // 密码
	Mobile   string `json:"mobile"`                     // 手机号
	Captcha  string `json:"captcha"`                    // 验证码
	Code     string `json:"code"`                       // 登陆方式
	Type     string `json:"type"`                       // 登陆类型 <系统>:<类型>:<备注>
	Role     string `json:"role"`                       // 角色
	Reset    bool   `json:"reset"`                      // 重置登陆

}

// SigninResult 登陆返回值
type SigninResult struct {
	Status       string        `json:"status" default:"ok"`    // 'ok' | 'error' 不适用boolean类型是为了以后可以增加扩展
	Token        string        `json:"token,omitempty"`        // 令牌
	Expired      int64         `json:"expired,omitempty"`      // 过期时间
	RefreshToken string        `json:"refreshToken,omitempty"` // 刷新令牌
	Message      string        `json:"message,omitempty"`      // 消息,有限显示
	Roles        []interface{} `json:"roles,omitempty"`        // 多角色的时候，返回角色，重新确认登录
}

var _ auth.UserInfo = &SigninUser{}

// SigninUser 登陆用户信息
type SigninUser struct {
	UserName string
	UserID   string
	RoleID   string
	TokenID  string
	Issuer   string
	Audience string
}

// GetUserName 用户名
func (s *SigninUser) GetUserName() string {
	return s.UserName
}

// GetUserID 用户ID
func (s *SigninUser) GetUserID() string {
	return s.UserID
}

// GetRoleID 角色ID
func (s *SigninUser) GetRoleID() string {
	return s.RoleID
}

// GetTokenID 令牌ID, 主要用于验证或者销毁令牌等关于令牌的操作
func (s *SigninUser) GetTokenID() string {
	return s.TokenID
}

// GetProps 获取私有属性,该内容会被加密, 注意:内容敏感,不要存储太多的内容
func (s *SigninUser) GetProps() (interface{}, bool) {
	return nil, false
}

// GetIssuer 令牌签发者
func (s *SigninUser) GetIssuer() string {
	return s.Issuer
}

// GetAudience 令牌接收者
func (s *SigninUser) GetAudience() string {
	return s.Audience
}
