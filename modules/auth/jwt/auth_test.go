package jwt

import (
	"context"
	"testing"

	"github.com/suisrc/zgo/modules/auth/jwt/store/buntdb"
	"github.com/suisrc/zgo/modules/logger"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	store, err := buntdb.NewStore(":memory:")
	assert.Nil(t, err)
	// var store Storer

	var ref TokenRef

	jwtAuth := New(store, Option(func(o *options) {
		o.tokenFunc = func(ctx context.Context) (string, error) {
			return ref.ref, nil
		}
	}))

	defer jwtAuth.Release()

	ctx := context.Background()

	user := &UserInfo{
		UserName: "Json",
		UserID:   "123",
		RoleID:   "789",
		//TokenID:  "456",
		Issuer:   "abc.com",
		Audience: "def.com",
	}

	token, err := jwtAuth.GenerateToken(ctx, user)
	assert.Nil(t, err)
	assert.NotNil(t, token)

	data, err := token.EncodeToJSON()
	logger.Infof(ctx, "%s", string(data))

	ref.ref = token.GetAccessToken()
	uInfo, err := jwtAuth.GetUserInfo(ctx)
	assert.Nil(t, err)
	assert.Equal(t, user.UserID, uInfo.GetUserID())

	err = jwtAuth.DestroyToken(ctx, uInfo)
	assert.Nil(t, err)

	uInfo, err = jwtAuth.GetUserInfo(ctx)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid token")
	//assert.Empty(t, id)
}

type TokenRef struct {
	ref string
}

// UserInfo 用户信息声明
type UserInfo struct {
	UserName string
	UserID   string
	RoleID   string
	TokenID  string
	Issuer   string
	Audience string
}

// GetUserName name
func (u *UserInfo) GetUserName() string {
	return u.UserName
}

// GetUserID user
func (u *UserInfo) GetUserID() string {
	return u.UserID
}

// GetRoleID role
func (u *UserInfo) GetRoleID() string {
	return u.RoleID
}

// GetTokenID token
func (u *UserInfo) GetTokenID() string {
	return u.TokenID
}

// GetIssuer issuer
func (u *UserInfo) GetIssuer() string {
	return u.Issuer
}

// GetAudience audience
func (u *UserInfo) GetAudience() string {
	return u.Audience
}

// GetProps props
func (u *UserInfo) GetProps() (interface{}, bool) {
	return nil, false
}
