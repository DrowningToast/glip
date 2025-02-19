package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type WarehouseConnectionDataGateway interface {
	CreateWarehouseConnection(ctx context.Context, connection *entity.WarehouseConnection) (*entity.WarehouseConnection, error)
	GetWarehouseConnectionById(ctx context.Context, id int) (*entity.WarehouseConnection, error)
	GetWarehouseConnectionByApiKey(ctx context.Context, apiKey string) (*entity.WarehouseConnection, error)
	ListWarehouseConnections(ctx context.Context, limit int, offset int) ([]*entity.WarehouseConnection, error)
	ListWarehouseConnectionsByStatus(ctx context.Context, status entity.WarehouseConnectionStatus, limit int, offset int) ([]*entity.WarehouseConnection, error)

	UpdateWarehouseConnection(ctx context.Context, connection *entity.WarehouseConnection) (*entity.WarehouseConnection, error)
	RevokeWarehouseConnection(ctx context.Context, id int) error
}
