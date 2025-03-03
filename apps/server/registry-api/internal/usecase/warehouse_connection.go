package usecase

import (
	"context"
	"encoding/json"

	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
	"github.com/pingcap/errors"
)

func (uc *Usecase) CreateWarehouseConnection(ctx context.Context, warehouseConn entity.WarehouseConnection) (*entity.WarehouseConnection, error) {
	warehouseId := warehouseConn.WarehouseId

	// find if the warehouse id exists in the json config file or not
	var regions map[string]map[string]entity.Warehouse
	bytes, err := json.Marshal(uc.WarehouseRegions)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to marshal warehouse regions")
	}

	if err := json.Unmarshal(bytes, &regions); err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to unmarshal warehouse regions")
	}

	for _, region := range regions {
		for _, w := range region {
			if w.Id == warehouseId {
				return uc.WarehouseConnectionDg.CreateWarehouseConnection(ctx, &warehouseConn)
			}
		}
	}

	return nil, errors.Wrap(errs.ErrInvalidArgument, "warehouse id not found")
}

type GetWarehouseConnectionQuery struct {
	Id     *int
	ApiKey *string
}

// query by warehouse connection id or api key.
// If both field are provided, the id will be used
func (uc *Usecase) GetWarehouseConnection(ctx context.Context, query GetWarehouseConnectionQuery) (*entity.WarehouseConnection, error) {
	if query.Id != nil {
		warehouseConn, err := uc.WarehouseConnectionDg.GetWarehouseConnectionById(ctx, *query.Id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get warehouse connection by id")
		}
		return warehouseConn, nil
	}

	if query.ApiKey != nil {
		warehouseConn, err := uc.WarehouseConnectionDg.GetWarehouseConnectionByApiKey(ctx, *query.ApiKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get warehouse connection by api key")
		}
		return warehouseConn, nil
	}

	return nil, errors.Wrap(errs.ErrInvalidArgument, "id or api key is required")
}

type ListWarehouseConnectionsQuery struct {
	Status *entity.WarehouseConnectionStatus

	Offset int
	Limit  int
}

func (uc *Usecase) ListWarehouseConnections(ctx context.Context, query ListWarehouseConnectionsQuery) ([]*entity.WarehouseConnection, error) {
	if query.Status != nil {
		return uc.WarehouseConnectionDg.ListWarehouseConnectionsByStatus(ctx, *query.Status, query.Limit, query.Offset)
	}

	return uc.WarehouseConnectionDg.ListWarehouseConnections(ctx, query.Limit, query.Offset)
}

func (uc *Usecase) UpdateWarehouseConnection(ctx context.Context, warehouseConn entity.WarehouseConnection) (*entity.WarehouseConnection, error) {
	return uc.WarehouseConnectionDg.UpdateWarehouseConnection(ctx, &warehouseConn)
}

func (uc *Usecase) RevokeWarehouseConnection(ctx context.Context, id int) error {
	return uc.WarehouseConnectionDg.RevokeWarehouseConnection(ctx, id)
}
