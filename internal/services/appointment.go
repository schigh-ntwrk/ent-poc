package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/domain/appointment"
	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type AppointmentService interface {
	GetByID(context.Context, uuid.UUID) (*ent.Appointment, error)
}

var _ AppointmentService = (*appointment.Service)(nil)

func NewAppointment(client *ent.Client) (AppointmentService, error) {
	return appointment.New(appointment.WithENTClient(client.Appointment)), nil
}
