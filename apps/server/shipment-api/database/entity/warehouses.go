package entity

import "time"

type WarehouseStatus string

const (
	WarehouseStatusActive   WarehouseStatus = "ACTIVE"
	WarehouseStatusInactive WarehouseStatus = "INACTIVE"
)

type Warehouse struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Location string          `json:"location"`
	Capacity int             `json:"capacity"`
	Status   WarehouseStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
