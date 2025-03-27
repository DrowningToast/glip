package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/internal/utils/pgmapper"
	database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

var _ datagateway.ShipmentDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateShipment(ctx context.Context, s *entity.Shipment) (*entity.Shipment, error) {
	if s == nil {
		return nil, errors.Wrap(errs.ErrInternal, "shipment is nil")
	}

	log.Debug(s)

	shipment, err := r.queries.CreateShipment(ctx, database.CreateShipmentParams{
		Route:                  lo.Map(s.Route, func(id string, _ int) string { return id }),
		LastWarehouseID:        pgmapper.MapStringPtrToPgText(s.LastWarehouseId),
		DepartureWarehouseID:   s.DepartureWarehouseId,
		DepartureAddress:       pgmapper.MapStringPtrToPgText(s.DepartureAddress),
		DestinationWarehouseID: s.DestinationWarehouseId,
		DestinationAddress:     s.DestinationAddress,
		// CarrierID:           pgmapper.MapIntPtrToPgInt4(s.CarrierId),
		Status:              string(s.Status),
		TotalWeight:         pgmapper.MapDecimalToPgNumeric(s.TotalWeight),
		TotalVolume:         pgmapper.MapDecimalToPgNumeric(s.TotalVolume),
		SpecialInstructions: pgmapper.MapStringPtrToPgText(s.SpecialInstructions),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, err.Error())
		}
		return nil, err
	}

	log.Debug(shipment)

	return mapShipmentModelToEntity(&shipment), nil
}

func (r *PostgresRepository) CreateShipmentWithOwner(ctx context.Context, s *entity.Shipment) (*entity.Shipment, error) {
	if s == nil {
		return nil, errors.Wrap(errs.ErrInternal, "shipment is nil")
	}

	log.Debug(s)

	shipment, err := r.queries.CreateShipmentWithOwner(ctx, database.CreateShipmentWithOwnerParams{
		Route:                  lo.Map(s.Route, func(id string, _ int) string { return id }),
		LastWarehouseID:        pgmapper.MapStringPtrToPgText(s.LastWarehouseId),
		DepartureWarehouseID:   s.DepartureWarehouseId,
		DepartureAddress:       pgmapper.MapStringPtrToPgText(s.DepartureAddress),
		DestinationWarehouseID: s.DestinationWarehouseId,
		DestinationAddress:     s.DestinationAddress,
		// CarrierID:           pgmapper.MapIntPtrToPgInt4(s.CarrierId),
		Status:              string(s.Status),
		TotalWeight:         pgmapper.MapDecimalToPgNumeric(s.TotalWeight),
		TotalVolume:         pgmapper.MapDecimalToPgNumeric(s.TotalVolume),
		SpecialInstructions: pgmapper.MapStringPtrToPgText(s.SpecialInstructions),
		OwnerID:             pgmapper.MapIntPtrToPgInt4(s.OwnerId),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, err.Error())
		}
		return nil, err
	}

	log.Debug(shipment.OwnerID)

	return mapShipmentModelToEntity(&shipment), nil
}

func (r *PostgresRepository) GetShipmentById(ctx context.Context, id int) (*entity.Shipment, error) {
	shipment, err := r.queries.GetShipmentById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errs.ErrNotFound, err.Error())
		}
		return nil, err
	}
	return mapShipmentModelToEntity(&shipment), nil
}

