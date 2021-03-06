// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/migrate"

	"github.com/schigh-ntwrk/ent-poc/internal/ent/appointment"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/clinic"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/customer"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/pet"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/user"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/veterinarian"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Appointment is the client for interacting with the Appointment builders.
	Appointment *AppointmentClient
	// Clinic is the client for interacting with the Clinic builders.
	Clinic *ClinicClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Veterinarian is the client for interacting with the Veterinarian builders.
	Veterinarian *VeterinarianClient
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
	c.Appointment = NewAppointmentClient(c.config)
	c.Clinic = NewClinicClient(c.config)
	c.Customer = NewCustomerClient(c.config)
	c.Pet = NewPetClient(c.config)
	c.User = NewUserClient(c.config)
	c.Veterinarian = NewVeterinarianClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Appointment:  NewAppointmentClient(cfg),
		Clinic:       NewClinicClient(cfg),
		Customer:     NewCustomerClient(cfg),
		Pet:          NewPetClient(cfg),
		User:         NewUserClient(cfg),
		Veterinarian: NewVeterinarianClient(cfg),
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
		config:       cfg,
		Appointment:  NewAppointmentClient(cfg),
		Clinic:       NewClinicClient(cfg),
		Customer:     NewCustomerClient(cfg),
		Pet:          NewPetClient(cfg),
		User:         NewUserClient(cfg),
		Veterinarian: NewVeterinarianClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Appointment.
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
	c.Appointment.Use(hooks...)
	c.Clinic.Use(hooks...)
	c.Customer.Use(hooks...)
	c.Pet.Use(hooks...)
	c.User.Use(hooks...)
	c.Veterinarian.Use(hooks...)
}

// AppointmentClient is a client for the Appointment schema.
type AppointmentClient struct {
	config
}

