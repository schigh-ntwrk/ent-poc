// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/appointment"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/pet"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/veterinarian"
)

// Appointment is the model entity for the Appointment schema.
type Appointment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RemovedAt holds the value of the "removed_at" field.
	RemovedAt *time.Time `json:"removed_at,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt time.Time `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt time.Time `json:"end_at,omitempty"`
	// PaidAt holds the value of the "paid_at" field.
	PaidAt time.Time `json:"paid_at,omitempty"`
	// Charge holds the value of the "charge" field.
	Charge float64 `json:"charge,omitempty"`
	// Paid holds the value of the "paid" field.
	Paid bool `json:"paid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppointmentQuery when eager-loading is set.
	Edges           AppointmentEdges `json:"edges"`
	pet_id          *uuid.UUID
	veterinarian_id *uuid.UUID
}

// AppointmentEdges holds the relations/edges for other nodes in the graph.
type AppointmentEdges struct {
	// Pets holds the value of the pets edge.
	Pets *Pet
	// Veterinarians holds the value of the veterinarians edge.
	Veterinarians *Veterinarian
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppointmentEdges) PetsOrErr() (*Pet, error) {
	if e.loadedTypes[0] {
		if e.Pets == nil {
			// The edge pets was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// VeterinariansOrErr returns the Veterinarians value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppointmentEdges) VeterinariansOrErr() (*Veterinarian, error) {
	if e.loadedTypes[1] {
		if e.Veterinarians == nil {
			// The edge veterinarians was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: veterinarian.Label}
		}
		return e.Veterinarians, nil
	}
	return nil, &NotLoadedError{edge: "veterinarians"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Appointment) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},       // id
		&sql.NullTime{},    // created_at
		&sql.NullTime{},    // updated_at
		&sql.NullTime{},    // removed_at
		&sql.NullTime{},    // start_at
		&sql.NullTime{},    // end_at
		&sql.NullTime{},    // paid_at
		&sql.NullFloat64{}, // charge
		&sql.NullBool{},    // paid
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Appointment) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // pet_id
		&uuid.UUID{}, // veterinarian_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Appointment fields.
func (a *Appointment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(appointment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		a.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		a.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field removed_at", values[2])
	} else if value.Valid {
		a.RemovedAt = new(time.Time)
		*a.RemovedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field start_at", values[3])
	} else if value.Valid {
		a.StartAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field end_at", values[4])
	} else if value.Valid {
		a.EndAt = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field paid_at", values[5])
	} else if value.Valid {
		a.PaidAt = value.Time
	}
	if value, ok := values[6].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field charge", values[6])
	} else if value.Valid {
		a.Charge = value.Float64
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field paid", values[7])
	} else if value.Valid {
		a.Paid = value.Bool
	}
	values = values[8:]
	if len(values) == len(appointment.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field pet_id", values[0])
		} else if value != nil {
			a.pet_id = value
		}
		if value, ok := values[1].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field veterinarian_id", values[1])
		} else if value != nil {
			a.veterinarian_id = value
		}
	}
	return nil
}

// QueryPets queries the pets edge of the Appointment.
func (a *Appointment) QueryPets() *PetQuery {
	return (&AppointmentClient{config: a.config}).QueryPets(a)
}

// QueryVeterinarians queries the veterinarians edge of the Appointment.
func (a *Appointment) QueryVeterinarians() *VeterinarianQuery {
	return (&AppointmentClient{config: a.config}).QueryVeterinarians(a)
}

// Update returns a builder for updating this Appointment.
// Note that, you need to call Appointment.Unwrap() before calling this method, if this Appointment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Appointment) Update() *AppointmentUpdateOne {
	return (&AppointmentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Appointment) Unwrap() *Appointment {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Appointment is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Appointment) String() string {
	var builder strings.Builder
	builder.WriteString("Appointment(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	if v := a.RemovedAt; v != nil {
		builder.WriteString(", removed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", start_at=")
	builder.WriteString(a.StartAt.Format(time.ANSIC))
	builder.WriteString(", end_at=")
	builder.WriteString(a.EndAt.Format(time.ANSIC))
	builder.WriteString(", paid_at=")
	builder.WriteString(a.PaidAt.Format(time.ANSIC))
	builder.WriteString(", charge=")
	builder.WriteString(fmt.Sprintf("%v", a.Charge))
	builder.WriteString(", paid=")
	builder.WriteString(fmt.Sprintf("%v", a.Paid))
	builder.WriteByte(')')
	return builder.String()
}

// Appointments is a parsable slice of Appointment.
type Appointments []*Appointment

func (a Appointments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
