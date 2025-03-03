package config

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/drowningtoast/glip/apps/server/internal/config"
	"github.com/drowningtoast/glip/apps/server/internal/services"
	"github.com/joho/godotenv"
)

type RegistryConfig struct {
	config.Config

	RegistryAuthConfig RegistryAuthConfig      `envPrefix:"INVENTORY_REGISTRY_AUTH_"`
	RegistryPgConfig   services.PostgresConfig `envPrefix:"INVENTORY_REGISTRY_PG_"`
}

const (
	DefaultConfigPath = "./.env"
)

func ExtendConfig(config *config.Config, path *string) (*RegistryConfig, error) {
	if config == nil {
		return nil, errors.New("config cannot be nil")
	}

	var (
		extendedConfig RegistryConfig
		configPath     string
	)

	if path != nil {
		configPath = *path
	} else if os.Getenv("ENV_PATH") != "" {
		configPath = os.Getenv("ENV_PATH")
	} else {
		configPath = DefaultConfigPath
	}

	extendedConfig.Config = *config
	if err := godotenv.Load(configPath); err != nil {
		return nil, err
	}

	if err := env.ParseWithOptions(&extendedConfig, env.Options{
		RequiredIfNoDef: false,
	}); err != nil {
		return nil, err
	}

	return &extendedConfig, nil
}
