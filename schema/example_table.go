package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

type Example struct {
	ent.Schema
}

func (Example) Config() ent.Config {
	return ent.Config{
		Table: Prefix + "example",
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
