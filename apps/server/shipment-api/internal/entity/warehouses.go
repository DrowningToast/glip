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
	Id              int             `json:"id"`
	Name            string          `json:"name"`
	Location        string          `json:"location"`
	Country         string          `json:"country"`
	City            string          `json:"city"`
	TotalCapacity   decimal.Decimal `json:"total_capacity"`
	CurrentCapacity decimal.Decimal `json:"current_capacity"`
	Description     *string         `json:"description"`
	Status          WarehouseStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
