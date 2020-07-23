package ser

import (
	"context"
	"errors"

	"github.com/suisrc/zgo/app/ent"

	dbx "github.com/suisrc/zgo/modules/db"
)

// User 用户
type User struct {
	db *ent.Client
}

// UpdateUser 更新用户信息
func (s *User) UpdateUser(ctx context.Context, body map[string]interface{}) (string, error) {
	ref := &dbx.ResultRef{}
	err := WithTx(ctx, s.db, func(tx *ent.Tx) error {
		// do nothing
		pms := body["name"]
		if v, ok := pms.(string); ok {
			ref.Data = v
			return nil
		}
		return errors.New("params is not string")
	})
	if err != nil {
		// 发生异常
		return "", err
	}
	return ref.Data.(string), nil
}
