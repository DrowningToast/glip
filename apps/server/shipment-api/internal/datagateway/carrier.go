package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type CarrierShipmentStats struct {
	Carrier *entity.Carrier

	TotalShipments     int
	DeliveredShipments int
	CanceledShipments  int
	AvgDelayHours      float64
}

type CarrierDataGateway interface {
	CreateCarrier(ctx context.Context, carrier *entity.Carrier) (*entity.Carrier, error)
	GetCarrierById(ctx context.Context, id int) (*entity.Carrier, error)
	ListCarriers(ctx context.Context, limit int, offset int) ([]*entity.Carrier, error)
	UpdateCarrier(ctx context.Context, carrier *entity.Carrier) (*entity.Carrier, error)

	GetCarrierShipmentStats(ctx context.Context, id int) (*CarrierShipmentStats, error)
}
