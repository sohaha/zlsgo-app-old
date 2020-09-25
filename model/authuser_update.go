// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"app/model/predicate"
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AuthUserUpdate is the builder for updating AuthUser entities.
type AuthUserUpdate struct {
	config
	hooks      []Hook
	mutation   *AuthUserMutation
	predicates []predicate.AuthUser
}

// Where adds a new predicate for the builder.
func (auu *AuthUserUpdate) Where(ps ...predicate.AuthUser) *AuthUserUpdate {
	auu.predicates = append(auu.predicates, ps...)
	return auu
}

// SetUsername sets the username field.
func (auu *AuthUserUpdate) SetUsername(s string) *AuthUserUpdate {
	auu.mutation.SetUsername(s)
	return auu
}

// SetPassword sets the password field.
func (auu *AuthUserUpdate) SetPassword(s string) *AuthUserUpdate {
	auu.mutation.SetPassword(s)
	return auu
}

// SetNickname sets the nickname field.
func (auu *AuthUserUpdate) SetNickname(s string) *AuthUserUpdate {
	auu.mutation.SetNickname(s)
	return auu
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (auu *AuthUserUpdate) SetNillableNickname(s *string) *AuthUserUpdate {
	if s != nil {
		auu.SetNickname(*s)
	}
	return auu
}

// SetEmail sets the email field.
func (auu *AuthUserUpdate) SetEmail(s string) *AuthUserUpdate {
	auu.mutation.SetEmail(s)
	return auu
}

// SetNillableEmail sets the email field if the given value is not nil.
func (auu *AuthUserUpdate) SetNillableEmail(s *string) *AuthUserUpdate {
	if s != nil {
		auu.SetEmail(*s)
	}
	return auu
}

// SetRemark sets the remark field.
func (auu *AuthUserUpdate) SetRemark(s string) *AuthUserUpdate {
	auu.mutation.SetRemark(s)
	return auu
}

// SetNillableRemark sets the remark field if the given value is not nil.
func (auu *AuthUserUpdate) SetNillableRemark(s *string) *AuthUserUpdate {
	if s != nil {
		auu.SetRemark(*s)
	}
	return auu
}

// ClearRemark clears the value of remark.
func (auu *AuthUserUpdate) ClearRemark() *AuthUserUpdate {
	auu.mutation.ClearRemark()
	return auu
}

// SetAvatar sets the avatar field.
func (auu *AuthUserUpdate) SetAvatar(s string) *AuthUserUpdate {
	auu.mutation.SetAvatar(s)
	return auu
}

// SetNillableAvatar sets the avatar field if the given value is not nil.
func (auu *AuthUserUpdate) SetNillableAvatar(s *string) *AuthUserUpdate {
	if s != nil {
		auu.SetAvatar(*s)
	}
	return auu
}

// ClearAvatar clears the value of avatar.
func (auu *AuthUserUpdate) ClearAvatar() *AuthUserUpdate {
	auu.mutation.ClearAvatar()
	return auu
}

// SetStatus sets the status field.
func (auu *AuthUserUpdate) SetStatus(u uint8) *AuthUserUpdate {
	auu.mutation.ResetStatus()
	auu.mutation.SetStatus(u)
	return auu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (auu *AuthUserUpdate) SetNillableStatus(u *uint8) *AuthUserUpdate {
	if u != nil {
		auu.SetStatus(*u)
	}
	return auu
}

// AddStatus adds u to status.
func (auu *AuthUserUpdate) AddStatus(u uint8) *AuthUserUpdate {
	auu.mutation.AddStatus(u)
	return auu
}

// SetUpdateTime sets the update_time field.
func (auu *AuthUserUpdate) SetUpdateTime(t time.Time) *AuthUserUpdate {
	auu.mutation.SetUpdateTime(t)
	return auu
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (auu *AuthUserUpdate) SetNillableUpdateTime(t *time.Time) *AuthUserUpdate {
	if t != nil {
		auu.SetUpdateTime(*t)
	}
	return auu
}

// Mutation returns the AuthUserMutation object of the builder.
func (auu *AuthUserUpdate) Mutation() *AuthUserMutation {
	return auu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (auu *AuthUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(auu.hooks) == 0 {
		if err = auu.check(); err != nil {
			return 0, err
		}
		affected, err = auu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auu.check(); err != nil {
				return 0, err
			}
			auu.mutation = mutation
			affected, err = auu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(auu.hooks) - 1; i >= 0; i-- {
			mut = auu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (auu *AuthUserUpdate) SaveX(ctx context.Context) int {
	affected, err := auu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auu *AuthUserUpdate) Exec(ctx context.Context) error {
	_, err := auu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auu *AuthUserUpdate) ExecX(ctx context.Context) {
	if err := auu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auu *AuthUserUpdate) check() error {
	if v, ok := auu.mutation.Password(); ok {
		if err := authuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("model: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (auu *AuthUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authuser.Table,
			Columns: authuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authuser.FieldID,
			},
		},
	}
	if ps := auu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldUsername,
		})
	}
	if value, ok := auu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldPassword,
		})
	}
	if value, ok := auu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldNickname,
		})
	}
	if value, ok := auu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldEmail,
		})
	}
	if value, ok := auu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldRemark,
		})
	}
	if auu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: authuser.FieldRemark,
		})
	}
	if value, ok := auu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldAvatar,
		})
	}
	if auu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: authuser.FieldAvatar,
		})
	}
	if value, ok := auu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: authuser.FieldStatus,
		})
	}
	if value, ok := auu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: authuser.FieldStatus,
		})
	}
	if value, ok := auu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authuser.FieldUpdateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AuthUserUpdateOne is the builder for updating a single AuthUser entity.
