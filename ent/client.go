// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"rahnit-rmm/ent/migrate"

	"rahnit-rmm/ent/device"
	"rahnit-rmm/ent/revocation"
	"rahnit-rmm/ent/tunnelconfig"
	"rahnit-rmm/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Device is the client for interacting with the Device builders.
	Device *DeviceClient
	// Revocation is the client for interacting with the Revocation builders.
	Revocation *RevocationClient
	// TunnelConfig is the client for interacting with the TunnelConfig builders.
	TunnelConfig *TunnelConfigClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Device = NewDeviceClient(c.config)
	c.Revocation = NewRevocationClient(c.config)
	c.TunnelConfig = NewTunnelConfigClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Device:       NewDeviceClient(cfg),
		Revocation:   NewRevocationClient(cfg),
		TunnelConfig: NewTunnelConfigClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Device:       NewDeviceClient(cfg),
		Revocation:   NewRevocationClient(cfg),
		TunnelConfig: NewTunnelConfigClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Device.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Device.Use(hooks...)
	c.Revocation.Use(hooks...)
	c.TunnelConfig.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Device.Intercept(interceptors...)
	c.Revocation.Intercept(interceptors...)
	c.TunnelConfig.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DeviceMutation:
		return c.Device.mutate(ctx, m)
	case *RevocationMutation:
		return c.Revocation.mutate(ctx, m)
	case *TunnelConfigMutation:
		return c.TunnelConfig.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DeviceClient is a client for the Device schema.
type DeviceClient struct {
	config
}

