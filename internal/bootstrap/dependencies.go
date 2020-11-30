package bootstrap

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql" // needed for mysql driver

	"github.com/schigh-ntwrk/ent-poc/internal/config"
	appointmentHTTP "github.com/schigh-ntwrk/ent-poc/internal/domain/appointment/transport/httpx"
	clinicHTTP "github.com/schigh-ntwrk/ent-poc/internal/domain/clinic/transport/httpx"
	customerHTTP "github.com/schigh-ntwrk/ent-poc/internal/domain/customer/transport/httpx"
	petHTTP "github.com/schigh-ntwrk/ent-poc/internal/domain/pet/transport/httpx"
	userHTTP "github.com/schigh-ntwrk/ent-poc/internal/domain/user/transport/httpx"
	veterinarianHTTP "github.com/schigh-ntwrk/ent-poc/internal/domain/veterinarian/transport/httpx"
	"github.com/schigh-ntwrk/ent-poc/internal/ent"
	"github.com/schigh-ntwrk/ent-poc/internal/httpx"
	"github.com/schigh-ntwrk/ent-poc/internal/services"
)

// Dependencies is a bootstrapping artifact that aggregates all dependent
// services upon startup.  This struct exists strictly as an output for the
// dependency injection framework we are using in this demo (wire)
type Dependencies struct {
	Config              *config.Config
	EntClient           *ent.Client
	Server              *httpx.Server
	ClinicService       services.ClinicService
	AppointmentService  services.AppointmentService
	PetService          services.PetService
	UserService         services.UserService
	CustomerService     services.CustomerService
	VeterinarianService services.VeterinarianService
}

// NewDependencies is another ugly func that exists as an
// entrypoint for the dependency injector
func NewDependencies(
	cfg *config.Config,
	entClient *ent.Client,
	server *httpx.Server,
	clinicService services.ClinicService,
	appointmentService services.AppointmentService,
	petService services.PetService,
	userService services.UserService,
	customerService services.CustomerService,
	veterinarianService services.VeterinarianService,
) Dependencies {
	return Dependencies{
		Config:              cfg,
		EntClient:           entClient,
		Server:              server,
		ClinicService:       clinicService,
		AppointmentService:  appointmentService,
		PetService:          petService,
		UserService:         userService,
		CustomerService:     customerService,
		VeterinarianService: veterinarianService,
	}
}

// NewEntClient wraps the creation of ent.Client
func NewEntClient(ctx context.Context, cfg *config.Config) (*ent.Client, error) {
	client, clientErr := ent.Open("mysql", cfg.Database.DSN())
	if clientErr != nil {
		return nil, clientErr
	}

	if cfg.Database.Migrate {
		_ = client.Schema.WriteTo(ctx, os.Stderr)
		return client, client.Schema.Create(ctx)
	}

	return client, nil
}

// NewServer is used to allow wire to properly instantiate the application server
func NewServer(cfg *config.Config,
	clinicService services.ClinicService,
	userService services.UserService,
	customerService services.CustomerService,
	veterinarianService services.VeterinarianService,
	petService services.PetService,
	appointmentService services.AppointmentService) (*httpx.Server, error) {
	return httpx.New(
		httpx.WithConfig(cfg),
		httpx.WithRoutes(clinicHTTP.Routes(clinicService)...),
		httpx.WithRoutes(userHTTP.Routes(userService)...),
		httpx.WithRoutes(customerHTTP.Routes(customerService)...),
		httpx.WithRoutes(veterinarianHTTP.Routes(veterinarianService)...),
		httpx.WithRoutes(petHTTP.Routes(petService)...),
		httpx.WithRoutes(appointmentHTTP.Routes(appointmentService)...),
	)
}
