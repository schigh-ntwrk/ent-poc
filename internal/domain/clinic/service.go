package clinic

import (
	"context"

	"github.com/google/uuid"

	"github.com/schigh-ntwrk/ent-poc/internal/ent"
)

type Service struct {
	client *ent.ClinicClient
}

type Option func(*Service)

func WithENTClient(client *ent.ClinicClient) Option {
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

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*ent.Clinic, error) {
	return s.client.Get(ctx, id)
}

func (s *Service) Save(ctx context.Context, clinic *ent.Clinic) (*ent.Clinic, error) {
	if uuidZero(clinic.ID) {
		id, idErr := uuid.NewUUID()
		if idErr != nil {
			return nil, idErr
		}

		return s.client.Create().
			SetID(id).
			SetName(clinic.Name).
			SetAddress(clinic.Address).
			SetPhone(clinic.Phone).
			SetWebURL(clinic.WebURL).Save(ctx)
	}

	return s.client.UpdateOne(clinic).
		SetName(clinic.Name).
		SetAddress(clinic.Address).
		SetPhone(clinic.Phone).
		SetWebURL(clinic.WebURL).
		Save(ctx)
}

func uuidZero(id uuid.UUID) bool {
	// This compares a uuid to its zero value.
	// The only way a UUID can be equivalent to
	// its zero value is if it was not instantiated
	// via a New... func
	return id[0]|id[1]|id[2]|id[3]|id[4]|id[5]|id[6]|id[7]|id[8]|id[9]|id[10]|id[11]|id[12]|id[13]|id[14]|id[15] == 0
}
