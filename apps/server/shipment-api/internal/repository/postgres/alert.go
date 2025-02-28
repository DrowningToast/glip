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

var _ datagateway.AlertDataGateway = &PostgresRepository{}

func (r *PostgresRepository) CreateShipmentAlert(ctx context.Context, newAlert *entity.Alert) (*entity.Alert, error) {
	alert, err := r.queries.CreateShipmentAlert(ctx, shipment_database.CreateShipmentAlertParams{
		AlertType:         string(newAlert.Type),
		Severity:          string(newAlert.Severity),
		Status:            string(newAlert.Status),
		RelatedEntityType: string(newAlert.RelatedEntityType),
		RelatedEntityID:   int32(newAlert.RelatedEntityId),
		Description:       mapStringPtrToPgText(newAlert.Description),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, err.Error())
		}

		return nil, err
	}

	return mapAlertModelToEntity(&alert), nil
}

func (r *PostgresRepository) GetShipmentAlertById(ctx context.Context, id int) (*entity.Alert, error) {
	alert, err := r.queries.GetShipmentAlertById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errs.ErrNotFound, err.Error())
		}
		return nil, err
	}

	return mapAlertModelToEntity(&alert), nil
}

func (r *PostgresRepository) ListShipmentActiveAlerts(ctx context.Context, limit int, offset int) ([]*entity.Alert, error) {
	alerts, err := r.queries.ListShipmentActiveAlerts(ctx, shipment_database.ListShipmentActiveAlertsParams{
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(alerts, func(alert shipment_database.Alert, _ int) *entity.Alert {
		return mapAlertModelToEntity(&alert)
	}), nil
}

func (r *PostgresRepository) ListShipmentAlertsByType(ctx context.Context, alertType entity.AlertType, limit int, offset int) ([]*entity.Alert, error) {
	alerts, err := r.queries.ListShipmentAlertsByType(ctx, shipment_database.ListShipmentAlertsByTypeParams{
		AlertType:    string(alertType),
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(alerts, func(alert shipment_database.Alert, _ int) *entity.Alert {
		return mapAlertModelToEntity(&alert)
	}), nil
}

func (r *PostgresRepository) ListShipmentAlertsBySeverity(ctx context.Context, severity entity.AlertSeverity, limit int, offset int) ([]*entity.Alert, error) {
	alerts, err := r.queries.ListShipmentAlertsBySeverity(ctx, shipment_database.ListShipmentAlertsBySeverityParams{
		Severity:     string(severity),
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(alerts, func(alert shipment_database.Alert, _ int) *entity.Alert {
		return mapAlertModelToEntity(&alert)
	}), nil
}

func (r *PostgresRepository) ListShipmentAlertsByEntityType(ctx context.Context, entityType entity.AlertRelatedEntityType, limit int, offset int) ([]*entity.Alert, error) {
	alerts, err := r.queries.ListShipmentAlertsByEntityType(ctx, shipment_database.ListShipmentAlertsByEntityTypeParams{
		EntityType:   string(entityType),
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(alerts, func(alert shipment_database.Alert, _ int) *entity.Alert {
		return mapAlertModelToEntity(&alert)
	}), nil
}

func (r *PostgresRepository) ListShipmentAlertsByEntityId(ctx context.Context, entityType entity.AlertRelatedEntityType, entityID int, limit int, offset int) ([]*entity.Alert, error) {
	alerts, err := r.queries.ListShipmentAlertsByEntityId(ctx, shipment_database.ListShipmentAlertsByEntityIdParams{
		EntityType:   string(entityType),
		EntityID:     int32(entityID),
		ReturnLimit:  mapIntPtrToPgInt4(&limit),
		ReturnOffset: mapIntPtrToPgInt4(&offset),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(alerts, func(alert shipment_database.Alert, _ int) *entity.Alert {
		return mapAlertModelToEntity(&alert)
	}), nil
}

func (r *PostgresRepository) UpdateShipmentAlertStatus(ctx context.Context, id int, status entity.AlertStatus) (*entity.Alert, error) {
	alert, err := r.queries.UpdateShipmentAlertStatus(ctx, shipment_database.UpdateShipmentAlertStatusParams{
		ID:     int32(id),
		Status: string(status),
	})
	if err != nil {
		return nil, err
	}

	return mapAlertModelToEntity(&alert), nil
}
