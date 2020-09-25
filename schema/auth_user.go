package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

type AuthUser struct {
	ent.Schema
}

func (AuthUser) Config() ent.Config {
	return ent.Config{
		Table: Prefix + "auth_user",
	}
}

func (AuthUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Comment("用户名"),
		field.String("password").MaxLen(200).Sensitive().Comment("用户密码"),
		field.String("nickname").Default(""),
		field.String("email").Default(""),
		field.String("remark").Optional().Default(""),
		field.String("avatar").Optional().Default(""),
		field.Uint8("status").Default(0).Comment("状态:-1软删除,0待激活,1正常,2禁止"),
		field.Time("create_time").Default(time.Now).Immutable().Comment("创建时间"),
		field.Time("update_time").Default(time.Now).Comment("更新时间"),
	}
}

func (AuthUser) Edges() []ent.Edge {
	return nil
}
