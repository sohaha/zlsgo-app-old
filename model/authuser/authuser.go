// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package authuser

import (
	"time"
)

const (
	// Label holds the string label denoting the authuser type in the database.
	Label = "auth_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"

	// Table holds the table name of the authuser in the database.
	Table = "z_auth_user"
)

// Columns holds all SQL columns for authuser fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldNickname,
	FieldEmail,
	FieldRemark,
	FieldAvatar,
	FieldStatus,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultNickname holds the default value on creation for the nickname field.
	DefaultNickname string
	// DefaultEmail holds the default value on creation for the email field.
	DefaultEmail string
	// DefaultRemark holds the default value on creation for the remark field.
	DefaultRemark string
	// DefaultAvatar holds the default value on creation for the avatar field.
	DefaultAvatar string
	// DefaultStatus holds the default value on creation for the status field.
	DefaultStatus uint8
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
)
