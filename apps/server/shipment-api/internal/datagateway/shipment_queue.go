package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type ShipmentQueueDataGateway interface {
	CreateToReceivedShipment(ctx context.Context, shipment *entity.Shipment, warehouseId string) error

	WatchReceivedShipment(ctx context.Context, shipmentChan chan<- entity.ShipmentQueue, errorChan chan error, terminateChan <-chan struct{}) error

	ListOutboundShipments(ctx context.Context) (map[string][]entity.ShipmentQueue, error)
	ListInboundShipments(ctx context.Context) (map[string][]entity.ShipmentQueue, error)
}
