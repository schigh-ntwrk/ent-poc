package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/domain/clinic"
	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type ClinicService interface {
	GetByID(context.Context, uuid.UUID) (*ent.Clinic, error)
	Save(context.Context, *ent.Clinic) (*ent.Clinic, error)
}

var _ ClinicService = (*clinic.Service)(nil)

func NewClinic(client *ent.Client) (ClinicService, error) {
	return clinic.New(clinic.WithENTClient(client.Clinic)), nil
}
