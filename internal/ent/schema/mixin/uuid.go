package mixin

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

type UUID struct {
	ent.Mixin
}

func (UUID) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).SchemaType(map[string]string{
			dialect.MySQL: "char(36) binary",
		}),
	}
}

func (UUID) Indexes() []ent.Index {
	return nil
}

func (UUID) Edges() []ent.Edge {
	return nil
}

func (UUID) Hooks() []ent.Hook {
	return nil
}

func (UUID) Policy() ent.Policy {
	return nil
}

func (UUID) Annotations() []schema.Annotation {
	return nil
}
