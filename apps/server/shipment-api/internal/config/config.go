package config

import (
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/config"
	"github.com/drowningtoast/glip/apps/server/internal/services"
	"github.com/joho/godotenv"
)

type ShipmentConfig struct {
	config.Config

	// static config
	WarehouseRegions WarehouseRegions
	WarehouseRoutes  WarehouseRoutes

	// API config
	RegistryApiKey string `env:"SHIPMENT_REGISTRY_API_KEY"`

	RabbitMQConfig     services.RabbitMQConfig `envPrefix:"SHIPMENT_RABBITMQ_"`
	ShipmentAuthConfig config.AuthConfig       `envPrefix:"SHIPMENT_AUTH_"`
	ShipmentPgConfig   services.PostgresConfig `envPrefix:"SHIPMENT_PG_"`
}

const (
	DefaultConfigPath = "./.env"
)

func ExtendConfig(config *config.Config, path *string) (*ShipmentConfig, error) {
	if config == nil {
		return nil, errors.New("config cannot be nil")
	}

	var (
		extendedConfig ShipmentConfig
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

	// load local config
	warehouseRegions, err := LoadWarehouseConfig()
	if err != nil {
		return nil, err
	}
	extendedConfig.WarehouseRegions = *warehouseRegions
	warehouseRoutes, err := GetWarehouseRoutes()
	if err != nil {
		return nil, err
	}
	extendedConfig.WarehouseRoutes = *warehouseRoutes

	return &extendedConfig, nil
}
