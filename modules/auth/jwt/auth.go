package jwt

/*
 为什么使用反向验证(只记录登出的用户, 因为我们确信点击登出的操作比点击登陆的操作要少的多的多)
*/
import (
	"context"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/suisrc/zgo/modules/auth"
	"github.com/suisrc/zgo/modules/logger"
)

type options struct {
	tokenType     string                                                                // 令牌类型,传递给TokenInfo
	expired       int                                                                   // 过期间隔
	signingMethod jwt.SigningMethod                                                     // 签名方法
	signingSecret interface{}                                                           // 签名密钥
	keyFunc       func(*jwt.Token, jwt.SigningMethod, interface{}) (interface{}, error) // JWT中获取密钥, 该内容可以忽略默认的signingMethod和signingSecret
	claimsFunc    func(jwt.Claims, jwt.SigningMethod) (*jwt.Token, error)               // JWT构建令牌, 该内容可以忽略默认的signingMethod
	tokenFunc     func(context.Context) (string, error)                                 // 获取令牌
}

// Option 定义参数项
type Option func(*options)

// SetSigningMethod 设定签名方式
func SetSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// SetSigningSecret 设定签名方式
func SetSigningSecret(secret string) Option {
	return func(o *options) {
		o.signingSecret = []byte(secret)
	}
}

// SetExpired 设定令牌过期时长(单位秒，默认7200)
func SetExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// SetKeyFunc 设定签名key
func SetKeyFunc(f func(*jwt.Token, jwt.SigningMethod, interface{}) (interface{}, error)) Option {
	return func(o *options) {
		o.keyFunc = f
	}
}

// SetNewClaims 设定声明内容
func SetNewClaims(f func(jwt.Claims, jwt.SigningMethod) (*jwt.Token, error)) Option {
	return func(o *options) {
		o.claimsFunc = f
	}
}

// SetTokenFunc 设定令牌Token
func SetTokenFunc(f func(context.Context) (string, error)) Option {
	return func(o *options) {
		o.tokenFunc = f
	}
}

//===================================================
// 分割线
//===================================================

// New 创建认证实例
func New(store Storer, opts ...Option) *Auther {
	o := options{
		tokenType:     "JWT",
		expired:       7200,
		signingMethod: jwt.SigningMethodHS512,
		keyFunc:       KeyFuncCallback,
		claimsFunc:    NewWithClaims,
		tokenFunc:     GetBearerToken,
	}
	for _, opt := range opts {
		opt(&o)
	}
	if o.signingSecret == nil {
		o.signingSecret = []byte(NewRandomID()) // 默认随机生成
		logger.Infof(nil, "new random signing secret: %s", o.signingSecret)
	}

	return &Auther{
		opts:  &o,
		store: store,
	}
}

// Release 释放资源
func (a *Auther) Release() error {
	return a.callStore(func(store Storer) error {
		return store.Close()
	})
}

//===================================================
// 分割线
//===================================================

var _ auth.Auther = &Auther{}

// Auther jwt认证
type Auther struct {
	opts  *options
	store Storer
}

// GetUserInfo 获取用户
func (a *Auther) GetUserInfo(c context.Context) (auth.UserInfo, error) {
	tokenString, err := a.opts.tokenFunc(c)
	if err != nil {
		return nil, err
	}

	claims, err := a.parseToken(tokenString)
	if err != nil {
		var e *jwt.ValidationError
		if errors.As(err, &e) {
			return nil, auth.ErrInvalidToken
		}
		return nil, err
	}

	err = a.callStore(func(store Storer) error {
		// 反向验证该用户是否已经登出
		if exists, err := store.Check(c, claims.GetTokenID()); err != nil {
			return err
		} else if exists {
			return auth.ErrInvalidToken
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// GenerateToken 生成令牌
func (a *Auther) GenerateToken(c context.Context, user auth.UserInfo) (auth.TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expired) * time.Second).Unix()
	issuedAt := now.Unix()

	claims := NewUserInfo(user)
	claims.IssuedAt = issuedAt
	claims.NotBefore = issuedAt
	claims.ExpiresAt = expiresAt

	token, err := a.opts.claimsFunc(claims, a.opts.signingMethod)
	if err != nil {
		return nil, err
	}

	tokenString, err := token.SignedString(a.opts.signingSecret)
	if err != nil {
		return nil, err
	}

	tokenInfo := &TokenInfo{
		AccessToken: tokenString,
		TokenStatus: "ok",
		TokenType:   a.opts.tokenType,
		ExpiresAt:   expiresAt,
	}
	return tokenInfo, nil
}

// DestroyToken 销毁令牌
func (a *Auther) DestroyToken(c context.Context, user auth.UserInfo) error {
	claims, ok := user.(*UserClaims)
	if !ok {
		return auth.ErrInvalidToken
	}

	// 如果设定了存储，则将未过期的令牌放入
	return a.callStore(func(store Storer) error {
		expired := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())
		return store.Set(c, claims.GetTokenID(), expired)
	})
}

//===================================================
// 分割线
//===================================================

// 解析令牌
func (a *Auther) parseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, a.keyFunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, auth.ErrInvalidToken
	}

	return token.Claims.(*UserClaims), nil
}

// 获取密钥
func (a *Auther) keyFunc(t *jwt.Token) (interface{}, error) {
	return a.opts.keyFunc(t, a.opts.signingMethod, a.opts.signingSecret)
}

// 调用存储方法
func (a *Auther) callStore(fn func(Storer) error) error {
	if store := a.store; store != nil {
		return fn(store)
	}
	return nil
}

//===================================================
// 分割线
//===================================================

// KeyFuncCallback 解析方法使用此回调函数来提供验证密钥。
// 该函数接收解析后的内容，但未验证的令牌。这使您可以在令牌的标头（例如 kid），以标识要使用的密钥。
func KeyFuncCallback(token *jwt.Token, method jwt.SigningMethod, secret interface{}) (interface{}, error) {
	//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//	return nil, auth.ErrInvalidToken // 无法验证
	//}
	//kid := token.Header["kid"]
	//if kid == "" {
	//	return nil, auth.ErrInvalidToken // 无法验证
	//}
	token.Method = method // 强制使用配置, 防止alg使用none而跳过验证
	return secret, nil
}

// NewWithClaims new claims
// jwt.NewWithClaims
func NewWithClaims(claims jwt.Claims, method jwt.SigningMethod) (*jwt.Token, error) {
	return &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
			// "kid": "zgo123456",
		},
		Claims: claims,
		Method: method,
	}, nil
}

// NewRandomID new ID
func NewRandomID() string {
	// uuid, err := uuid.NewRandom()
	// if err != nil {
	// 	panic(err)
	// }
	// strid := uuid.String()
	// return strid
	return auth.UUID(32)
}