// NewAppointmentClient returns a client for the Appointment from the given config.
func NewAppointmentClient(c config) *AppointmentClient {
	return &AppointmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `appointment.Hooks(f(g(h())))`.
func (c *AppointmentClient) Use(hooks ...Hook) {
	c.hooks.Appointment = append(c.hooks.Appointment, hooks...)
}

// Create returns a create builder for Appointment.
func (c *AppointmentClient) Create() *AppointmentCreate {
	mutation := newAppointmentMutation(c.config, OpCreate)
	return &AppointmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Appointment entities.
func (c *AppointmentClient) CreateBulk(builders ...*AppointmentCreate) *AppointmentCreateBulk {
	return &AppointmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Appointment.
func (c *AppointmentClient) Update() *AppointmentUpdate {
	mutation := newAppointmentMutation(c.config, OpUpdate)
	return &AppointmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppointmentClient) UpdateOne(a *Appointment) *AppointmentUpdateOne {
	mutation := newAppointmentMutation(c.config, OpUpdateOne, withAppointment(a))
	return &AppointmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppointmentClient) UpdateOneID(id uuid.UUID) *AppointmentUpdateOne {
	mutation := newAppointmentMutation(c.config, OpUpdateOne, withAppointmentID(id))
	return &AppointmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Appointment.
func (c *AppointmentClient) Delete() *AppointmentDelete {
	mutation := newAppointmentMutation(c.config, OpDelete)
	return &AppointmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AppointmentClient) DeleteOne(a *Appointment) *AppointmentDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AppointmentClient) DeleteOneID(id uuid.UUID) *AppointmentDeleteOne {
	builder := c.Delete().Where(appointment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppointmentDeleteOne{builder}
}

// Query returns a query builder for Appointment.
func (c *AppointmentClient) Query() *AppointmentQuery {
	return &AppointmentQuery{config: c.config}
}

// Get returns a Appointment entity by its id.
func (c *AppointmentClient) Get(ctx context.Context, id uuid.UUID) (*Appointment, error) {
	return c.Query().Where(appointment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppointmentClient) GetX(ctx context.Context, id uuid.UUID) *Appointment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPets queries the pets edge of a Appointment.
func (c *AppointmentClient) QueryPets(a *Appointment) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(appointment.Table, appointment.FieldID, id),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appointment.PetsTable, appointment.PetsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVeterinarians queries the veterinarians edge of a Appointment.
func (c *AppointmentClient) QueryVeterinarians(a *Appointment) *VeterinarianQuery {
	query := &VeterinarianQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(appointment.Table, appointment.FieldID, id),
			sqlgraph.To(veterinarian.Table, veterinarian.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appointment.VeterinariansTable, appointment.VeterinariansColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AppointmentClient) Hooks() []Hook {
	return c.hooks.Appointment
}

// ClinicClient is a client for the Clinic schema.
type ClinicClient struct {
	config
}

// NewClinicClient returns a client for the Clinic from the given config.
func NewClinicClient(c config) *ClinicClient {
	return &ClinicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `clinic.Hooks(f(g(h())))`.
func (c *ClinicClient) Use(hooks ...Hook) {
	c.hooks.Clinic = append(c.hooks.Clinic, hooks...)
}

// Create returns a create builder for Clinic.
func (c *ClinicClient) Create() *ClinicCreate {
	mutation := newClinicMutation(c.config, OpCreate)
	return &ClinicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Clinic entities.
func (c *ClinicClient) CreateBulk(builders ...*ClinicCreate) *ClinicCreateBulk {
	return &ClinicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Clinic.
func (c *ClinicClient) Update() *ClinicUpdate {
	mutation := newClinicMutation(c.config, OpUpdate)
	return &ClinicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClinicClient) UpdateOne(cl *Clinic) *ClinicUpdateOne {
	mutation := newClinicMutation(c.config, OpUpdateOne, withClinic(cl))
	return &ClinicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClinicClient) UpdateOneID(id uuid.UUID) *ClinicUpdateOne {
	mutation := newClinicMutation(c.config, OpUpdateOne, withClinicID(id))
	return &ClinicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Clinic.
func (c *ClinicClient) Delete() *ClinicDelete {
	mutation := newClinicMutation(c.config, OpDelete)
	return &ClinicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ClinicClient) DeleteOne(cl *Clinic) *ClinicDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ClinicClient) DeleteOneID(id uuid.UUID) *ClinicDeleteOne {
	builder := c.Delete().Where(clinic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClinicDeleteOne{builder}
}

// Query returns a query builder for Clinic.
func (c *ClinicClient) Query() *ClinicQuery {
	return &ClinicQuery{config: c.config}
}

// Get returns a Clinic entity by its id.
func (c *ClinicClient) Get(ctx context.Context, id uuid.UUID) (*Clinic, error) {
	return c.Query().Where(clinic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClinicClient) GetX(ctx context.Context, id uuid.UUID) *Clinic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVeterinarians queries the veterinarians edge of a Clinic.
func (c *ClinicClient) QueryVeterinarians(cl *Clinic) *VeterinarianQuery {
	query := &VeterinarianQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clinic.Table, clinic.FieldID, id),
			sqlgraph.To(veterinarian.Table, veterinarian.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinic.VeterinariansTable, clinic.VeterinariansColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClinicClient) Hooks() []Hook {
	return c.hooks.Clinic
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Create returns a create builder for Customer.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id uuid.UUID) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CustomerClient) DeleteOneID(id uuid.UUID) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{config: c.config}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id uuid.UUID) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Customer.
func (c *CustomerClient) QueryUser(cu *Customer) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, customer.UserTable, customer.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPets queries the pets edge of a Customer.
func (c *CustomerClient) QueryPets(cu *Customer) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.PetsTable, customer.PetsColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pet.Hooks(f(g(h())))`.
func (c *PetClient) Use(hooks ...Hook) {
	c.hooks.Pet = append(c.hooks.Pet, hooks...)
}

// Create returns a create builder for Pet.
func (c *PetClient) Create() *PetCreate {
	mutation := newPetMutation(c.config, OpCreate)
	return &PetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Pet entities.
func (c *PetClient) CreateBulk(builders ...*PetCreate) *PetCreateBulk {
	return &PetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	mutation := newPetMutation(c.config, OpUpdate)
	return &PetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPet(pe))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id uuid.UUID) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPetID(id))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	mutation := newPetMutation(c.config, OpDelete)
	return &PetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PetClient) DeleteOneID(id uuid.UUID) *PetDeleteOne {
	builder := c.Delete().Where(pet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PetDeleteOne{builder}
}

// Query returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{config: c.config}
}

// Get returns a Pet entity by its id.
func (c *PetClient) Get(ctx context.Context, id uuid.UUID) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id uuid.UUID) *Pet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Pet.
func (c *PetClient) QueryOwner(pe *Pet) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pet.OwnerTable, pet.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAppointments queries the appointments edge of a Pet.
func (c *PetClient) QueryAppointments(pe *Pet) *AppointmentQuery {
	query := &AppointmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, id),
			sqlgraph.To(appointment.Table, appointment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pet.AppointmentsTable, pet.AppointmentsColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PetClient) Hooks() []Hook {
	return c.hooks.Pet
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

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVeterinarian queries the veterinarian edge of a User.
func (c *UserClient) QueryVeterinarian(u *User) *VeterinarianQuery {
	query := &VeterinarianQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(veterinarian.Table, veterinarian.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.VeterinarianTable, user.VeterinarianColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomer queries the customer edge of a User.
func (c *UserClient) QueryCustomer(u *User) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.CustomerTable, user.CustomerColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VeterinarianClient is a client for the Veterinarian schema.
type VeterinarianClient struct {
	config
}

// NewVeterinarianClient returns a client for the Veterinarian from the given config.
func NewVeterinarianClient(c config) *VeterinarianClient {
	return &VeterinarianClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `veterinarian.Hooks(f(g(h())))`.
func (c *VeterinarianClient) Use(hooks ...Hook) {
	c.hooks.Veterinarian = append(c.hooks.Veterinarian, hooks...)
}

// Create returns a create builder for Veterinarian.
func (c *VeterinarianClient) Create() *VeterinarianCreate {
	mutation := newVeterinarianMutation(c.config, OpCreate)
	return &VeterinarianCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Veterinarian entities.
func (c *VeterinarianClient) CreateBulk(builders ...*VeterinarianCreate) *VeterinarianCreateBulk {
	return &VeterinarianCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Veterinarian.
func (c *VeterinarianClient) Update() *VeterinarianUpdate {
	mutation := newVeterinarianMutation(c.config, OpUpdate)
	return &VeterinarianUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VeterinarianClient) UpdateOne(v *Veterinarian) *VeterinarianUpdateOne {
	mutation := newVeterinarianMutation(c.config, OpUpdateOne, withVeterinarian(v))
	return &VeterinarianUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VeterinarianClient) UpdateOneID(id uuid.UUID) *VeterinarianUpdateOne {
	mutation := newVeterinarianMutation(c.config, OpUpdateOne, withVeterinarianID(id))
	return &VeterinarianUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Veterinarian.
func (c *VeterinarianClient) Delete() *VeterinarianDelete {
	mutation := newVeterinarianMutation(c.config, OpDelete)
	return &VeterinarianDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VeterinarianClient) DeleteOne(v *Veterinarian) *VeterinarianDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VeterinarianClient) DeleteOneID(id uuid.UUID) *VeterinarianDeleteOne {
	builder := c.Delete().Where(veterinarian.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VeterinarianDeleteOne{builder}
}

// Query returns a query builder for Veterinarian.
func (c *VeterinarianClient) Query() *VeterinarianQuery {
	return &VeterinarianQuery{config: c.config}
}

// Get returns a Veterinarian entity by its id.
func (c *VeterinarianClient) Get(ctx context.Context, id uuid.UUID) (*Veterinarian, error) {
	return c.Query().Where(veterinarian.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VeterinarianClient) GetX(ctx context.Context, id uuid.UUID) *Veterinarian {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Veterinarian.
func (c *VeterinarianClient) QueryUser(v *Veterinarian) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(veterinarian.Table, veterinarian.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, veterinarian.UserTable, veterinarian.UserColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryClinic queries the clinic edge of a Veterinarian.
func (c *VeterinarianClient) QueryClinic(v *Veterinarian) *ClinicQuery {
	query := &ClinicQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(veterinarian.Table, veterinarian.FieldID, id),
			sqlgraph.To(clinic.Table, clinic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, veterinarian.ClinicTable, veterinarian.ClinicColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAppointments queries the appointments edge of a Veterinarian.
func (c *VeterinarianClient) QueryAppointments(v *Veterinarian) *AppointmentQuery {
	query := &AppointmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(veterinarian.Table, veterinarian.FieldID, id),
			sqlgraph.To(appointment.Table, appointment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, veterinarian.AppointmentsTable, veterinarian.AppointmentsColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VeterinarianClient) Hooks() []Hook {
	return c.hooks.Veterinarian
}
