package entity

import "time"

type AccountRole string

const (
	AccountRoleAdmin         AccountRole = "ADMIN"
	AccountRoleCarrierStaff  AccountRole = "CARRIER_STAFF"
	AccountRoleCarrierViewer AccountRole = "CARRIER_VIEWER"
	AccountRoleOwner         AccountRole = "OWNER"
)

type Account struct {
	Id       int         `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Role     AccountRole `json:"role"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