type AuthUserUpdateOne struct {
	config
	hooks    []Hook
	mutation *AuthUserMutation
}

// SetUsername sets the username field.
func (auuo *AuthUserUpdateOne) SetUsername(s string) *AuthUserUpdateOne {
	auuo.mutation.SetUsername(s)
	return auuo
}

// SetPassword sets the password field.
func (auuo *AuthUserUpdateOne) SetPassword(s string) *AuthUserUpdateOne {
	auuo.mutation.SetPassword(s)
	return auuo
}

// SetNickname sets the nickname field.
func (auuo *AuthUserUpdateOne) SetNickname(s string) *AuthUserUpdateOne {
	auuo.mutation.SetNickname(s)
	return auuo
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (auuo *AuthUserUpdateOne) SetNillableNickname(s *string) *AuthUserUpdateOne {
	if s != nil {
		auuo.SetNickname(*s)
	}
	return auuo
}

// SetEmail sets the email field.
func (auuo *AuthUserUpdateOne) SetEmail(s string) *AuthUserUpdateOne {
	auuo.mutation.SetEmail(s)
	return auuo
}

// SetNillableEmail sets the email field if the given value is not nil.
func (auuo *AuthUserUpdateOne) SetNillableEmail(s *string) *AuthUserUpdateOne {
	if s != nil {
		auuo.SetEmail(*s)
	}
	return auuo
}

// SetRemark sets the remark field.
func (auuo *AuthUserUpdateOne) SetRemark(s string) *AuthUserUpdateOne {
	auuo.mutation.SetRemark(s)
	return auuo
}

// SetNillableRemark sets the remark field if the given value is not nil.
func (auuo *AuthUserUpdateOne) SetNillableRemark(s *string) *AuthUserUpdateOne {
	if s != nil {
		auuo.SetRemark(*s)
	}
	return auuo
}

// ClearRemark clears the value of remark.
func (auuo *AuthUserUpdateOne) ClearRemark() *AuthUserUpdateOne {
	auuo.mutation.ClearRemark()
	return auuo
}

// SetAvatar sets the avatar field.
func (auuo *AuthUserUpdateOne) SetAvatar(s string) *AuthUserUpdateOne {
	auuo.mutation.SetAvatar(s)
	return auuo
}

// SetNillableAvatar sets the avatar field if the given value is not nil.
func (auuo *AuthUserUpdateOne) SetNillableAvatar(s *string) *AuthUserUpdateOne {
	if s != nil {
		auuo.SetAvatar(*s)
	}
	return auuo
}

// ClearAvatar clears the value of avatar.
func (auuo *AuthUserUpdateOne) ClearAvatar() *AuthUserUpdateOne {
	auuo.mutation.ClearAvatar()
	return auuo
}

// SetStatus sets the status field.
func (auuo *AuthUserUpdateOne) SetStatus(u uint8) *AuthUserUpdateOne {
	auuo.mutation.ResetStatus()
	auuo.mutation.SetStatus(u)
	return auuo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (auuo *AuthUserUpdateOne) SetNillableStatus(u *uint8) *AuthUserUpdateOne {
	if u != nil {
		auuo.SetStatus(*u)
	}
	return auuo
}

// AddStatus adds u to status.
func (auuo *AuthUserUpdateOne) AddStatus(u uint8) *AuthUserUpdateOne {
	auuo.mutation.AddStatus(u)
	return auuo
}

// SetUpdateTime sets the update_time field.
func (auuo *AuthUserUpdateOne) SetUpdateTime(t time.Time) *AuthUserUpdateOne {
	auuo.mutation.SetUpdateTime(t)
	return auuo
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (auuo *AuthUserUpdateOne) SetNillableUpdateTime(t *time.Time) *AuthUserUpdateOne {
	if t != nil {
		auuo.SetUpdateTime(*t)
	}
	return auuo
}

// Mutation returns the AuthUserMutation object of the builder.
func (auuo *AuthUserUpdateOne) Mutation() *AuthUserMutation {
	return auuo.mutation
}

// Save executes the query and returns the updated entity.
func (auuo *AuthUserUpdateOne) Save(ctx context.Context) (*AuthUser, error) {
	var (
		err  error
		node *AuthUser
	)
	if len(auuo.hooks) == 0 {
		if err = auuo.check(); err != nil {
			return nil, err
		}
		node, err = auuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auuo.check(); err != nil {
				return nil, err
			}
			auuo.mutation = mutation
			node, err = auuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auuo.hooks) - 1; i >= 0; i-- {
			mut = auuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auuo *AuthUserUpdateOne) SaveX(ctx context.Context) *AuthUser {
	node, err := auuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auuo *AuthUserUpdateOne) Exec(ctx context.Context) error {
	_, err := auuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auuo *AuthUserUpdateOne) ExecX(ctx context.Context) {
	if err := auuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auuo *AuthUserUpdateOne) check() error {
	if v, ok := auuo.mutation.Password(); ok {
		if err := authuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("model: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (auuo *AuthUserUpdateOne) sqlSave(ctx context.Context) (_node *AuthUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authuser.Table,
			Columns: authuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authuser.FieldID,
			},
		},
	}
	id, ok := auuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AuthUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldUsername,
		})
	}
	if value, ok := auuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldPassword,
		})
	}
	if value, ok := auuo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldNickname,
		})
	}
	if value, ok := auuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldEmail,
		})
	}
	if value, ok := auuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldRemark,
		})
	}
	if auuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: authuser.FieldRemark,
		})
	}
	if value, ok := auuo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldAvatar,
		})
	}
	if auuo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: authuser.FieldAvatar,
		})
	}
	if value, ok := auuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: authuser.FieldStatus,
		})
	}
	if value, ok := auuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: authuser.FieldStatus,
		})
	}
	if value, ok := auuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authuser.FieldUpdateTime,
		})
	}
	_node = &AuthUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
