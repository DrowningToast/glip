package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type WarehouseStatus string

const (
	WarehouseStatusActive   WarehouseStatus = "ACTIVE"
	WarehouseStatusInactive WarehouseStatus = "INACTIVE"
)

type Warehouse struct {
	Id              string          `json:"id" validate:"required"`
	Name            string          `json:"name" validate:"required"`
	Location        string          `json:"location" validate:"required"`
	Country         string          `json:"country" validate:"required"`
	City            string          `json:"city" validate:"required"`
	TotalCapacity   decimal.Decimal `json:"total_capacity"`
	CurrentCapacity decimal.Decimal `json:"current_capacity"`
	Description     *string         `json:"description" validate:"required"`
	Status          WarehouseStatus `json:"status" validate:"required"`

	EndPoint *string
}

type WarehouseConnectionStatus string

const (
	WarehouseConnectionStatusActive   WarehouseConnectionStatus = "ACTIVE"
	WarehouseConnectionStatusInactive WarehouseConnectionStatus = "INACTIVE"
	WarehouseConnectionStatusRevoked  WarehouseConnectionStatus = "REVOKED"
)

type WarehouseConnection struct {
	Id string `json:"id"`
	// The warehouse id that the connection is for
	WarehouseId string                    `json:"warehouse_id"`
	ApiKey      string                    `json:"api_key"`
	Name        string                    `json:"name"`
	Status      WarehouseConnectionStatus `json:"status"`
	CreatedAt   *time.Time                `json:"created_at"`
	UpdatedAt   *time.Time                `json:"updated_at"`
	LastUsedAt  *time.Time                `json:"last_used_at"`
}
