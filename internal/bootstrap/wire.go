// +build wireinject

package bootstrap

import (
	"context"

	"github.com/google/wire"

	"github.com/schigh-ntwrk/ent-poc/internal/config"
	"github.com/schigh-ntwrk/ent-poc/internal/services"
)

func Up(ctx context.Context) (Dependencies, error) {
	wire.Build(
		NewDependencies,
		config.New,
		NewEntClient,
		services.NewClinic,
		services.NewAppointment,
		services.NewPet,
		services.NewUser,
		services.NewVeterinarian,
		services.NewCustomer,
		NewServer,
	)
	return Dependencies{}, nil
}
