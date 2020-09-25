// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"app/model/predicate"
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AuthUserDelete is the builder for deleting a AuthUser entity.
type AuthUserDelete struct {
	config
	hooks      []Hook
	mutation   *AuthUserMutation
	predicates []predicate.AuthUser
}

// Where adds a new predicate to the delete builder.
func (aud *AuthUserDelete) Where(ps ...predicate.AuthUser) *AuthUserDelete {
	aud.predicates = append(aud.predicates, ps...)
	return aud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aud *AuthUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aud.hooks) == 0 {
		affected, err = aud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aud.mutation = mutation
			affected, err = aud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aud.hooks) - 1; i >= 0; i-- {
			mut = aud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aud *AuthUserDelete) ExecX(ctx context.Context) int {
	n, err := aud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aud *AuthUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authuser.FieldID,
			},
		},
	}
	if ps := aud.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aud.driver, _spec)
}

// AuthUserDeleteOne is the builder for deleting a single AuthUser entity.
type AuthUserDeleteOne struct {
	aud *AuthUserDelete
}

// Exec executes the deletion query.
func (audo *AuthUserDeleteOne) Exec(ctx context.Context) error {
	n, err := audo.aud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (audo *AuthUserDeleteOne) ExecX(ctx context.Context) {
	audo.aud.ExecX(ctx)
}
