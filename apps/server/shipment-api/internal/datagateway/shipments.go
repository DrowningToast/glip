package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type ShipmentDataGateway interface {
	CreateShipment(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error)
	CreateShipmentWithOwner(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error)

	GetShipmentById(ctx context.Context, id int) (*entity.Shipment, error)
	ListShipments(ctx context.Context, limit int, offset int) ([]*entity.Shipment, error)

	// Get many shipments by last warehouse
	ListShipmentsByLastWarehouse(ctx context.Context, lastWarehouseId string, limit int, offset int) ([]*entity.Shipment, error)
	ListShipmentsByStatus(ctx context.Context, status entity.ShipmentStatus, limit int, offset int) ([]*entity.Shipment, error)
	ListShipmentsByStatusAndLastWarehouse(ctx context.Context, status entity.ShipmentStatus, lastWarehouseId string, limit int, offset int) ([]*entity.Shipment, error)
	ListShipmentsByAccountUsername(ctx context.Context, username string, limit int, offset int, status *entity.ShipmentStatus) ([]*entity.Shipment, error)

	UpdateShipment(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error)
	UpdateShipmentStatus(ctx context.Context, id int, status entity.ShipmentStatus) (*entity.Shipment, error)
}
