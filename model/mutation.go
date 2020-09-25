// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"app/model/example"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAuthUser = "AuthUser"
	TypeExample  = "Example"
)

// AuthUserMutation represents an operation that mutate the AuthUsers
// nodes in the graph.
type AuthUserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	password      *string
	nickname      *string
	email         *string
	remark        *string
	avatar        *string
	status        *uint8
	addstatus     *uint8
	create_time   *time.Time
	update_time   *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*AuthUser, error)
}

var _ ent.Mutation = (*AuthUserMutation)(nil)

// authuserOption allows to manage the mutation configuration using functional options.
type authuserOption func(*AuthUserMutation)

// newAuthUserMutation creates new mutation for $n.Name.
func newAuthUserMutation(c config, op Op, opts ...authuserOption) *AuthUserMutation {
	m := &AuthUserMutation{
		config:        c,
		op:            op,
		typ:           TypeAuthUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAuthUserID sets the id field of the mutation.
func withAuthUserID(id int) authuserOption {
	return func(m *AuthUserMutation) {
		var (
			err   error
			once  sync.Once
			value *AuthUser
		)
		m.oldValue = func(ctx context.Context) (*AuthUser, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().AuthUser.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAuthUser sets the old AuthUser of the mutation.
func withAuthUser(node *AuthUser) authuserOption {
	return func(m *AuthUserMutation) {
		m.oldValue = func(context.Context) (*AuthUser, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AuthUserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AuthUserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("model: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *AuthUserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUsername sets the username field.
func (m *AuthUserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the username value in the mutation.
func (m *AuthUserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old username value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUsername is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername reset all changes of the "username" field.
func (m *AuthUserMutation) ResetUsername() {
	m.username = nil
}

// SetPassword sets the password field.
func (m *AuthUserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the password value in the mutation.
func (m *AuthUserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old password value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPassword is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword reset all changes of the "password" field.
func (m *AuthUserMutation) ResetPassword() {
	m.password = nil
}

// SetNickname sets the nickname field.
func (m *AuthUserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the nickname value in the mutation.
func (m *AuthUserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old nickname value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNickname is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ResetNickname reset all changes of the "nickname" field.
func (m *AuthUserMutation) ResetNickname() {
	m.nickname = nil
}

// SetEmail sets the email field.
func (m *AuthUserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the email value in the mutation.
func (m *AuthUserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old email value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail reset all changes of the "email" field.
func (m *AuthUserMutation) ResetEmail() {
	m.email = nil
}

// SetRemark sets the remark field.
func (m *AuthUserMutation) SetRemark(s string) {
	m.remark = &s
}

// Remark returns the remark value in the mutation.
func (m *AuthUserMutation) Remark() (r string, exists bool) {
	v := m.remark
	if v == nil {
		return
	}
	return *v, true
}

// OldRemark returns the old remark value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldRemark(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRemark is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRemark requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRemark: %w", err)
	}
	return oldValue.Remark, nil
}

// ClearRemark clears the value of remark.
func (m *AuthUserMutation) ClearRemark() {
	m.remark = nil
	m.clearedFields[authuser.FieldRemark] = struct{}{}
}

// RemarkCleared returns if the field remark was cleared in this mutation.
func (m *AuthUserMutation) RemarkCleared() bool {
	_, ok := m.clearedFields[authuser.FieldRemark]
	return ok
}

// ResetRemark reset all changes of the "remark" field.
func (m *AuthUserMutation) ResetRemark() {
	m.remark = nil
	delete(m.clearedFields, authuser.FieldRemark)
}

// SetAvatar sets the avatar field.
func (m *AuthUserMutation) SetAvatar(s string) {
	m.avatar = &s
}

// Avatar returns the avatar value in the mutation.
func (m *AuthUserMutation) Avatar() (r string, exists bool) {
	v := m.avatar
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatar returns the old avatar value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldAvatar(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAvatar is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAvatar requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatar: %w", err)
	}
	return oldValue.Avatar, nil
}

// ClearAvatar clears the value of avatar.
func (m *AuthUserMutation) ClearAvatar() {
	m.avatar = nil
	m.clearedFields[authuser.FieldAvatar] = struct{}{}
}

// AvatarCleared returns if the field avatar was cleared in this mutation.
func (m *AuthUserMutation) AvatarCleared() bool {
	_, ok := m.clearedFields[authuser.FieldAvatar]
	return ok
}

// ResetAvatar reset all changes of the "avatar" field.
func (m *AuthUserMutation) ResetAvatar() {
	m.avatar = nil
	delete(m.clearedFields, authuser.FieldAvatar)
}

// SetStatus sets the status field.
func (m *AuthUserMutation) SetStatus(u uint8) {
	m.status = &u
	m.addstatus = nil
}

// Status returns the status value in the mutation.
func (m *AuthUserMutation) Status() (r uint8, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old status value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldStatus(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds u to status.
func (m *AuthUserMutation) AddStatus(u uint8) {
	if m.addstatus != nil {
		*m.addstatus += u
	} else {
		m.addstatus = &u
	}
}

// AddedStatus returns the value that was added to the status field in this mutation.
func (m *AuthUserMutation) AddedStatus() (r uint8, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus reset all changes of the "status" field.
func (m *AuthUserMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// SetCreateTime sets the create_time field.
func (m *AuthUserMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the create_time value in the mutation.
func (m *AuthUserMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old create_time value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime reset all changes of the "create_time" field.
func (m *AuthUserMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the update_time field.
func (m *AuthUserMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the update_time value in the mutation.
func (m *AuthUserMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old update_time value of the AuthUser.
// If the AuthUser object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AuthUserMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime reset all changes of the "update_time" field.
func (m *AuthUserMutation) ResetUpdateTime() {
	m.update_time = nil
}

// Op returns the operation name.
func (m *AuthUserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (AuthUser).
func (m *AuthUserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *AuthUserMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.username != nil {
		fields = append(fields, authuser.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, authuser.FieldPassword)
	}
	if m.nickname != nil {
		fields = append(fields, authuser.FieldNickname)
	}
	if m.email != nil {
		fields = append(fields, authuser.FieldEmail)
	}
	if m.remark != nil {
		fields = append(fields, authuser.FieldRemark)
	}
	if m.avatar != nil {
		fields = append(fields, authuser.FieldAvatar)
	}
	if m.status != nil {
		fields = append(fields, authuser.FieldStatus)
	}
	if m.create_time != nil {
		fields = append(fields, authuser.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, authuser.FieldUpdateTime)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *AuthUserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case authuser.FieldUsername:
		return m.Username()
	case authuser.FieldPassword:
		return m.Password()
	case authuser.FieldNickname:
		return m.Nickname()
	case authuser.FieldEmail:
		return m.Email()
	case authuser.FieldRemark:
		return m.Remark()
	case authuser.FieldAvatar:
		return m.Avatar()
	case authuser.FieldStatus:
		return m.Status()
	case authuser.FieldCreateTime:
		return m.CreateTime()
	case authuser.FieldUpdateTime:
		return m.UpdateTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *AuthUserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case authuser.FieldUsername:
		return m.OldUsername(ctx)
	case authuser.FieldPassword:
		return m.OldPassword(ctx)
	case authuser.FieldNickname:
		return m.OldNickname(ctx)
	case authuser.FieldEmail:
		return m.OldEmail(ctx)
	case authuser.FieldRemark:
		return m.OldRemark(ctx)
	case authuser.FieldAvatar:
		return m.OldAvatar(ctx)
	case authuser.FieldStatus:
		return m.OldStatus(ctx)
	case authuser.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case authuser.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	}
	return nil, fmt.Errorf("unknown AuthUser field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *AuthUserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case authuser.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case authuser.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case authuser.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case authuser.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case authuser.FieldRemark:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRemark(v)
		return nil
	case authuser.FieldAvatar:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatar(v)
		return nil
	case authuser.FieldStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case authuser.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case authuser.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	}
	return fmt.Errorf("unknown AuthUser field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *AuthUserMutation) AddedFields() []string {
	var fields []string
	if m.addstatus != nil {
		fields = append(fields, authuser.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *AuthUserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case authuser.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *AuthUserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case authuser.FieldStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown AuthUser numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *AuthUserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(authuser.FieldRemark) {
		fields = append(fields, authuser.FieldRemark)
	}
	if m.FieldCleared(authuser.FieldAvatar) {
		fields = append(fields, authuser.FieldAvatar)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *AuthUserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *AuthUserMutation) ClearField(name string) error {
	switch name {
	case authuser.FieldRemark:
		m.ClearRemark()
		return nil
	case authuser.FieldAvatar:
		m.ClearAvatar()
		return nil
	}
	return fmt.Errorf("unknown AuthUser nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *AuthUserMutation) ResetField(name string) error {
	switch name {
	case authuser.FieldUsername:
		m.ResetUsername()
		return nil
	case authuser.FieldPassword:
		m.ResetPassword()
		return nil
	case authuser.FieldNickname:
		m.ResetNickname()
		return nil
	case authuser.FieldEmail:
		m.ResetEmail()
		return nil
	case authuser.FieldRemark:
		m.ResetRemark()
		return nil
	case authuser.FieldAvatar:
		m.ResetAvatar()
		return nil
	case authuser.FieldStatus:
		m.ResetStatus()
		return nil
	case authuser.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case authuser.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	}
	return fmt.Errorf("unknown AuthUser field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *AuthUserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *AuthUserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *AuthUserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *AuthUserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *AuthUserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *AuthUserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *AuthUserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown AuthUser unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *AuthUserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown AuthUser edge %s", name)
}

// ExampleMutation represents an operation that mutate the Examples
// nodes in the graph.
type ExampleMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Example, error)
}

var _ ent.Mutation = (*ExampleMutation)(nil)

// exampleOption allows to manage the mutation configuration using functional options.
type exampleOption func(*ExampleMutation)

// newExampleMutation creates new mutation for $n.Name.
func newExampleMutation(c config, op Op, opts ...exampleOption) *ExampleMutation {
	m := &ExampleMutation{
		config:        c,
		op:            op,
		typ:           TypeExample,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withExampleID sets the id field of the mutation.
func withExampleID(id int) exampleOption {
	return func(m *ExampleMutation) {
		var (
			err   error
			once  sync.Once
			value *Example
		)
		m.oldValue = func(ctx context.Context) (*Example, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Example.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withExample sets the old Example of the mutation.
func withExample(node *Example) exampleOption {
	return func(m *ExampleMutation) {
		m.oldValue = func(context.Context) (*Example, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ExampleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ExampleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("model: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *ExampleMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *ExampleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *ExampleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Example.
// If the Example object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *ExampleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *ExampleMutation) ResetName() {
	m.name = nil
}

// Op returns the operation name.
func (m *ExampleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Example).
func (m *ExampleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *ExampleMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, example.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *ExampleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case example.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *ExampleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case example.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Example field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ExampleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case example.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Example field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *ExampleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *ExampleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ExampleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Example numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *ExampleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *ExampleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *ExampleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Example nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *ExampleMutation) ResetField(name string) error {
	switch name {
	case example.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Example field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *ExampleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *ExampleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *ExampleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *ExampleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *ExampleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *ExampleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *ExampleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Example unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *ExampleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Example edge %s", name)
}