// NewDeviceClient returns a client for the Device from the given config.
func NewDeviceClient(c config) *DeviceClient {
	return &DeviceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `device.Hooks(f(g(h())))`.
func (c *DeviceClient) Use(hooks ...Hook) {
	c.hooks.Device = append(c.hooks.Device, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `device.Intercept(f(g(h())))`.
func (c *DeviceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Device = append(c.inters.Device, interceptors...)
}

// Create returns a builder for creating a Device entity.
func (c *DeviceClient) Create() *DeviceCreate {
	mutation := newDeviceMutation(c.config, OpCreate)
	return &DeviceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Device entities.
func (c *DeviceClient) CreateBulk(builders ...*DeviceCreate) *DeviceCreateBulk {
	return &DeviceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DeviceClient) MapCreateBulk(slice any, setFunc func(*DeviceCreate, int)) *DeviceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DeviceCreateBulk{err: fmt.Errorf("calling to DeviceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DeviceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DeviceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Device.
func (c *DeviceClient) Update() *DeviceUpdate {
	mutation := newDeviceMutation(c.config, OpUpdate)
	return &DeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeviceClient) UpdateOne(d *Device) *DeviceUpdateOne {
	mutation := newDeviceMutation(c.config, OpUpdateOne, withDevice(d))
	return &DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeviceClient) UpdateOneID(id int) *DeviceUpdateOne {
	mutation := newDeviceMutation(c.config, OpUpdateOne, withDeviceID(id))
	return &DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Device.
func (c *DeviceClient) Delete() *DeviceDelete {
	mutation := newDeviceMutation(c.config, OpDelete)
	return &DeviceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DeviceClient) DeleteOne(d *Device) *DeviceDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DeviceClient) DeleteOneID(id int) *DeviceDeleteOne {
	builder := c.Delete().Where(device.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeviceDeleteOne{builder}
}

// Query returns a query builder for Device.
func (c *DeviceClient) Query() *DeviceQuery {
	return &DeviceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDevice},
		inters: c.Interceptors(),
	}
}

// Get returns a Device entity by its id.
func (c *DeviceClient) Get(ctx context.Context, id int) (*Device, error) {
	return c.Query().Where(device.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeviceClient) GetX(ctx context.Context, id int) *Device {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTunnelConfig queries the tunnel_config edge of a Device.
func (c *DeviceClient) QueryTunnelConfig(d *Device) *TunnelConfigQuery {
	query := (&TunnelConfigClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, id),
			sqlgraph.To(tunnelconfig.Table, tunnelconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, device.TunnelConfigTable, device.TunnelConfigColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeviceClient) Hooks() []Hook {
	return c.hooks.Device
}

// Interceptors returns the client interceptors.
func (c *DeviceClient) Interceptors() []Interceptor {
	return c.inters.Device
}

func (c *DeviceClient) mutate(ctx context.Context, m *DeviceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DeviceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DeviceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Device mutation op: %q", m.Op())
	}
}

// RevocationClient is a client for the Revocation schema.
type RevocationClient struct {
	config
}

// NewRevocationClient returns a client for the Revocation from the given config.
func NewRevocationClient(c config) *RevocationClient {
	return &RevocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `revocation.Hooks(f(g(h())))`.
func (c *RevocationClient) Use(hooks ...Hook) {
	c.hooks.Revocation = append(c.hooks.Revocation, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `revocation.Intercept(f(g(h())))`.
func (c *RevocationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Revocation = append(c.inters.Revocation, interceptors...)
}

// Create returns a builder for creating a Revocation entity.
func (c *RevocationClient) Create() *RevocationCreate {
	mutation := newRevocationMutation(c.config, OpCreate)
	return &RevocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Revocation entities.
func (c *RevocationClient) CreateBulk(builders ...*RevocationCreate) *RevocationCreateBulk {
	return &RevocationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RevocationClient) MapCreateBulk(slice any, setFunc func(*RevocationCreate, int)) *RevocationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RevocationCreateBulk{err: fmt.Errorf("calling to RevocationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RevocationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RevocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Revocation.
func (c *RevocationClient) Update() *RevocationUpdate {
	mutation := newRevocationMutation(c.config, OpUpdate)
	return &RevocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RevocationClient) UpdateOne(r *Revocation) *RevocationUpdateOne {
	mutation := newRevocationMutation(c.config, OpUpdateOne, withRevocation(r))
	return &RevocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RevocationClient) UpdateOneID(id int) *RevocationUpdateOne {
	mutation := newRevocationMutation(c.config, OpUpdateOne, withRevocationID(id))
	return &RevocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Revocation.
func (c *RevocationClient) Delete() *RevocationDelete {
	mutation := newRevocationMutation(c.config, OpDelete)
	return &RevocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RevocationClient) DeleteOne(r *Revocation) *RevocationDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RevocationClient) DeleteOneID(id int) *RevocationDeleteOne {
	builder := c.Delete().Where(revocation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RevocationDeleteOne{builder}
}

// Query returns a query builder for Revocation.
func (c *RevocationClient) Query() *RevocationQuery {
	return &RevocationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRevocation},
		inters: c.Interceptors(),
	}
}

// Get returns a Revocation entity by its id.
func (c *RevocationClient) Get(ctx context.Context, id int) (*Revocation, error) {
	return c.Query().Where(revocation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RevocationClient) GetX(ctx context.Context, id int) *Revocation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RevocationClient) Hooks() []Hook {
	return c.hooks.Revocation
}

// Interceptors returns the client interceptors.
func (c *RevocationClient) Interceptors() []Interceptor {
	return c.inters.Revocation
}

func (c *RevocationClient) mutate(ctx context.Context, m *RevocationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RevocationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RevocationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RevocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RevocationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Revocation mutation op: %q", m.Op())
	}
}

// TunnelConfigClient is a client for the TunnelConfig schema.
type TunnelConfigClient struct {
	config
}

// NewTunnelConfigClient returns a client for the TunnelConfig from the given config.
func NewTunnelConfigClient(c config) *TunnelConfigClient {
	return &TunnelConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tunnelconfig.Hooks(f(g(h())))`.
func (c *TunnelConfigClient) Use(hooks ...Hook) {
	c.hooks.TunnelConfig = append(c.hooks.TunnelConfig, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tunnelconfig.Intercept(f(g(h())))`.
func (c *TunnelConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.TunnelConfig = append(c.inters.TunnelConfig, interceptors...)
}

// Create returns a builder for creating a TunnelConfig entity.
func (c *TunnelConfigClient) Create() *TunnelConfigCreate {
	mutation := newTunnelConfigMutation(c.config, OpCreate)
	return &TunnelConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TunnelConfig entities.
func (c *TunnelConfigClient) CreateBulk(builders ...*TunnelConfigCreate) *TunnelConfigCreateBulk {
	return &TunnelConfigCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TunnelConfigClient) MapCreateBulk(slice any, setFunc func(*TunnelConfigCreate, int)) *TunnelConfigCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TunnelConfigCreateBulk{err: fmt.Errorf("calling to TunnelConfigClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TunnelConfigCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TunnelConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TunnelConfig.
func (c *TunnelConfigClient) Update() *TunnelConfigUpdate {
	mutation := newTunnelConfigMutation(c.config, OpUpdate)
	return &TunnelConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TunnelConfigClient) UpdateOne(tc *TunnelConfig) *TunnelConfigUpdateOne {
	mutation := newTunnelConfigMutation(c.config, OpUpdateOne, withTunnelConfig(tc))
	return &TunnelConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TunnelConfigClient) UpdateOneID(id int) *TunnelConfigUpdateOne {
	mutation := newTunnelConfigMutation(c.config, OpUpdateOne, withTunnelConfigID(id))
	return &TunnelConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TunnelConfig.
func (c *TunnelConfigClient) Delete() *TunnelConfigDelete {
	mutation := newTunnelConfigMutation(c.config, OpDelete)
	return &TunnelConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TunnelConfigClient) DeleteOne(tc *TunnelConfig) *TunnelConfigDeleteOne {
	return c.DeleteOneID(tc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TunnelConfigClient) DeleteOneID(id int) *TunnelConfigDeleteOne {
	builder := c.Delete().Where(tunnelconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TunnelConfigDeleteOne{builder}
}

// Query returns a query builder for TunnelConfig.
func (c *TunnelConfigClient) Query() *TunnelConfigQuery {
	return &TunnelConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTunnelConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a TunnelConfig entity by its id.
func (c *TunnelConfigClient) Get(ctx context.Context, id int) (*TunnelConfig, error) {
	return c.Query().Where(tunnelconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TunnelConfigClient) GetX(ctx context.Context, id int) *TunnelConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDevice queries the device edge of a TunnelConfig.
func (c *TunnelConfigClient) QueryDevice(tc *TunnelConfig) *DeviceQuery {
	query := (&DeviceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tunnelconfig.Table, tunnelconfig.FieldID, id),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, tunnelconfig.DeviceTable, tunnelconfig.DeviceColumn),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TunnelConfigClient) Hooks() []Hook {
	return c.hooks.TunnelConfig
}

// Interceptors returns the client interceptors.
func (c *TunnelConfigClient) Interceptors() []Interceptor {
	return c.inters.TunnelConfig
}

func (c *TunnelConfigClient) mutate(ctx context.Context, m *TunnelConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TunnelConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TunnelConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TunnelConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TunnelConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TunnelConfig mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Device, Revocation, TunnelConfig, User []ent.Hook
	}
	inters struct {
		Device, Revocation, TunnelConfig, User []ent.Interceptor
	}
)
