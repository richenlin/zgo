package entc

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suisrc/zgo/app/model/ent"
)

// WithTx 执行带有事务的方法, 在一个事务中完成所有的内容
func WithTx(ctx context.Context, db *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

// WithTxV 执行带有事务的方法, 在一个事务中完成所有的内容
func WithTxV(ctx context.Context, db *ent.Client, fn func(*ent.Tx) (interface{}, error)) (interface{}, error) {
	fnx := func(tx *ent.Tx, rr interface{}) (interface{}, error) {
		return fn(tx)
	}
	return WithTxVx(ctx, db, nil, fnx)
}

// WithTxVx 执行带有事务的方法, 在一个事务中完成所有的内容
func WithTxVx(ctx context.Context, db *ent.Client, rr interface{}, fn func(*ent.Tx, interface{}) (interface{}, error)) (interface{}, error) {
	tx, err := db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := recover(); err != nil {
			// 发生中断异常
			tx.Rollback()
			panic(err)
		}
	}()
	res, err := fn(tx, rr)
	if err != nil {
		// 执行内容发生异常
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		// 提交发生异常
		return nil, errors.Wrapf(err, "committing transaction: %v", err)
	}
	return res, nil
}
