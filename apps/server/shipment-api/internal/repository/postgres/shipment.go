package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
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

	shipment, err := r.queries.CreateShipment(ctx, database.CreateShipmentParams{
		Route:               lo.Map(s.Route, func(id int, _ int) int32 { return int32(id) }),
		LastWarehouseID:     mapIntPtrToPgInt4(s.LastWarehouseId),
		DestinationAddress:  s.DestinationAddress,
		CarrierID:           mapIntPtrToPgInt4(s.CarrierId),
		ScheduledDeparture:  mapTimeToTimestamp(s.ScheduledDeparture),
		ScheduledArrival:    mapTimeToTimestamp(s.ScheduledArrival),
		ActualDeparture:     mapTimePtrToTimestamp(s.ActualDeparture),
		ActualArrival:       mapTimePtrToTimestamp(s.ActualArrival),
		Status:              string(s.Status),
		TotalWeight:         mapDecimalToPgNumeric(s.TotalWeight),
		TotalVolume:         mapDecimalToPgNumeric(s.TotalVolume),
		SpecialInstructions: mapStringPtrToPgText(s.SpecialInstructions),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, err.Error())
		}
		return nil, err
	}

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

func (r *PostgresRepository) ListShipmentsByLastWarehouse(ctx context.Context, lastWarehouseId int, limit int, offset int) ([]*entity.Shipment, error) {
	shipments, err := r.queries.ListShipmentsByLastWarehouse(ctx, database.ListShipmentsByLastWarehouseParams{
		WarehouseID:  pgtype.Int4{Int32: int32(lastWarehouseId)},
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

func (r *PostgresRepository) UpdateShipment(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error) {
	if shipment == nil {
		return nil, errors.Wrap(errs.ErrInternal, "shipment is nil")
	}

	updatedShipment, err := r.queries.UpdateShipment(ctx, database.UpdateShipmentParams{
		ID:                 int32(shipment.Id),
		Route:              lo.Map(shipment.Route, func(id int, _ int) int32 { return int32(id) }),
		LastWarehouseID:    mapIntPtrToPgInt4(shipment.LastWarehouseId),
		DestinationAddress: shipment.DestinationAddress,
		CarrierID:          mapIntPtrToPgInt4(shipment.CarrierId),
		ScheduledDeparture: mapTimeToTimestamp(shipment.ScheduledDeparture),
		ScheduledArrival:   mapTimeToTimestamp(shipment.ScheduledArrival),
		ActualDeparture:    mapTimePtrToTimestamp(shipment.ActualDeparture),
		ActualArrival:      mapTimePtrToTimestamp(shipment.ActualArrival),
		Status:             string(shipment.Status),
		TotalWeight:        mapDecimalToPgNumeric(shipment.TotalWeight),
		TotalVolume:        mapDecimalToPgNumeric(shipment.TotalVolume),
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
