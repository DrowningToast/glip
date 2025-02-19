package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
)

var _ datagateway.CarrierDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateCarrier(ctx context.Context, carrierPtr *entity.Carrier) (*entity.Carrier, error) {
	if carrierPtr == nil {
		return nil, errors.Wrap(errs.ErrInternal, "carrier is nil")
	}

	carrier, err := r.queries.CreateCarrier(ctx, shipment_database.CreateCarrierParams{
		Name:          carrierPtr.Name,
		ContactPerson: mapStringPtrToPgText(carrierPtr.ContactPerson),
		ContactPhone:  mapStringPtrToPgText(carrierPtr.ContactPhone),
		Email:         mapStringPtrToPgText(carrierPtr.Email),
		Description:   mapStringPtrToPgText(carrierPtr.Description),
		Status:        string(carrierPtr.Status),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create carrier")
	}

	return mapCarrierModelToEntity(&carrier), nil
}

func (r *PostgresRepository) GetCarrierById(ctx context.Context, id int) (*entity.Carrier, error) {
	carrier, err := r.queries.GetCarrierById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get carrier by id")
	}

	return mapCarrierModelToEntity(&carrier), nil
}

func (r *PostgresRepository) ListCarriers(ctx context.Context, limit int, offset int) ([]*entity.Carrier, error) {
	carriers, err := r.queries.ListCarriers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list carriers")
	}

	return lo.Map(carriers, func(carrier shipment_database.Carrier, _ int) *entity.Carrier {
		return mapCarrierModelToEntity(&carrier)
	}), nil
}

func (r *PostgresRepository) UpdateCarrier(ctx context.Context, carrierPtr *entity.Carrier) (*entity.Carrier, error) {
	carrier, err := r.queries.UpdateCarrier(ctx, shipment_database.UpdateCarrierParams{
		ID: int32(carrierPtr.Id),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update carrier")
	}

	return mapCarrierModelToEntity(&carrier), nil
}

func (r *PostgresRepository) GetCarrierShipmentStats(ctx context.Context, id int) (*datagateway.CarrierShipmentStats, error) {
	stats, err := r.queries.GetCarrierShipmentStats(ctx, int32(id))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get carrier shipment stats")
	}

	carrier := mapCarrierModelToEntity(&shipment_database.Carrier{
		ID:            int32(id),
		Name:          stats.Name,
		Status:        stats.Status,
		ContactPerson: stats.ContactPerson,
		ContactPhone:  stats.ContactPhone,
		Email:         stats.Email,
		Description:   stats.Description,
	})

	return &datagateway.CarrierShipmentStats{
		Carrier: carrier,

		TotalShipments:     int(stats.TotalShipments),
		DeliveredShipments: int(stats.DeliveredShipments),
		CanceledShipments:  int(stats.CancelledShipments),
		AvgDelayHours:      float64(stats.AvgDelayHours),
	}, nil
}
