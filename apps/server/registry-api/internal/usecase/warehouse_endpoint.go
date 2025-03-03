package usecase

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
)

func (uc *Usecase) GetWarehouseEndpoint(ctx context.Context, warehouseId string) (*entity.WarehouseEndpoint, error) {
	return uc.WarehouseEndpointDg.GetEndpointByWarehouseId(ctx, warehouseId)
}

func (uc *Usecase) ListWarehouseEndpoints(ctx context.Context) ([]*entity.WarehouseEndpoint, error) {
	return uc.WarehouseEndpointDg.ListEndpoints(ctx)
}

func (uc *Usecase) UpdateWarehouseEndpoint(ctx context.Context, warehouseId string, endpoint string) error {
	return uc.WarehouseEndpointDg.UpdateEndpoint(ctx, warehouseId, endpoint)
}

func (uc *Usecase) DeleteWarehouseEndpoint(ctx context.Context, warehouseId string) error {
	return uc.WarehouseEndpointDg.DeleteEndpoint(ctx, warehouseId)
}
