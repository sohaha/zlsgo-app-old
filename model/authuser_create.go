// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AuthUserCreate is the builder for creating a AuthUser entity.
type AuthUserCreate struct {
	config
	mutation *AuthUserMutation
	hooks    []Hook
}

// SetUsername sets the username field.
func (auc *AuthUserCreate) SetUsername(s string) *AuthUserCreate {
	auc.mutation.SetUsername(s)
	return auc
}

// SetPassword sets the password field.
func (auc *AuthUserCreate) SetPassword(s string) *AuthUserCreate {
	auc.mutation.SetPassword(s)
	return auc
}

// SetNickname sets the nickname field.
func (auc *AuthUserCreate) SetNickname(s string) *AuthUserCreate {
	auc.mutation.SetNickname(s)
	return auc
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableNickname(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetNickname(*s)
	}
	return auc
}

// SetEmail sets the email field.
func (auc *AuthUserCreate) SetEmail(s string) *AuthUserCreate {
	auc.mutation.SetEmail(s)
	return auc
}

// SetNillableEmail sets the email field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableEmail(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetEmail(*s)
	}
	return auc
}

// SetRemark sets the remark field.
func (auc *AuthUserCreate) SetRemark(s string) *AuthUserCreate {
	auc.mutation.SetRemark(s)
	return auc
}

// SetNillableRemark sets the remark field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableRemark(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetRemark(*s)
	}
	return auc
}

// SetAvatar sets the avatar field.
func (auc *AuthUserCreate) SetAvatar(s string) *AuthUserCreate {
	auc.mutation.SetAvatar(s)
	return auc
}

// SetNillableAvatar sets the avatar field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableAvatar(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetAvatar(*s)
	}
	return auc
}

// SetStatus sets the status field.
func (auc *AuthUserCreate) SetStatus(u uint8) *AuthUserCreate {
	auc.mutation.SetStatus(u)
	return auc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableStatus(u *uint8) *AuthUserCreate {
	if u != nil {
		auc.SetStatus(*u)
	}
	return auc
}

// SetCreateTime sets the create_time field.
func (auc *AuthUserCreate) SetCreateTime(t time.Time) *AuthUserCreate {
	auc.mutation.SetCreateTime(t)
	return auc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableCreateTime(t *time.Time) *AuthUserCreate {
	if t != nil {
		auc.SetCreateTime(*t)
	}
	return auc
}

// SetUpdateTime sets the update_time field.
func (auc *AuthUserCreate) SetUpdateTime(t time.Time) *AuthUserCreate {
	auc.mutation.SetUpdateTime(t)
	return auc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableUpdateTime(t *time.Time) *AuthUserCreate {
	if t != nil {
		auc.SetUpdateTime(*t)
	}
	return auc
}

// Mutation returns the AuthUserMutation object of the builder.
func (auc *AuthUserCreate) Mutation() *AuthUserMutation {
	return auc.mutation
}

// Save creates the AuthUser in the database.
func (auc *AuthUserCreate) Save(ctx context.Context) (*AuthUser, error) {
	var (
		err  error
		node *AuthUser
	)
	auc.defaults()
	if len(auc.hooks) == 0 {
		if err = auc.check(); err != nil {
			return nil, err
		}
		node, err = auc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auc.check(); err != nil {
				return nil, err
			}
			auc.mutation = mutation
			node, err = auc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auc.hooks) - 1; i >= 0; i-- {
			mut = auc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AuthUserCreate) SaveX(ctx context.Context) *AuthUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (auc *AuthUserCreate) defaults() {
	if _, ok := auc.mutation.Nickname(); !ok {
		v := authuser.DefaultNickname
		auc.mutation.SetNickname(v)
	}
	if _, ok := auc.mutation.Email(); !ok {
		v := authuser.DefaultEmail
		auc.mutation.SetEmail(v)
	}
	if _, ok := auc.mutation.Remark(); !ok {
		v := authuser.DefaultRemark
		auc.mutation.SetRemark(v)
	}
	if _, ok := auc.mutation.Avatar(); !ok {
		v := authuser.DefaultAvatar
		auc.mutation.SetAvatar(v)
	}
	if _, ok := auc.mutation.Status(); !ok {
		v := authuser.DefaultStatus
		auc.mutation.SetStatus(v)
	}
	if _, ok := auc.mutation.CreateTime(); !ok {
		v := authuser.DefaultCreateTime()
		auc.mutation.SetCreateTime(v)
	}
	if _, ok := auc.mutation.UpdateTime(); !ok {
		v := authuser.DefaultUpdateTime()
		auc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *AuthUserCreate) check() error {
	if _, ok := auc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("model: missing required field \"username\"")}
	}
	if _, ok := auc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("model: missing required field \"password\"")}
	}
	if v, ok := auc.mutation.Password(); ok {
		if err := authuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("model: validator failed for field \"password\": %w", err)}
		}
	}
	if _, ok := auc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New("model: missing required field \"nickname\"")}
	}
	if _, ok := auc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("model: missing required field \"email\"")}
	}
	if _, ok := auc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("model: missing required field \"status\"")}
	}
	if _, ok := auc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("model: missing required field \"create_time\"")}
	}
	if _, ok := auc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("model: missing required field \"update_time\"")}
	}
	return nil
}

func (auc *AuthUserCreate) sqlSave(ctx context.Context) (*AuthUser, error) {
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (auc *AuthUserCreate) createSpec() (*AuthUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthUser{config: auc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: authuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authuser.FieldID,
			},
		}
	)
	if value, ok := auc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := auc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := auc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := auc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := auc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := auc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := auc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: authuser.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := auc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authuser.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := auc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authuser.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	return _node, _spec
}

// AuthUserCreateBulk is the builder for creating a bulk of AuthUser entities.
type AuthUserCreateBulk struct {
	config
	builders []*AuthUserCreate
}

// Save creates the AuthUser entities in the database.
func (aucb *AuthUserCreateBulk) Save(ctx context.Context) ([]*AuthUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AuthUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (aucb *AuthUserCreateBulk) SaveX(ctx context.Context) []*AuthUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
