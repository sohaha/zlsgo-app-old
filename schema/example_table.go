package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"

	"app/global"
)

type Example struct {
	ent.Schema
}

func (Example) Config() ent.Config {
	global.Read(false)
	return ent.Config{
		// 表前缀 + 表名
		Table: global.DatabaseConf().Prefix + "example",
	}
}

func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

func (Example) Edges() []ent.Edge {
	return nil
}
