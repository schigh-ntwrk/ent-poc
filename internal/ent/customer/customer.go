// Code generated by entc, DO NOT EDIT.

package customer

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRemovedAt holds the string denoting the removed_at field in the database.
	FieldRemovedAt = "removed_at"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"

	// Table holds the table name of the customer in the database.
	Table = "customers"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "customers"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// PetsTable is the table the holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "customer_id"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRemovedAt,
	FieldAddress,
	FieldPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Customer type.
var ForeignKeys = []string{
	"user_id",
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
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)