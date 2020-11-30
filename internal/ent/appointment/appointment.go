// Code generated by entc, DO NOT EDIT.

package appointment

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appointment type in the database.
	Label = "appointment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRemovedAt holds the string denoting the removed_at field in the database.
	FieldRemovedAt = "removed_at"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldPaidAt holds the string denoting the paid_at field in the database.
	FieldPaidAt = "paid_at"
	// FieldCharge holds the string denoting the charge field in the database.
	FieldCharge = "charge"
	// FieldPaid holds the string denoting the paid field in the database.
	FieldPaid = "paid"

	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeVeterinarians holds the string denoting the veterinarians edge name in mutations.
	EdgeVeterinarians = "veterinarians"

	// Table holds the table name of the appointment in the database.
	Table = "appointments"
	// PetsTable is the table the holds the pets relation/edge.
	PetsTable = "appointments"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "pet_id"
	// VeterinariansTable is the table the holds the veterinarians relation/edge.
	VeterinariansTable = "appointments"
	// VeterinariansInverseTable is the table name for the Veterinarian entity.
	// It exists in this package in order to avoid circular dependency with the "veterinarian" package.
	VeterinariansInverseTable = "veterinarians"
	// VeterinariansColumn is the table column denoting the veterinarians relation/edge.
	VeterinariansColumn = "veterinarian_id"
)

// Columns holds all SQL columns for appointment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRemovedAt,
	FieldStartAt,
	FieldEndAt,
	FieldPaidAt,
	FieldCharge,
	FieldPaid,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Appointment type.
var ForeignKeys = []string{
	"pet_id",
	"veterinarian_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCharge holds the default value on creation for the charge field.
	DefaultCharge float64
	// DefaultPaid holds the default value on creation for the paid field.
	DefaultPaid bool
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
