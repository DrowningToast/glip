package config

import (
	"os"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/services"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	InventoryServiceRegistry services.EtcdConfig `envPrefix:"REGISTRY_"`

	InventoryPgConfig       services.PostgresConfig `envPrefix:"INVENTORY_PG_"`
	InventoryRegistryConfig services.EtcdConfig     `envPrefix:"REGISTRY_"`
}

const (
	DefaultConfigPath = "./.env"
)

func Load() (*Config, error) {
	var (
		config     Config
		configPath string
	)

	if os.Getenv("ENV_PATH") != "" {
		configPath = os.Getenv("ENV_PATH")
	} else {
		configPath = DefaultConfigPath
	}

	if err := godotenv.Load(configPath); err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	if err := env.ParseWithOptions(&config, env.Options{
		RequiredIfNoDef: false,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}

	return &config, nil
}
