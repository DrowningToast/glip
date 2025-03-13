package entity

import (
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
