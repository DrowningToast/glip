package entity

import "time"

type WarehouseConnectionStatus string

const (
	WarehouseConnectionStatusActive   WarehouseConnectionStatus = "ACTIVE"
	WarehouseConnectionStatusInactive WarehouseConnectionStatus = "INACTIVE"
	WarehouseConnectionStatusRevoked  WarehouseConnectionStatus = "REVOKED"
)

type WarehouseConnection struct {
	Id int `json:"id"`
	// The warehouse id that the connection is for
	WarehouseId int                       `json:"warehouse_id"`
	ApiKey      string                    `json:"api_key"`
	Name        string                    `json:"name"`
	Status      WarehouseConnectionStatus `json:"status"`
	CreatedAt   *time.Time                `json:"created_at"`
	UpdatedAt   *time.Time                `json:"updated_at"`
	LastUsedAt  *time.Time                `json:"last_used_at"`
	// The account id of the user who created the connection
	CreatedBy int `json:"created_by"`
}
