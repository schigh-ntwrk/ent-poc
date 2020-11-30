package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"github.com/schigh-ntwrk/entc-poc/internal/ent/schema/mixin"
)

// Clinic holds the schema definition for the Clinic entity.
type Clinic struct {
	ent.Schema
}

// Fields of the Clinic.
func (Clinic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(64).NotEmpty(),
		field.String("address").MaxLen(255).NotEmpty(),
		field.String("phone").MaxLen(20).NotEmpty(),
		field.String("web_url").MaxLen(255),
	}
}

func (Clinic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
		mixin.RemovableEntityTimestamps{},
	}
}

// Edges of the Clinic.
func (Clinic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("veterinarians", Veterinarian.Type).StorageKey(edge.Column("clinic_id")),
	}
}
