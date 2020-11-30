package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"github.com/schigh-ntwrk/ent-poc/internal/ent/schema/mixin"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(45).NotEmpty(),
		field.Enum("species").NamedValues("Dog", "DOG", "Cat", "CAT", "Bird", "BIRD", "Rodent", "RODENT", "Lizard", "LIZARD", "Fish", "FISH"),
		field.Time("birth_date").Nillable().Optional(),
		field.String("details").MaxLen(1024).Nillable().Optional(),
	}
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
		mixin.RemovableEntityTimestamps{},
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Customer.Type).Ref("pets").Unique(),
		edge.To("appointments", Appointment.Type).StorageKey(edge.Column("pet_id")),
	}
}
