// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/example"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// Example is the model entity for the Example schema.
type Example struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Example) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Example fields.
func (e *Example) assignValues(values ...interface{}) error {
	if m, n := len(values), len(example.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		e.Name = value.String
	}
	return nil
}

// Update returns a builder for updating this Example.
// Note that, you need to call Example.Unwrap() before calling this method, if this Example
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Example) Update() *ExampleUpdateOne {
	return (&ExampleClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Example) Unwrap() *Example {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("model: Example is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Example) String() string {
	var builder strings.Builder
	builder.WriteString("Example(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", name=")
	builder.WriteString(e.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Examples is a parsable slice of Example.
type Examples []*Example

func (e Examples) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
