package postgres

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

var _ datagateway.WarehouseDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateWarehouse(ctx context.Context, warehousePtr *entity.Warehouse) (*entity.Warehouse, error) {
	if warehousePtr == nil {
		return nil, errors.Wrap(errs.ErrInternal, "warehouse is empty")
	}

	warehouse, err := r.queries.CreateWarehouse(ctx, shipment_database.CreateWarehouseParams{
		Name:            warehousePtr.Name,
		LocationAddress: warehousePtr.Location,
		Country:         warehousePtr.Country,
		City:            warehousePtr.City,
		TotalCapacity:   mapDecimalToPgNumeric(warehousePtr.TotalCapacity),
		CurrentCapacity: mapDecimalToPgNumeric(warehousePtr.CurrentCapacity),
		Description:     mapStringPtrToPgText(warehousePtr.Description),
		Status:          pgtype.Text{String: string(warehousePtr.Status), Valid: true},
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, "warehouse already exists")
		}
		return nil, errors.Wrap(errs.ErrInternal, "failed to create warehouse")
	}

	return mapWarehouseModelToEntity(&warehouse), nil
}

func (r *PostgresRepository) GetWarehouse(ctx context.Context, id int) (*entity.Warehouse, error) {
	warehouse, err := r.queries.GetWarehouse(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse")
	}

	return mapWarehouseModelToEntity(&warehouse), nil
}

func (r *PostgresRepository) GetWarehouseByStatus(ctx context.Context, status entity.WarehouseStatus) ([]*entity.Warehouse, error) {
	warehouses, err := r.queries.GetWarehouseByStatus(ctx, pgtype.Text{String: string(status), Valid: true})
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse by status")
	}

	return lo.Map(warehouses, func(w shipment_database.Warehouse, _ int) *entity.Warehouse {
		return mapWarehouseModelToEntity(&w)
	}), nil
}

func (r *PostgresRepository) GetWarehouseByCountry(ctx context.Context, country string) ([]*entity.Warehouse, error) {
	warehouses, err := r.queries.GetWarehousesByCountry(ctx, country)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse by country")
	}

	return lo.Map(warehouses, func(w shipment_database.Warehouse, _ int) *entity.Warehouse {
		return mapWarehouseModelToEntity(&w)
	}), nil
}

func (r *PostgresRepository) GetWarehouseByCity(ctx context.Context, city string) ([]*entity.Warehouse, error) {
	warehouses, err := r.queries.GetWarehousesByCity(ctx, city)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse by city")
	}

	return lo.Map(warehouses, func(w shipment_database.Warehouse, _ int) *entity.Warehouse {
		return mapWarehouseModelToEntity(&w)
	}), nil
}

func (r *PostgresRepository) ListWarehouses(ctx context.Context) ([]*entity.Warehouse, error) {
	warehouses, err := r.queries.ListWarehouses(ctx)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to list warehouses")
	}

	return lo.Map(warehouses, func(w shipment_database.Warehouse, _ int) *entity.Warehouse {
		return mapWarehouseModelToEntity(&w)
	}), nil
}

func (r *PostgresRepository) UpdateWarehouse(ctx context.Context, warehousePtr *entity.Warehouse) (*entity.Warehouse, error) {
	warehouse, err := r.queries.UpdateWarehouse(ctx, shipment_database.UpdateWarehouseParams{
		ID:              int32(warehousePtr.Id),
		Name:            warehousePtr.Name,
		LocationAddress: warehousePtr.Location,
		Country:         warehousePtr.Country,
		City:            warehousePtr.City,
		TotalCapacity:   mapDecimalToPgNumeric(warehousePtr.TotalCapacity),
		CurrentCapacity: mapDecimalToPgNumeric(warehousePtr.CurrentCapacity),
		Description:     mapStringPtrToPgText(warehousePtr.Description),
		Status:          pgtype.Text{String: string(warehousePtr.Status), Valid: true},
	})
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to update warehouse")
	}

	return mapWarehouseModelToEntity(&warehouse), nil
}

func (r *PostgresRepository) UpdateWarehouseCapacity(ctx context.Context, warehousePtr *entity.Warehouse) (*entity.Warehouse, error) {
	warehouse, err := r.queries.UpdateWarehouseCapacity(ctx, shipment_database.UpdateWarehouseCapacityParams{
		ID:              int32(warehousePtr.Id),
		CurrentCapacity: mapDecimalToPgNumeric(warehousePtr.CurrentCapacity),
	})
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to update warehouse capacity")
	}

	return mapWarehouseModelToEntity(&warehouse), nil
}

func (r *PostgresRepository) DeleteWarehouse(ctx context.Context, id int) error {
	err := r.queries.DeleteWarehouse(ctx, int32(id))
	if err != nil {
		return errors.Wrap(errs.ErrInternal, "failed to delete warehouse")
	}

	return nil
}
