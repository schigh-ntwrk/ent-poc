// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/predicate"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/veterinarian"
)

// VeterinarianDelete is the builder for deleting a Veterinarian entity.
type VeterinarianDelete struct {
	config
	hooks    []Hook
	mutation *VeterinarianMutation
}

// Where adds a new predicate to the delete builder.
func (vd *VeterinarianDelete) Where(ps ...predicate.Veterinarian) *VeterinarianDelete {
	vd.mutation.predicates = append(vd.mutation.predicates, ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VeterinarianDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VeterinarianMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VeterinarianDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VeterinarianDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: veterinarian.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: veterinarian.FieldID,
			},
		},
	}
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// VeterinarianDeleteOne is the builder for deleting a single Veterinarian entity.
type VeterinarianDeleteOne struct {
	vd *VeterinarianDelete
}

// Exec executes the deletion query.
func (vdo *VeterinarianDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{veterinarian.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VeterinarianDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
