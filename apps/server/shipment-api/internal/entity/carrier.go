package entity

import "time"

type CarrierStatus string

const (
	CarrierStatusActive   CarrierStatus = "ACTIVE"
	CarrierStatusInactive CarrierStatus = "INACTIVE"
)

type Carrier struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	ContactPerson *string       `json:"contact_person"`
	ContactPhone  *string       `json:"contact_phone"`
	Email         *string       `json:"email"`
	Description   *string       `json:"description"`
	Status        CarrierStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
