package entity

type WarehouseStatus string

const (
	WarehouseStatusActive   WarehouseStatus = "ACTIVE"
	WarehouseStatusInactive WarehouseStatus = "INACTIVE"
)

type Warehouse struct {
	Id       string          `json:"id" validate:"required"`
	Name     string          `json:"name" validate:"required"`
	Location string          `json:"location" validate:"required"`
	Country  string          `json:"country" validate:"required"`
	City     string          `json:"city" validate:"required"`
	Status   WarehouseStatus `json:"status" validate:"required"`

	EndPoint *string
}
