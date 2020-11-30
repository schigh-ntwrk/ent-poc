// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AppointmentsColumns holds the columns for the "appointments" table.
	AppointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, SchemaType: map[string]string{"mysql": "char(36) binary"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "removed_at", Type: field.TypeTime, Nullable: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "paid_at", Type: field.TypeTime},
		{Name: "charge", Type: field.TypeFloat64},
		{Name: "paid", Type: field.TypeBool},
		{Name: "pet_id", Type: field.TypeUUID, Nullable: true},
		{Name: "veterinarian_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppointmentsTable holds the schema information for the "appointments" table.
	AppointmentsTable = &schema.Table{
		Name:       "appointments",
		Columns:    AppointmentsColumns,
		PrimaryKey: []*schema.Column{AppointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "appointments_pets_appointments",
				Columns: []*schema.Column{AppointmentsColumns[9]},

				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "appointments_veterinarians_appointments",
				Columns: []*schema.Column{AppointmentsColumns[10]},

				RefColumns: []*schema.Column{VeterinariansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "appointment_removed_at",
				Unique:  false,
				Columns: []*schema.Column{AppointmentsColumns[3]},
			},
			{
				Name:    "idx_appointment_time",
				Unique:  false,
				Columns: []*schema.Column{AppointmentsColumns[4], AppointmentsColumns[5]},
			},
			{
				Name:    "idx_appointment_paid",
				Unique:  false,
				Columns: []*schema.Column{AppointmentsColumns[8], AppointmentsColumns[6]},
			},
		},
	}
	// ClinicsColumns holds the columns for the "clinics" table.
	ClinicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, SchemaType: map[string]string{"mysql": "char(36) binary"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "removed_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "address", Type: field.TypeString, Size: 255},
		{Name: "phone", Type: field.TypeString, Size: 20},
		{Name: "web_url", Type: field.TypeString, Size: 255},
	}
	// ClinicsTable holds the schema information for the "clinics" table.
	ClinicsTable = &schema.Table{
		Name:        "clinics",
		Columns:     ClinicsColumns,
		PrimaryKey:  []*schema.Column{ClinicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "clinic_removed_at",
				Unique:  false,
				Columns: []*schema.Column{ClinicsColumns[3]},
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, SchemaType: map[string]string{"mysql": "char(36) binary"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "removed_at", Type: field.TypeTime, Nullable: true},
		{Name: "address", Type: field.TypeString, Size: 255},
		{Name: "phone", Type: field.TypeString, Size: 20},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "customers_users_customer",
				Columns: []*schema.Column{CustomersColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "customer_removed_at",
				Unique:  false,
				Columns: []*schema.Column{CustomersColumns[3]},
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, SchemaType: map[string]string{"mysql": "char(36) binary"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "removed_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 45},
		{Name: "species", Type: field.TypeEnum, Enums: []string{"DOG", "CAT", "BIRD", "RODENT", "LIZARD", "FISH"}},
		{Name: "birth_date", Type: field.TypeTime, Nullable: true},
		{Name: "details", Type: field.TypeString, Nullable: true, Size: 1024},
		{Name: "customer_id", Type: field.TypeUUID, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "pets_customers_pets",
				Columns: []*schema.Column{PetsColumns[8]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "pet_removed_at",
				Unique:  false,
				Columns: []*schema.Column{PetsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, SchemaType: map[string]string{"mysql": "char(36) binary"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "removed_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "first_name", Type: field.TypeString, Size: 45},
		{Name: "last_name", Type: field.TypeString, Size: 45},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_removed_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// VeterinariansColumns holds the columns for the "veterinarians" table.
	VeterinariansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, SchemaType: map[string]string{"mysql": "char(36) binary"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "removed_at", Type: field.TypeTime, Nullable: true},
		{Name: "phone", Type: field.TypeString, Size: 20},
		{Name: "clinic_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// VeterinariansTable holds the schema information for the "veterinarians" table.
	VeterinariansTable = &schema.Table{
		Name:       "veterinarians",
		Columns:    VeterinariansColumns,
		PrimaryKey: []*schema.Column{VeterinariansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "veterinarians_clinics_veterinarians",
				Columns: []*schema.Column{VeterinariansColumns[5]},

				RefColumns: []*schema.Column{ClinicsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "veterinarians_users_veterinarian",
				Columns: []*schema.Column{VeterinariansColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "veterinarian_removed_at",
				Unique:  false,
				Columns: []*schema.Column{VeterinariansColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppointmentsTable,
		ClinicsTable,
		CustomersTable,
		PetsTable,
		UsersTable,
		VeterinariansTable,
	}
)

func init() {
	AppointmentsTable.ForeignKeys[0].RefTable = PetsTable
	AppointmentsTable.ForeignKeys[1].RefTable = VeterinariansTable
	CustomersTable.ForeignKeys[0].RefTable = UsersTable
	PetsTable.ForeignKeys[0].RefTable = CustomersTable
	VeterinariansTable.ForeignKeys[0].RefTable = ClinicsTable
	VeterinariansTable.ForeignKeys[1].RefTable = UsersTable
}