package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func (uc *Usecase) FindShipmentRoute(ctx context.Context, departureWarehouseId string, destinationWarehouseId string) ([]string, error) {
	departureRegion, err := uc.Config.WarehouseRegions.GetWarehouseRegion(departureWarehouseId)
	if err != nil {
		return nil, err
	}

	route, ok := uc.Config.WarehouseRoutes[string(departureRegion)][departureWarehouseId][destinationWarehouseId]
	if !ok {
		return nil, errors.Wrap(errs.ErrNotFound, "route not found")
	}

	return route.Route, nil
}

type CreateShipmentParams struct {
	DepartureAddress   string `json:"departure_address" validate:"required"`
	DepartureCity      string `json:"departure_city" validate:"required"`
	DestinationAddress string `json:"destination_address" validate:"required"`
	DestinationCity    string `json:"destination_city" validate:"required"`
	// CarrierId           int             `json:"carrier_id"`
	TotalWeight         decimal.Decimal `json:"total_weight" validate:"required"`
	TotalVolume         decimal.Decimal `json:"total_volume" validate:"required"`
	SpecialInstructions *string         `json:"special_instructions,omitempty"`
}

func (uc *Usecase) CreateShipment(ctx context.Context, shipment CreateShipmentParams) (*entity.Shipment, error) {
	departureWarehouse, err := uc.Config.WarehouseRegions.GetWarehouseByCity(shipment.DepartureCity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get departure warehouse")
	}

	destinationWarehouse, err := uc.Config.WarehouseRegions.GetWarehouseByCity(shipment.DestinationCity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get destination warehouse")
	}

	route, err := uc.FindShipmentRoute(ctx, departureWarehouse.Id, destinationWarehouse.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find shipment route")
	}

	createdShipment, err := uc.ShipmentDg.CreateShipment(ctx, &entity.Shipment{
		// Id is set to 0, because the repository will set it
		DepartureAddress:       &shipment.DepartureAddress,
		DepartureWarehouseId:   route[0],
		DestinationAddress:     shipment.DestinationAddress,
		DestinationWarehouseId: route[len(route)-1],
		// CarrierId:              &shipment.CarrierId,
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
	err = uc.ShipmentQueueDg.CreateToReceivedShipment(ctx, createdShipment, route[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to create expected shipment arrives at warehouse event")
	}

	return createdShipment, nil
}

func (uc *Usecase) WatchShipmentUpdates(ctx context.Context, errorChan chan error) {
	shipmentQueueChan := make(chan entity.ShipmentQueue)
	go uc.ShipmentQueueDg.WatchReceivedShipment(ctx, shipmentQueueChan, errorChan, ctx.Done())

	for {
		select {
		// incoming shipment queue
		case shipmentQueue := <-shipmentQueueChan:
			// validate shipment queue
			if shipmentQueue.QueueType == entity.ShipmentQueueTypeOutbound {
				errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment queue type")
				continue
			}

			oldShipment, err := uc.ShipmentDg.GetShipmentById(ctx, shipmentQueue.Id)
			if err != nil {
				log.Warnf("failed to get shipment: %v", err)
				continue
			}
			if oldShipment == nil {
				log.Warnf("shipment not found")
				errorChan <- errors.Wrap(errs.ErrNotFound, "shipment not found")
				continue
			}
			if oldShipment.LastWarehouseId == nil {
				log.Warnf("invalid shipment status")
				errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment status")
				continue
			}

			switch shipmentQueue.Status {
			case entity.ShipmentStatusArrivedAtWarehouse:
				if oldShipment.LastWarehouseId == nil {
					if *shipmentQueue.FromWarehouseId != oldShipment.Route[0] {
						log.Warnf("invalid shipment route")
						errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment route")
						continue
					}

					// First arrival
				} else {
					// Not the first arrival
					_, currentIndex, found := lo.FindIndexOf(oldShipment.Route, func(warehouseId string) bool {
						return warehouseId == *shipmentQueue.FromWarehouseId
					})
					if !found {
						log.Warnf("last warehouse id not found in route")
						errorChan <- errors.Wrap(errs.ErrNotFound, "last warehouse id not found in route")
						continue
					}
					if currentIndex == len(oldShipment.Route)-1 {
						log.Warnf("invalid shipment route")
						errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment route")
						continue
					}
					if oldShipment.Route[currentIndex-1] != *shipmentQueue.LastWarehouseId {
						log.Warnf("invalid shipment route")
						errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment route")
						continue
					}
				}

				oldShipment.LastWarehouseId = shipmentQueue.FromWarehouseId
				oldShipment.Status = entity.ShipmentStatusArrivedAtWarehouse
				oldShipment, err = uc.ShipmentDg.UpdateShipment(ctx, oldShipment)
				if err != nil {
					log.Warnf("failed to update shipment: %v", err)
					errorChan <- errors.Wrap(err, "failed to update shipment")
					continue
				}
			case entity.ShipmentStatusInTransitOnTheWay:
			case entity.ShipmentStatusDelivered:
				if shipmentQueue.LastWarehouseId == nil || shipmentQueue.FromWarehouseId == nil || *shipmentQueue.FromWarehouseId != *shipmentQueue.LastWarehouseId || *shipmentQueue.FromWarehouseId != shipmentQueue.DestinationWarehouseId {
					log.Warnf("invalid shipment status")
					errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment status")
					continue
				}

				oldShipment.Status = shipmentQueue.Status
				oldShipment, err = uc.ShipmentDg.UpdateShipment(ctx, oldShipment)
				if err != nil {
					log.Warnf("failed to update shipment: %v", err)
					errorChan <- errors.Wrap(err, "failed to update shipment")
					continue
				}
				break

			default:
				log.Warnf("invalid shipment status")
				errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment status")
				continue
			}

			oldShipment, err = uc.ShipmentDg.GetShipmentById(ctx, shipmentQueue.Id)
			if err != nil {
				log.Warnf("failed to get shipment: %v", err)
				errorChan <- errors.Wrap(err, "failed to get shipment")
				continue
			}
			if oldShipment == nil {
				log.Warnf("shipment not found")
				errorChan <- errors.Wrap(errs.ErrNotFound, "shipment not found")
				continue
			}

			// update shipment status
			uc.ShipmentDg.UpdateShipment(ctx, &entity.Shipment{
				Id:     shipmentQueue.Id,
				Status: entity.ShipmentStatus(shipmentQueue.Status),
			})
			break
		case err := <-errorChan:
			log.Warnf("error: %v", err)
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (uc *Usecase) GetShipmentById(ctx context.Context, id int) (*entity.Shipment, error) {
	shipment, err := uc.ShipmentDg.GetShipmentById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shipment")
	}
	return shipment, nil
}

type ListShipmentsParams struct {
	Limit           int
	Offset          int
	Status          *entity.ShipmentStatus
	LastWarehouseId *string
}

func (uc *Usecase) ListShipments(ctx context.Context, params ListShipmentsParams) ([]*entity.Shipment, error) {
	if params.Status != nil && params.LastWarehouseId != nil {
		shipments, err := uc.ShipmentDg.ListShipmentsByStatusAndLastWarehouse(ctx, *params.Status, *params.LastWarehouseId, params.Limit, params.Offset)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list shipments")
		}
		return shipments, nil
	} else if params.Status == nil && params.LastWarehouseId != nil {
		shipments, err := uc.ShipmentDg.ListShipmentsByLastWarehouse(ctx, *params.LastWarehouseId, params.Limit, params.Offset)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list shipments")
		}
		return shipments, nil
	} else if params.Status != nil && params.LastWarehouseId == nil {
		shipments, err := uc.ShipmentDg.ListShipmentsByStatus(ctx, *params.Status, params.Limit, params.Offset)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list shipments")
		}
		return shipments, nil
	} else {
		shipments, err := uc.ShipmentDg.ListShipments(ctx, params.Limit, params.Offset)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list shipments")
		}
		return shipments, nil
	}
}
