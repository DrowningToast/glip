package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
)

var _ datagateway.WarehouseConnectionDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateWarehouseConnection(ctx context.Context, newConnection *entity.WarehouseConnection) (*entity.WarehouseConnection, error) {
	connection, err := r.queries.CreateWarehouseConnection(ctx, shipment_database.CreateWarehouseConnectionParams{
		WarehouseID: int32(newConnection.WarehouseId),
		ApiKey:      newConnection.ApiKey,
		Name:        newConnection.Name,
		Status:      string(newConnection.Status),
		CreatedBy:   mapIntPtrToPgInt4(&newConnection.CreatedBy),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, "warehouse connection already exists")
		}
		return nil, errors.Wrap(errs.ErrInternal, "failed to create warehouse connection")
	}

	return mapWarehouseConnectionModelToEntity(&connection), nil
}

func (r *PostgresRepository) GetWarehouseConnectionById(ctx context.Context, id int) (*entity.WarehouseConnection, error) {
	connection, err := r.queries.GetWarehouseConnectionById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse connection by id")
	}

	return mapWarehouseConnectionModelToEntity(&connection), nil
}

func (r *PostgresRepository) GetWarehouseConnectionByApiKey(ctx context.Context, apiKey string) (*entity.WarehouseConnection, error) {
	connection, err := r.queries.GetWarehouseConnectionByApiKey(ctx, apiKey)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse connection by api key")
	}
	return mapWarehouseConnectionModelToEntity(&connection), nil
}

func (r *PostgresRepository) ListWarehouseConnections(ctx context.Context, limit int, offset int) ([]*entity.WarehouseConnection, error) {
	connections, err := r.queries.ListWarehouseConnections(ctx, shipment_database.ListWarehouseConnectionsParams{
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to list warehouse connections")
	}
	return lo.Map(connections, func(connection shipment_database.WarehouseConnection, _ int) *entity.WarehouseConnection {
		return mapWarehouseConnectionModelToEntity(&connection)
	}), nil
}

func (r *PostgresRepository) ListWarehouseConnectionsByStatus(ctx context.Context, status entity.WarehouseConnectionStatus, limit int, offset int) ([]*entity.WarehouseConnection, error) {
	connections, err := r.queries.ListWarehouseConnectionsByStatus(ctx, shipment_database.ListWarehouseConnectionsByStatusParams{
		Status:       string(status),
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to list warehouse connections by status")
	}
	return lo.Map(connections, func(connection shipment_database.WarehouseConnection, _ int) *entity.WarehouseConnection {
		return mapWarehouseConnectionModelToEntity(&connection)
	}), nil
}

func (r *PostgresRepository) UpdateWarehouseConnection(ctx context.Context, newConnection *entity.WarehouseConnection) (*entity.WarehouseConnection, error) {
	connection, err := r.queries.UpdateWarehouseConnection(ctx, shipment_database.UpdateWarehouseConnectionParams{
		ID:          int32(newConnection.Id),
		Name:        newConnection.Name,
		Status:      string(newConnection.Status),
		LastUsedAt:  mapTimePtrToTimestamp(newConnection.LastUsedAt),
		ApiKey:      newConnection.ApiKey,
		WarehouseID: int32(newConnection.WarehouseId),
		CreatedBy:   mapIntPtrToPgInt4(&newConnection.CreatedBy),
	})
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to update warehouse connection")
	}
	return mapWarehouseConnectionModelToEntity(&connection), nil
}

func (r *PostgresRepository) RevokeWarehouseConnection(ctx context.Context, id int) error {
	_, err := r.queries.RevokeWarehouseConnection(ctx, int32(id))
	if err != nil {
		return errors.Wrap(errs.ErrInternal, "failed to revoke warehouse connection")
	}

	return nil
}
