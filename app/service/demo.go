package service

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/suisrc/zgo/app/model/ent"
	"github.com/suisrc/zgo/app/model/entc"
	"github.com/suisrc/zgo/app/model/sqlxc"
)

// Demo 用户
type Demo struct {
	GPA
}

// T1WithTx 更新用户信息
func (s *Demo) T1WithTx(ctx context.Context, body map[string]interface{}) (string, error) {
	res, err := entc.WithTxV(ctx, s.DBE, func(tx *ent.Tx) (interface{}, error) {
		pname := body["name"]
		if v, ok := pname.(string); ok {
			return v, nil
		}
		return nil, errors.New("params has not name")
	})
	if err != nil {
		return "", err
	}
	return res.(string), nil
}

// T2WithTx 更新用户信息
func (s *Demo) T2WithTx(ctx context.Context, body map[string]interface{}) (string, error) {
	res, err := sqlxc.WithTxV(ctx, s.DBS, func(tx *sqlx.Tx) (interface{}, error) {
		pname := body["name"]
		if v, ok := pname.(string); ok {
			return v, nil
		}
		return nil, errors.New("params has not name")
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
		pname := body["name"]
		if v, ok := pname.(string); ok {
			ref.Data = v
			return nil
		}
		return errors.New("params has not name")
	})
	if err != nil {
		// 发生异常
		return "", err
	}
	return ref.Data.(string), nil
}
