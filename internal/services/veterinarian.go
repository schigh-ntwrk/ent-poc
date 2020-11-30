package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/domain/veterinarian"
	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type VeterinarianService interface {
	GetByID(context.Context, uuid.UUID) (*ent.Veterinarian, error)
}

var _ VeterinarianService = (*veterinarian.Service)(nil)

func NewVeterinarian(client *ent.Client) (VeterinarianService, error) {
	return veterinarian.New(veterinarian.WithENTClient(client.Veterinarian)), nil
}
