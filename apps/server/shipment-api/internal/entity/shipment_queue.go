package entity

import "github.com/rabbitmq/amqp091-go"

type ShipmentQueueType string

const (
	// Alert from the shipment api to the warehouse
	ShipmentQueueTypeOutbound ShipmentQueueType = "OUTBOUND"

	// Alert from the warehouse to the shipment api
	ShipmentQueueTypeInbound ShipmentQueueType = "INBOUND"
)

type ShipmentQueueStatus string

const (
	// Outbound, sent from the shipment api to the warehouse
	// Description: Expected shipment arrives at warehouse
	ShipmentQueueStatusIncomingShipment ShipmentQueueStatus = "INCOMING_SHIPMENT"

	// Inbound, sent from the warehouse to the shipment api

	// Description: Warehouse received shipment
	ShipmentQueueStatusWarehouseReceived ShipmentQueueStatus = "WAREHOUSE_RECEIVED"
	// Description: Warehouse departed from shipment
	ShipmentQueueStatusWarehouseDeparted ShipmentQueueStatus = "WAREHOUSE_DEPARTED"
	// Description: Shipment delivered
	ShipmentQueueStatusDelivered ShipmentQueueStatus = "DELIVERED"
)

func (s ShipmentQueueStatus) GetQueueType() ShipmentQueueType {
	switch s {
	case ShipmentQueueStatusIncomingShipment:
		return ShipmentQueueTypeOutbound
	}
	return ShipmentQueueTypeInbound
}

type ShipmentQueue struct {
	Shipment

	// Message queue sent from warehouse id to shipment api
	FromWarehouseId *string `json:"from_warehouse_id"`

	// Message queue sent from shipment api to warehouse id
	ToWarehouseId *string `json:"to_warehouse_id"`

	QueueType ShipmentQueueType `json:"type"`

	// Msg
	Msg *amqp091.Delivery `json:"msg,omitempty"`
}
