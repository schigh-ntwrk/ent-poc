// Code generated by entc, DO NOT EDIT.

package clinic

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the clinic type in the database.
	Label = "clinic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRemovedAt holds the string denoting the removed_at field in the database.
	FieldRemovedAt = "removed_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldWebURL holds the string denoting the web_url field in the database.
	FieldWebURL = "web_url"

	// EdgeVeterinarians holds the string denoting the veterinarians edge name in mutations.
	EdgeVeterinarians = "veterinarians"

	// Table holds the table name of the clinic in the database.
	Table = "clinics"
	// VeterinariansTable is the table the holds the veterinarians relation/edge.
	VeterinariansTable = "veterinarians"
	// VeterinariansInverseTable is the table name for the Veterinarian entity.
	// It exists in this package in order to avoid circular dependency with the "veterinarian" package.
	VeterinariansInverseTable = "veterinarians"
	// VeterinariansColumn is the table column denoting the veterinarians relation/edge.
	VeterinariansColumn = "clinic_id"
)

// Columns holds all SQL columns for clinic fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRemovedAt,
	FieldName,
	FieldAddress,
	FieldPhone,
	FieldWebURL,
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

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// WebURLValidator is a validator for the "web_url" field. It is called by the builders before save.
	WebURLValidator func(string) error
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
