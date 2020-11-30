package veterinarian

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/entc-poc/internal/ent"
)

type Service struct {
	client *ent.VeterinarianClient
}

type Option func(*Service)

func WithENTClient(client *ent.VeterinarianClient) Option {
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

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*ent.Veterinarian, error) {
	return s.client.Get(ctx, id)
}
