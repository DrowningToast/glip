package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
)

type WarehouseEndpointDataGateway interface {
	GetEndpointByWarehouseId(ctx context.Context, warehouseId string) (*entity.WarehouseEndpoint, error)
	ListEndpoints(ctx context.Context) ([]*entity.WarehouseEndpoint, error)
	UpdateEndpoint(ctx context.Context, warehouseId string, endpoint string) error
	DeleteEndpoint(ctx context.Context, warehouseId string) error
}
