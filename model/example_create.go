// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/example"
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ExampleCreate is the builder for creating a Example entity.
type ExampleCreate struct {
	config
	mutation *ExampleMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ec *ExampleCreate) SetName(s string) *ExampleCreate {
	ec.mutation.SetName(s)
	return ec
}

// Mutation returns the ExampleMutation object of the builder.
func (ec *ExampleCreate) Mutation() *ExampleMutation {
	return ec.mutation
}

// Save creates the Example in the database.
func (ec *ExampleCreate) Save(ctx context.Context) (*Example, error) {
	var (
		err  error
		node *Example
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExampleCreate) SaveX(ctx context.Context) *Example {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExampleCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("model: missing required field \"name\"")}
	}
	return nil
}

func (ec *ExampleCreate) sqlSave(ctx context.Context) (*Example, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *ExampleCreate) createSpec() (*Example, *sqlgraph.CreateSpec) {
	var (
		_node = &Example{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: example.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: example.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: example.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// ExampleCreateBulk is the builder for creating a bulk of Example entities.
type ExampleCreateBulk struct {
	config
	builders []*ExampleCreate
}

// Save creates the Example entities in the database.
func (ecb *ExampleCreateBulk) Save(ctx context.Context) ([]*Example, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Example, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExampleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ecb *ExampleCreateBulk) SaveX(ctx context.Context) []*Example {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
