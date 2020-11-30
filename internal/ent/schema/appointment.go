package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"

	"github.com/schigh-ntwrk/ent-poc/internal/ent/schema/mixin"
)

// Appointment holds the schema definition for the Appointment entity.
type Appointment struct {
	ent.Schema
}

// Fields of the Appointment.
func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_at"),
		field.Time("end_at"),
		field.Time("paid_at"),
		field.Float("charge").Default(0),
		field.Bool("paid").Default(false),
	}
}

func (Appointment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
		mixin.RemovableEntityTimestamps{},
	}
}

func (Appointment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("start_at", "end_at").StorageKey("idx_appointment_time"),
		index.Fields("paid", "paid_at").StorageKey("idx_appointment_paid"),
	}
}

// Edges of the Appointment.
func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pets", Pet.Type).Ref("appointments").Unique().Required(),
		edge.From("veterinarians", Veterinarian.Type).Ref("appointments").Unique().Required(),
	}
}
