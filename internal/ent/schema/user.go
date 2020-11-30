package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"github.com/schigh-ntwrk/ent-poc/internal/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").MaxLen(255).NotEmpty().Unique(),
		field.String("first_name").MaxLen(45).NotEmpty(),
		field.String("last_name").MaxLen(45).NotEmpty(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
		mixin.RemovableEntityTimestamps{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("veterinarian", Veterinarian.Type).Unique().StorageKey(edge.Column("user_id")),
		edge.To("customer", Customer.Type).Unique().StorageKey(edge.Column("user_id")),
	}
}