func (r *PostgresRepository) ListShipments(ctx context.Context, limit int, offset int) ([]*entity.Shipment, error) {
	shipments, err := r.queries.ListShipments(ctx, database.ListShipmentsParams{
		ReturnOffset: pgtype.Int4{Int32: int32(offset)},
		ReturnLimit:  pgtype.Int4{Int32: int32(limit)},
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(shipments, func(s database.Shipment, _ int) *entity.Shipment {
		return mapShipmentModelToEntity(&s)
	}), nil
}

func (r *PostgresRepository) ListShipmentsByLastWarehouse(ctx context.Context, lastWarehouseId string, limit int, offset int) ([]*entity.Shipment, error) {
	shipments, err := r.queries.ListShipmentsByLastWarehouse(ctx, database.ListShipmentsByLastWarehouseParams{
		WarehouseID:  pgtype.Text{String: lastWarehouseId},
		ReturnLimit:  pgtype.Int4{Int32: int32(limit)},
		ReturnOffset: pgtype.Int4{Int32: int32(offset)},
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(shipments, func(s database.Shipment, _ int) *entity.Shipment {
		return mapShipmentModelToEntity(&s)
	}), nil
}

func (r *PostgresRepository) ListShipmentsByStatus(ctx context.Context, status entity.ShipmentStatus, limit int, offset int) ([]*entity.Shipment, error) {
	shipments, err := r.queries.ListShipmentsByStatus(ctx, database.ListShipmentsByStatusParams{
		Status:       string(status),
		ReturnLimit:  pgtype.Int4{Int32: int32(limit)},
		ReturnOffset: pgtype.Int4{Int32: int32(offset)},
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(shipments, func(s database.Shipment, _ int) *entity.Shipment {
		return mapShipmentModelToEntity(&s)
	}), nil
}

func (r *PostgresRepository) ListShipmentsByStatusAndLastWarehouse(ctx context.Context, status entity.ShipmentStatus, lastWarehouseId string, limit int, offset int) ([]*entity.Shipment, error) {
	shipments, err := r.queries.ListShipmentsByStatusAndLastWarehouse(ctx, database.ListShipmentsByStatusAndLastWarehouseParams{
		Status:       string(status),
		WarehouseID:  pgtype.Text{String: lastWarehouseId},
		ReturnLimit:  pgtype.Int4{Int32: int32(limit)},
		ReturnOffset: pgtype.Int4{Int32: int32(offset)},
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(shipments, func(s database.Shipment, _ int) *entity.Shipment {
		return mapShipmentModelToEntity(&s)
	}), nil
}

func (r *PostgresRepository) ListShipmentsByAccountUsername(ctx context.Context, username string, limit int, offset int, status *entity.ShipmentStatus) ([]*entity.Shipment, error) {
	shipments, err := r.queries.ListShipmentsByAccountUsername(ctx, database.ListShipmentsByAccountUsernameParams{
		Username:     username,
		ReturnLimit:  pgtype.Int4{Int32: int32(limit)},
		ReturnOffset: pgtype.Int4{Int32: int32(offset)},
		Status:       pgmapper.MapStringPtrToPgText((*string)(status)),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(shipments, func(s database.ListShipmentsByAccountUsernameRow, _ int) *entity.Shipment {
		return mapShipmentJoinedAccountModelToEntity(s)
	}), nil
}

func (r *PostgresRepository) UpdateShipment(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error) {
	if shipment == nil {
		return nil, errors.Wrap(errs.ErrInternal, "shipment is nil")
	}

	updatedShipment, err := r.queries.UpdateShipment(ctx, database.UpdateShipmentParams{
		ID:                 int32(shipment.Id),
		Route:              lo.Map(shipment.Route, func(id string, _ int) string { return id }),
		LastWarehouseID:    pgmapper.MapStringPtrToPgText(shipment.LastWarehouseId),
		DestinationAddress: shipment.DestinationAddress,
		// CarrierID:          pgmapper.MapIntPtrToPgInt4(shipment.CarrierId),
		Status:      string(shipment.Status),
		TotalWeight: pgmapper.MapDecimalToPgNumeric(shipment.TotalWeight),
		TotalVolume: pgmapper.MapDecimalToPgNumeric(shipment.TotalVolume),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errs.ErrNotFound, err.Error())
		}
		return nil, err
	}

	return mapShipmentModelToEntity(&updatedShipment), nil
}

func (r *PostgresRepository) UpdateShipmentStatus(ctx context.Context, id int, status entity.ShipmentStatus) (*entity.Shipment, error) {
	updatedShipment, err := r.queries.UpdateShipmentStatus(ctx, database.UpdateShipmentStatusParams{
		ID:     int32(id),
		Status: string(status),
	})
	if err != nil {
		return nil, err
	}

	return mapShipmentModelToEntity(&updatedShipment), nil
}
