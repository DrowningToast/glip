package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

func (uc *Usecase) FindShipmentRoute(ctx context.Context, departureWarehouseId string, destinationWarehouseId string) ([]string, error) {
	departureRegion, err := uc.Config.WarehouseRegions.GetWarehouseRegion(departureWarehouseId)
	if err != nil {
		return nil, err
	}

	route, ok := uc.Config.WarehouseRoutes[string(*departureRegion)][departureWarehouseId][destinationWarehouseId]
	if !ok {
		return nil, errors.Wrap(errs.ErrNotFound, "route not found")
	}

	return route.Route, nil
}

func (uc *Usecase) CreateShipment(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error) {
	if shipment == nil {
		return nil, errors.Wrap(errs.ErrInvalidArgument, "shipment is nil")
	}

	route, err := uc.FindShipmentRoute(ctx, shipment.DepartureWarehouseId, shipment.DestinationWarehouseId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find shipment route")
	}

	createdShipment, err := uc.ShipmentDg.CreateShipment(ctx, &entity.Shipment{
		// Id is set to 0, because the repository will set it
		DestinationAddress:  shipment.DestinationAddress,
		CarrierId:           shipment.CarrierId,
		Status:              entity.ShipmentStatusWaitingForPickup,
		TotalWeight:         shipment.TotalWeight,
		TotalVolume:         shipment.TotalVolume,
		SpecialInstructions: shipment.SpecialInstructions,
		Route:               route,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create shipment")
	}

	// QUEUE: Create expected shipment arrives at warehouse event

	return createdShipment, nil
}
