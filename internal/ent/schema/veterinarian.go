package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"github.com/schigh-ntwrk/entc-poc/internal/ent/schema/mixin"
)

// Veterinarian holds the schema definition for the Veterinarian entity.
type Veterinarian struct {
	ent.Schema
}

// Fields of the Veterinarian.
func (Veterinarian) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").MaxLen(20).NotEmpty(),
	}
}

func (Veterinarian) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
		mixin.RemovableEntityTimestamps{},
	}
}

// Edges of the Veterinarian.
func (Veterinarian) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("veterinarian").Unique().Required(),
		edge.From("clinic", Clinic.Type).Ref("veterinarians").Unique().Required(),
		edge.To("appointments", Appointment.Type).StorageKey(edge.Column("veterinarian_id")),
	}
}
