package config

import (
	"os"

	"github.com/drowningtoast/glip/apps/server/services"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	InventoryPgConfig services.PostgresConfig `envPrefix:"INVENTORY_PG"`
	ShipmentPgConfig  services.PostgresConfig `envPrefix:"SHIPMENT_PG"`
}

const (
	DefaultConfigPath = "./.env"
)

func LoadConfig() *Config {
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
		panic(err)
	}

	if err := env.ParseWithOptions(&config, env.Options{
		RequiredIfNoDef: false,
	}); err != nil {
		panic(err)
	}
	return &config
}
