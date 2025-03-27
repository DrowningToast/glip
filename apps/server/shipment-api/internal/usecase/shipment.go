package usecase

import (
	"context"
	"fmt"

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
	OwnerId             *int            `json:"owner_id"`
	TotalWeight         decimal.Decimal `json:"total_weight" validate:"required"`
	TotalVolume         decimal.Decimal `json:"total_volume" validate:"required"`
	SpecialInstructions *string         `json:"special_instructions,omitempty"`
}

// if the account owner is nil, it'll presume it's created by root
func (uc *Usecase) CreateShipment(ctx context.Context, shipment CreateShipmentParams, ownerAccountUsername *string) (*entity.Shipment, error) {
	if ownerAccountUsername != nil {
		account, err := uc.AccountDg.GetAccountByUsername(ctx, *ownerAccountUsername)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get account to verify the customer (sender)")
		}
		if account == nil {
			return nil, errors.Wrap(errs.ErrNotFound, "customer info not found from provided username")
		}
		customer, err := uc.CustomerDg.GetShipmentOwnerByAccountId(ctx, account.Id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get the customer info")
		}
		if customer == nil {
			return nil, errors.Wrap(errs.ErrNotFound, "no customer info binded to the account entity found")
		}
		shipment.OwnerId = &customer.Id
	} else {
		// Assume the root created this shipment for some reaso:46

		shipment.OwnerId = nil
	}

	if shipment.OwnerId != nil {
		log.Debug(*shipment.OwnerId)
	}

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

	var createdShipment *entity.Shipment
	if shipment.OwnerId == nil {
		// created by admin
		createdShipment, err = uc.ShipmentDg.CreateShipment(ctx, &entity.Shipment{
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
	} else {
		// created by customer
		createdShipment, err = uc.ShipmentDg.CreateShipmentWithOwner(ctx, &entity.Shipment{
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
			OwnerId:             shipment.OwnerId,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to create shipment")
		}
	}

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
	go uc.ShipmentQueueDg.WatchReceivedShipmentQueue(ctx, shipmentQueueChan, errorChan, ctx.Done())

	for {
		select {
		// incoming shipment queue
		case shipmentQueue := <-shipmentQueueChan:
			// validate shipment queue
			log.Debug(fmt.Sprintf("MSG TO SHIPMENT API, TYPE : %s", shipmentQueue.QueueType))
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
			if shipmentQueue.LastWarehouseId == nil {
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
					_, _, found := lo.FindIndexOf(oldShipment.Route, func(warehouseId string) bool {
						return warehouseId == *shipmentQueue.FromWarehouseId
					})
					if !found {
						log.Warnf("last warehouse id not found in route")
						errorChan <- errors.Wrap(errs.ErrNotFound, "last warehouse id not found in route")
						continue
					}

					// if currentIndex == len(oldShipment.Route)-1 {
					// 	log.Warnf("invalid shipment route")
					// 	errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment route")
					// 	continue
					// }
					// // Check if we're not skiping
					// if oldShipment.Route[currentIndex-1] != *shipmentQueue.LastWarehouseId {
					// 	log.Warnf("invalid shipment route")
					// 	errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment route")
					// 	continue
					// }
				}

				oldShipment.LastWarehouseId = shipmentQueue.FromWarehouseId

				oldShipment.Status = entity.ShipmentStatusArrivedAtWarehouse
				oldShipment, err = uc.ShipmentDg.UpdateShipment(ctx, oldShipment)
				if err != nil {
					log.Warnf("failed to update shipment: %v", err)
					errorChan <- errors.Wrap(err, "failed to update shipment")
					continue
				}
				shipmentQueue.Msg.Ack(false)
				break
			case entity.ShipmentStatusInTransitOnTheWay, entity.ShipmentStatusDelivered:
				if shipmentQueue.LastWarehouseId == nil || shipmentQueue.FromWarehouseId == nil || *shipmentQueue.FromWarehouseId != *shipmentQueue.LastWarehouseId {
					log.Warnf("invalid shipment status")
					errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment status")
					continue
				}
				if shipmentQueue.Status == entity.ShipmentStatusDelivered && *shipmentQueue.FromWarehouseId != oldShipment.DestinationWarehouseId {
					log.Warnf("invalid shipment status")
					errorChan <- errors.Wrap(errs.ErrInvalidArgument, "destination warehouse id doesn't match the from warehouse id value")
					continue
				}

				oldShipment.Status = shipmentQueue.Status
				oldShipment, err = uc.ShipmentDg.UpdateShipment(ctx, oldShipment)
				if err != nil {
					log.Warnf("failed to update shipment: %v", err)
					errorChan <- errors.Wrap(err, "failed to update shipment")
					continue
				}

				if shipmentQueue.Status == entity.ShipmentStatusInTransitOnTheWay {
					_, currentIndex, _ := lo.FindIndexOf(oldShipment.Route, func(warehouseId string) bool {
						return warehouseId == *oldShipment.LastWarehouseId
					})
					nextWarehouseId := oldShipment.Route[currentIndex+1]
					uc.ShipmentQueueDg.CreateToReceivedShipment(ctx, oldShipment, nextWarehouseId)
				}
				shipmentQueue.Msg.Ack(false)

				break
			default:
				log.Warnf("invalid shipment status")
				errorChan <- errors.Wrap(errs.ErrInvalidArgument, "invalid shipment status")
				continue
			}

			// oldShipment, err = uc.ShipmentDg.GetShipmentById(ctx, shipmentQueue.Id)
			// if err != nil {
			// 	log.Warnf("failed to get shipment: %v", err)
			// 	errorChan <- errors.Wrap(err, "failed to get shipment")
			// 	continue
			// }
			// if oldShipment == nil {
			// 	log.Warnf("shipment not found")
			// 	errorChan <- errors.Wrap(errs.ErrNotFound, "shipment not found")
			// 	continue
			// }

			// // update shipment status
			// uc.ShipmentDg.UpdateShipment(ctx, &entity.Shipment{
			// 	Id:     shipmentQueue.Id,
			// 	Status: entity.ShipmentStatus(shipmentQueue.Status),
			// })
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

type ListShipmentsByAccountUser struct {
	Limit    int
	Offset   int
	Status   *entity.ShipmentStatus
	Username string
}

func (uc *Usecase) ListShipmentsByAccountUser(ctx context.Context, params ListShipmentsByAccountUser) ([]*entity.Shipment, error) {
	shipments, err := uc.ShipmentDg.ListShipmentsByAccountUsername(ctx, params.Username, params.Limit, params.Offset, params.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list shipments")
	}
	return shipments, nil
}

type GetShipmentByOwnerParams struct {
	Email      string `json:"email" validate:"required"`
	ShipmentId int    `json:"id" validate:"required"`
}

func (uc *Usecase) GetShipmentByOwner(ctx context.Context, params GetShipmentByOwnerParams, bypassOwnerCheck bool) (*entity.Shipment, error) {
	shipment, err := uc.ShipmentDg.GetShipmentById(ctx, params.ShipmentId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shipment")
	}

	owner, err := uc.CustomerDg.GetShipmentOwnerByEmail(ctx, params.Email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get owner")
	}
	if owner == nil {
		return nil, errors.Wrap(errs.ErrNotFound, "owner not found")
	}

	if !bypassOwnerCheck && shipment.OwnerId != nil {
		if *shipment.OwnerId != owner.Id {
			return nil, errors.Wrap(errs.ErrNotFound, "shipment not found")
		}
	}

	return shipment, nil
}
