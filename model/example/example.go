// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package example

const (
	// Label holds the string label denoting the example type in the database.
	Label = "example"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// Table holds the table name of the example in the database.
	Table = "z_example"
)

// Columns holds all SQL columns for example fields.
var Columns = []string{
	FieldID,
	FieldName,
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
