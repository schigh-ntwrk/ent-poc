package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/ent-poc/internal/domain/pet"
	"github.com/schigh-ntwrk/ent-poc/internal/ent"
)

type PetService interface {
	GetByID(context.Context, uuid.UUID) (*ent.Pet, error)
}

var _ PetService = (*pet.Service)(nil)

func NewPet(client *ent.Client) (PetService, error) {
	return pet.New(pet.WithENTClient(client.Pet)), nil
}
