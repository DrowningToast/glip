package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type AlertDataGateway interface {
	CreateShipmentAlert(ctx context.Context, alert *entity.Alert) (*entity.Alert, error)
	GetShipmentAlertById(ctx context.Context, id int) (*entity.Alert, error)
	ListShipmentActiveAlerts(ctx context.Context, limit int, offset int) ([]*entity.Alert, error)
	ListShipmentAlertsByType(ctx context.Context, alertType entity.AlertType, limit int, offset int) ([]*entity.Alert, error)
	ListShipmentAlertsBySeverity(ctx context.Context, severity entity.AlertSeverity, limit int, offset int) ([]*entity.Alert, error)
	ListShipmentAlertsByEntityType(ctx context.Context, entityType entity.AlertRelatedEntityType, limit int, offset int) ([]*entity.Alert, error)
	ListShipmentAlertsByEntityId(ctx context.Context, entityType entity.AlertRelatedEntityType, entityID int, limit int, offset int) ([]*entity.Alert, error)

	UpdateShipmentAlertStatus(ctx context.Context, id int, status entity.AlertStatus) (*entity.Alert, error)
}
