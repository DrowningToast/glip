package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type WarehouseDataGateway interface {
	GetWarehouse(ctx context.Context, id int) (*entity.Warehouse, error)
	GetWarehouseByStatus(ctx context.Context, status entity.WarehouseStatus) ([]*entity.Warehouse, error)
	GetWarehouseByCountry(ctx context.Context, country string) ([]*entity.Warehouse, error)
	GetWarehouseByCity(ctx context.Context, city string) ([]*entity.Warehouse, error)

	ListWarehousesByRegion(ctx context.Context, region string) ([]*entity.Warehouse, error)
	ListWarehouses(ctx context.Context) ([]*entity.Warehouse, error)

	GetWarehouseEndpoint(ctx context.Context, id int) (*string, error)
}
