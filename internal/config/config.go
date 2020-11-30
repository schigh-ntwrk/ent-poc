package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database Database
	Server   Server
}

type Database struct {
	DBHost  string `envconfig:"DB_HOST" required:"true"`
	DBName  string `envconfig:"DB_NAME" required:"true"`
	DBUser  string `envconfig:"DB_USER" required:"true"`
	DBPass  string `envconfig:"DB_PASS" required:"true"`
	Migrate bool   `envconfig:"DB_RUN_MIGRATION"`
}

type Server struct {
	IngressTimeout  time.Duration `envconfig:"SERVER_INGRESS_TIMEOUT" default:"1s"`
	ShutdownTimeout time.Duration `envconfig:"SERVER_SHUTDOWN_TIMEOUT" default:"5s"`
	ServicePort     int           `envconfig:"SERVER_SERVICE_PORT" default:"80"`
}

func (d Database) DSN() string {
	return fmt.Sprintf("%s:%s@/%s?parseTime=true", d.DBUser, d.DBPass, d.DBName)
}

func New() (*Config, error) {
	// This will load env files into the environment and set them as variables.
	// If no value is set for this variable, the file `.env` will be used as a
	// default.  If that file is not found, the godotenv utility will fail
	// silently.  Y can specify multiple env files by comma-separating them and
	// storing them into ENV_FILE.  They will be processed in the order they are
	// listed.
	envFile, ok := os.LookupEnv("ENV_FILE")
	if !ok || envFile == "" {
		pwd, pwdErr := os.Getwd()
		if pwdErr != nil {
			return nil, pwdErr
		}
		envFile = filepath.Join(pwd, ".env")
	}
	envFiles := strings.Split(envFile, ",")
	_ = godotenv.Load(envFiles...)

	// This will parse environment variables into a Config struct
	var out Config
	if err := envconfig.Process("", &out); err != nil {
		return nil, err
	}

	return &out, nil
}
