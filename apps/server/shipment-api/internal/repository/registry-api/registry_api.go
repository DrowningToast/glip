package registry_api

import "github.com/drowningtoast/glip/apps/server/shipment-api/internal/config"

type RegistryApiRepository struct {
	Url    string
	Port   string
	ApiKey string
}

func NewRepository(config *config.ShipmentConfig) *RegistryApiRepository {
	return &RegistryApiRepository{
		Url:    config.RegistryEndpoint,
		Port:   config.RegistryPort,
		ApiKey: config.RegistryApiKey,
	}
}
