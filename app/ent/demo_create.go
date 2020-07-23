// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/suisrc/zgo/app/ent/demo"
)

// DemoCreate is the builder for creating a Demo entity.
type DemoCreate struct {
	config
	mutation *DemoMutation
	hooks    []Hook
}

// Mutation returns the DemoMutation object of the builder.
func (dc *DemoCreate) Mutation() *DemoMutation {
	return dc.mutation
}

// Save creates the Demo in the database.
func (dc *DemoCreate) Save(ctx context.Context) (*Demo, error) {
	var (
		err  error
		node *Demo
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DemoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DemoCreate) SaveX(ctx context.Context) *Demo {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DemoCreate) sqlSave(ctx context.Context) (*Demo, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DemoCreate) createSpec() (*Demo, *sqlgraph.CreateSpec) {
	var (
		d     = &Demo{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: demo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: demo.FieldID,
			},
		}
	)
	return d, _spec
}
