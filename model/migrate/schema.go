// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ZAuthUserColumns holds the columns for the "z_auth_user" table.
	ZAuthUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Size: 200},
		{Name: "nickname", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Default: ""},
		{Name: "remark", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "status", Type: field.TypeUint8},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// ZAuthUserTable holds the schema information for the "z_auth_user" table.
	ZAuthUserTable = &schema.Table{
		Name:        "z_auth_user",
		Columns:     ZAuthUserColumns,
		PrimaryKey:  []*schema.Column{ZAuthUserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
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
		ZAuthUserTable,
		ZExampleTable,
	}
)

func init() {
}
