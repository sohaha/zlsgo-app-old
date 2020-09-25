// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ZExampleColumns holds the columns for the "z_example" table.
	ZExampleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// ZExampleTable holds the schema information for the "z_example" table.
	ZExampleTable = &schema.Table{
		Name:        "z_example",
		Columns:     ZExampleColumns,
		PrimaryKey:  []*schema.Column{ZExampleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ZExampleTable,
	}
)

func init() {
}
