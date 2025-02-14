package entity

import "time"

type ShipmentStatus string

const (
	ShipmentStatusInTransitOnTheWay    ShipmentStatus = "IN_TRANSIT_ON_THE_WAY"
	ShipmentStatusInTransitInWarehouse ShipmentStatus = "IN_TRANSIT_IN_WAREHOUSE"
	ShipmentStatusDelivered            ShipmentStatus = "DELIVERED"
	ShipmentStatusCancelled            ShipmentStatus = "CANCELLED"
)

func (s ShipmentStatus) String() string {
	return string(s)
}

func (s ShipmentStatus) Valid() bool {
	switch s {
	case ShipmentStatusInTransitOnTheWay,
		ShipmentStatusInTransitInWarehouse,
		ShipmentStatusDelivered,
		ShipmentStatusCancelled:
		return true
	default:
		return false
	}
}

type Shipment struct {
	Id                  int            `json:"id"`
	Route               []int          `json:"route"`
	LastWarehouseId     *int           `json:"last_warehouse_id"`
	DestinationAddress  string         `json:"destination_address"`
	CarrierId           *int           `json:"carrier_id"`
	ScheduledDeparture  time.Time      `json:"scheduled_departure"`
	ScheduledArrival    time.Time      `json:"scheduled_arrival"`
	ActualDeparture     *time.Time     `json:"actual_departure"`
	ActualArrival       *time.Time     `json:"actual_arrival"`
	Status              ShipmentStatus `json:"status"`
	TotalWeight         float64        `json:"total_weight"`
	TotalVolume         float64        `json:"total_volume"`
	SpecialInstructions *string        `json:"special_instructions"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
