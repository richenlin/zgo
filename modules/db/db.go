package db

import (
	"context"

	"github.com/pkg/errors"
)

// DB db
type DB interface {
	Tx(context.Context) (Tx, error)
}

// Tx tx
type Tx interface {
	Commit() error
	Rollback() error
}

// ResultRef ref
type ResultRef struct {
	Data interface{}
}

// WithTx 执行带有事务的方法, 在一个事务中完成所有的内容
func WithTx(ctx context.Context, db DB, fn func(tx Tx) error) error {
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
