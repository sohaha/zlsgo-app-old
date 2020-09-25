// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// AuthUser is the model entity for the AuthUser schema.
type AuthUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Status holds the value of the "status" field.
	Status uint8 `json:"status,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthUser) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // username
		&sql.NullString{}, // password
		&sql.NullString{}, // nickname
		&sql.NullString{}, // email
		&sql.NullString{}, // remark
		&sql.NullString{}, // avatar
		&sql.NullInt64{},  // status
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthUser fields.
func (au *AuthUser) assignValues(values ...interface{}) error {
	if m, n := len(values), len(authuser.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	au.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[0])
	} else if value.Valid {
		au.Username = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[1])
	} else if value.Valid {
		au.Password = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nickname", values[2])
	} else if value.Valid {
		au.Nickname = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[3])
	} else if value.Valid {
		au.Email = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field remark", values[4])
	} else if value.Valid {
		au.Remark = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field avatar", values[5])
	} else if value.Valid {
		au.Avatar = value.String
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[6])
	} else if value.Valid {
		au.Status = uint8(value.Int64)
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[7])
	} else if value.Valid {
		au.CreateTime = value.Time
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[8])
	} else if value.Valid {
		au.UpdateTime = value.Time
	}
	return nil
}

// Update returns a builder for updating this AuthUser.
// Note that, you need to call AuthUser.Unwrap() before calling this method, if this AuthUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *AuthUser) Update() *AuthUserUpdateOne {
	return (&AuthUserClient{config: au.config}).UpdateOne(au)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (au *AuthUser) Unwrap() *AuthUser {
	tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("model: AuthUser is not a transactional entity")
	}
	au.config.driver = tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *AuthUser) String() string {
	var builder strings.Builder
	builder.WriteString("AuthUser(")
	builder.WriteString(fmt.Sprintf("id=%v", au.ID))
	builder.WriteString(", username=")
	builder.WriteString(au.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", nickname=")
	builder.WriteString(au.Nickname)
	builder.WriteString(", email=")
	builder.WriteString(au.Email)
	builder.WriteString(", remark=")
	builder.WriteString(au.Remark)
	builder.WriteString(", avatar=")
	builder.WriteString(au.Avatar)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", au.Status))
	builder.WriteString(", create_time=")
	builder.WriteString(au.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(au.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AuthUsers is a parsable slice of AuthUser.
type AuthUsers []*AuthUser

func (au AuthUsers) config(cfg config) {
	for _i := range au {
		au[_i].config = cfg
	}
}
