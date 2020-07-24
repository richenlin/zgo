package service

import (
	"context"

	"github.com/suisrc/zgo/demo/model/ent"
	"github.com/suisrc/zgo/demo/model/entc"
)

// Demo 用户
type Demo struct {
	GPA
}

// T1WithTx 更新用户信息
func (s *Demo) T1WithTx(ctx context.Context) (string, error) {
	res, err := entc.WithTxV(ctx, s.DBE, func(tx *ent.Tx) (interface{}, error) {
		return "ok", nil
	})
	if err != nil {
		return "", err
	}
	return res.(string), nil
}

// T9WithTx 更新用户信息
func (s *Demo) T9WithTx(ctx context.Context, body map[string]interface{}) (string, error) {
	ref := &ResultRef{}
	err := entc.WithTx(ctx, s.DBE, func(tx *ent.Tx) error {
		ref.D = "ok"
		return nil
	})
	if err != nil {
		return "", err
	}
	return ref.D.(string), nil
}
