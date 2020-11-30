// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/customer"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/user"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/veterinarian"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RemovedAt holds the value of the "removed_at" field.
	RemovedAt *time.Time `json:"removed_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Veterinarian holds the value of the veterinarian edge.
	Veterinarian *Veterinarian
	// Customer holds the value of the customer edge.
	Customer *Customer
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VeterinarianOrErr returns the Veterinarian value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) VeterinarianOrErr() (*Veterinarian, error) {
	if e.loadedTypes[0] {
		if e.Veterinarian == nil {
			// The edge veterinarian was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: veterinarian.Label}
		}
		return e.Veterinarian, nil
	}
	return nil, &NotLoadedError{edge: "veterinarian"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[1] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // removed_at
		&sql.NullString{}, // email
		&sql.NullString{}, // first_name
		&sql.NullString{}, // last_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		u.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field removed_at", values[2])
	} else if value.Valid {
		u.RemovedAt = new(time.Time)
		*u.RemovedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[3])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field first_name", values[4])
	} else if value.Valid {
		u.FirstName = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field last_name", values[5])
	} else if value.Valid {
		u.LastName = value.String
	}
	return nil
}

// QueryVeterinarian queries the veterinarian edge of the User.
func (u *User) QueryVeterinarian() *VeterinarianQuery {
	return (&UserClient{config: u.config}).QueryVeterinarian(u)
}

// QueryCustomer queries the customer edge of the User.
func (u *User) QueryCustomer() *CustomerQuery {
	return (&UserClient{config: u.config}).QueryCustomer(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	if v := u.RemovedAt; v != nil {
		builder.WriteString(", removed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(u.LastName)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
