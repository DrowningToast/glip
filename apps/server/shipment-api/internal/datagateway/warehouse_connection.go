package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type WarehouseConnectionDataGateway interface {
	GetWarehouseConnectByApiKey(ctx context.Context, apiKey string) (*entity.WarehouseConnection, error)
}
