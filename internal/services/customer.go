package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/domain/customer"
	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type CustomerService interface {
	GetByID(context.Context, uuid.UUID) (*ent.Customer, error)
}

var _ CustomerService = (*customer.Service)(nil)

func NewCustomer(client *ent.Client) (CustomerService, error) {
	return customer.New(customer.WithENTClient(client.Customer)), nil
}
