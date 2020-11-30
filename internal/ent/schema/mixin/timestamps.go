package mixin

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

type BaseEntityTimestamps struct {
	ent.Mixin
}

func (BaseEntityTimestamps) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (BaseEntityTimestamps) Indexes() []ent.Index {
	return nil
}

func (BaseEntityTimestamps) Edges() []ent.Edge {
	return nil
}

func (BaseEntityTimestamps) Hooks() []ent.Hook {
	return nil
}

func (BaseEntityTimestamps) Policy() ent.Policy {
	return nil
}

func (BaseEntityTimestamps) Annotations() []schema.Annotation {
	return nil
}

// -------------------------

type RemovableEntityTimestamps struct {
	ent.Mixin
}

func (RemovableEntityTimestamps) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("removed_at").Nillable().Optional(),
	}
}

func (RemovableEntityTimestamps) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("removed_at"),
	}
}

func (RemovableEntityTimestamps) Edges() []ent.Edge {
	return nil
}

func (RemovableEntityTimestamps) Hooks() []ent.Hook {
	return nil
}

func (RemovableEntityTimestamps) Policy() ent.Policy {
	return nil
}

func (RemovableEntityTimestamps) Annotations() []schema.Annotation {
	return nil
}
