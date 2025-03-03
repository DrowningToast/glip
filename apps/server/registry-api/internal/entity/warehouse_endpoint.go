package entity

import "time"

type WarehouseEndpoint struct {
	WarehouseId string    `json:"warehouse_id"`
	Endpoint    string    `json:"endpoint"`
	UpdatedAt   time.Time `json:"updated_at"`
}
