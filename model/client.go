// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"context"
	"fmt"
	"log"

	"app/model/migrate"

	"app/model/example"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Example is the client for interacting with the Example builders.
	Example *ExampleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Example = NewExampleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("model: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("model: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Example: NewExampleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:  cfg,
		Example: NewExampleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Example.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Example.Use(hooks...)
}

// ExampleClient is a client for the Example schema.
type ExampleClient struct {
	config
}

// NewExampleClient returns a client for the Example from the given config.
func NewExampleClient(c config) *ExampleClient {
	return &ExampleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `example.Hooks(f(g(h())))`.
func (c *ExampleClient) Use(hooks ...Hook) {
	c.hooks.Example = append(c.hooks.Example, hooks...)
}

// Create returns a create builder for Example.
func (c *ExampleClient) Create() *ExampleCreate {
	mutation := newExampleMutation(c.config, OpCreate)
	return &ExampleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Example entities.
func (c *ExampleClient) CreateBulk(builders ...*ExampleCreate) *ExampleCreateBulk {
	return &ExampleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Example.
func (c *ExampleClient) Update() *ExampleUpdate {
	mutation := newExampleMutation(c.config, OpUpdate)
	return &ExampleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExampleClient) UpdateOne(e *Example) *ExampleUpdateOne {
	mutation := newExampleMutation(c.config, OpUpdateOne, withExample(e))
	return &ExampleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExampleClient) UpdateOneID(id int) *ExampleUpdateOne {
	mutation := newExampleMutation(c.config, OpUpdateOne, withExampleID(id))
	return &ExampleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Example.
func (c *ExampleClient) Delete() *ExampleDelete {
	mutation := newExampleMutation(c.config, OpDelete)
	return &ExampleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ExampleClient) DeleteOne(e *Example) *ExampleDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ExampleClient) DeleteOneID(id int) *ExampleDeleteOne {
	builder := c.Delete().Where(example.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExampleDeleteOne{builder}
}

// Query returns a query builder for Example.
func (c *ExampleClient) Query() *ExampleQuery {
	return &ExampleQuery{config: c.config}
}

// Get returns a Example entity by its id.
func (c *ExampleClient) Get(ctx context.Context, id int) (*Example, error) {
	return c.Query().Where(example.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExampleClient) GetX(ctx context.Context, id int) *Example {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ExampleClient) Hooks() []Hook {
	return c.hooks.Example
}
