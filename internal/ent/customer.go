// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/schigh-ntwrk/entc-poc/internal/ent/customer"
	"github.com/schigh-ntwrk/entc-poc/internal/ent/user"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RemovedAt holds the value of the "removed_at" field.
	RemovedAt *time.Time `json:"removed_at,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges   CustomerEdges `json:"edges"`
	user_id *uuid.UUID
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// User holds the value of the user edge.
	User *User
	// Pets holds the value of the pets edge.
	Pets []*Pet
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) PetsOrErr() ([]*Pet, error) {
	if e.loadedTypes[1] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // removed_at
		&sql.NullString{}, // address
		&sql.NullString{}, // phone
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Customer) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		c.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		c.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		c.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field removed_at", values[2])
	} else if value.Valid {
		c.RemovedAt = new(time.Time)
		*c.RemovedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[3])
	} else if value.Valid {
		c.Address = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[4])
	} else if value.Valid {
		c.Phone = value.String
	}
	values = values[5:]
	if len(values) == len(customer.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field user_id", values[0])
		} else if value != nil {
			c.user_id = value
		}
	}
	return nil
}

// QueryUser queries the user edge of the Customer.
func (c *Customer) QueryUser() *UserQuery {
	return (&CustomerClient{config: c.config}).QueryUser(c)
}

// QueryPets queries the pets edge of the Customer.
func (c *Customer) QueryPets() *PetQuery {
	return (&CustomerClient{config: c.config}).QueryPets(c)
}

// Update returns a builder for updating this Customer.
// Note that, you need to call Customer.Unwrap() before calling this method, if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	if v := c.RemovedAt; v != nil {
		builder.WriteString(", removed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", address=")
	builder.WriteString(c.Address)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
