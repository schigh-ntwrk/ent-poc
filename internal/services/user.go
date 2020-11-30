package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/domain/user"
	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type UserService interface {
	GetByID(context.Context, uuid.UUID) (*ent.User, error)
}

var _ UserService = (*user.Service)(nil)

func NewUser(client *ent.Client) (UserService, error) {
	return user.New(user.WithENTClient(client.User)), nil
}
