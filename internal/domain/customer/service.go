package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type Service struct {
	client *ent.CustomerClient
}

type Option func(*Service)

func WithENTClient(client *ent.CustomerClient) Option {
	return func(s *Service) {
		s.client = client
	}
}

func New(opts ...Option) *Service {
	s := Service{}
	for _, f := range opts {
		f(&s)
	}

	return &s
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*ent.Customer, error) {
	return s.client.Get(ctx, id)
}
