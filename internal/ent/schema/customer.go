package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"github.com/schigh-ntwrk/entc-poc/internal/ent/schema/mixin"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").MaxLen(255).NotEmpty(),
		field.String("phone").MaxLen(20).NotEmpty(),
	}
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
		mixin.RemovableEntityTimestamps{},
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("customer").Unique().Required(),
		edge.To("pets", Pet.Type).StorageKey(edge.Column("customer_id")),
	}
}
