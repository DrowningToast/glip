package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type ShipmentStatus string

const (
	ShipmentStatusWaitingForPickup   ShipmentStatus = "WAITING_FOR_PICKUP_TO_WAREHOUSE"
	ShipmentStatusInTransitOnTheWay  ShipmentStatus = "IN_TRANSIT_ON_THE_WAY"
	ShipmentStatusArrivedAtWarehouse ShipmentStatus = "ARRIVED_AT_WAREHOUSE"
	ShipmentStatusDelivered          ShipmentStatus = "DELIVERED"
	ShipmentStatusCancelled          ShipmentStatus = "CANCELLED"
	ShipmentStatusLost               ShipmentStatus = "LOST"
)

func (s ShipmentStatus) String() string {
	return string(s)
}

func (s ShipmentStatus) Valid() bool {
	switch s {
	case ShipmentStatusInTransitOnTheWay,
		ShipmentStatusWaitingForPickup,
		ShipmentStatusDelivered,
		ShipmentStatusCancelled:
		return true
	default:
		return false
	}
}

type Shipment struct {
	Id int `json:"id"`
	// Route is a list of warehouse ids that the shipment will pass through
	Route []string `json:"route"`

	// Last visited warehouse id
	LastWarehouseId *string `json:"last_warehouse_id"`

	// Departure warehouse id
	DepartureWarehouseId string  `json:"departure_warehouse_id"`
	DepartureAddress     *string `json:"departure_address"`

	// Destination warehouse id
	DestinationWarehouseId string `json:"destination_warehouse_id"`
	DestinationAddress     string `json:"destination_address"`

	// Who managed the shipment, (account entity)
	CreatedBy int `json:"created_by"`

	// Who the shipment belongs to, (customer entity)
	OwnerId int `json:"owner_id"`

	Status              ShipmentStatus  `json:"status"`
	TotalWeight         decimal.Decimal `json:"total_weight"`
	TotalVolume         decimal.Decimal `json:"total_volume"`
	SpecialInstructions *string         `json:"special_instructions"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
