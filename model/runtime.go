// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"app/schema"
	"time"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	authuserFields := schema.AuthUser{}.Fields()
	_ = authuserFields
	// authuserDescPassword is the schema descriptor for password field.
	authuserDescPassword := authuserFields[1].Descriptor()
	// authuser.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	authuser.PasswordValidator = authuserDescPassword.Validators[0].(func(string) error)
	// authuserDescNickname is the schema descriptor for nickname field.
	authuserDescNickname := authuserFields[2].Descriptor()
	// authuser.DefaultNickname holds the default value on creation for the nickname field.
	authuser.DefaultNickname = authuserDescNickname.Default.(string)
	// authuserDescEmail is the schema descriptor for email field.
	authuserDescEmail := authuserFields[3].Descriptor()
	// authuser.DefaultEmail holds the default value on creation for the email field.
	authuser.DefaultEmail = authuserDescEmail.Default.(string)
	// authuserDescRemark is the schema descriptor for remark field.
	authuserDescRemark := authuserFields[4].Descriptor()
	// authuser.DefaultRemark holds the default value on creation for the remark field.
	authuser.DefaultRemark = authuserDescRemark.Default.(string)
	// authuserDescAvatar is the schema descriptor for avatar field.
	authuserDescAvatar := authuserFields[5].Descriptor()
	// authuser.DefaultAvatar holds the default value on creation for the avatar field.
	authuser.DefaultAvatar = authuserDescAvatar.Default.(string)
	// authuserDescStatus is the schema descriptor for status field.
	authuserDescStatus := authuserFields[6].Descriptor()
	// authuser.DefaultStatus holds the default value on creation for the status field.
	authuser.DefaultStatus = authuserDescStatus.Default.(uint8)
	// authuserDescCreateTime is the schema descriptor for create_time field.
	authuserDescCreateTime := authuserFields[7].Descriptor()
	// authuser.DefaultCreateTime holds the default value on creation for the create_time field.
	authuser.DefaultCreateTime = authuserDescCreateTime.Default.(func() time.Time)
	// authuserDescUpdateTime is the schema descriptor for update_time field.
	authuserDescUpdateTime := authuserFields[8].Descriptor()
	// authuser.DefaultUpdateTime holds the default value on creation for the update_time field.
	authuser.DefaultUpdateTime = authuserDescUpdateTime.Default.(func() time.Time)
}
