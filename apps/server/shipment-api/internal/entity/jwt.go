package entity

import "github.com/golang-jwt/jwt/v5"

type ConnectionType string

const (
	ConnectionTypeWarehouse ConnectionType = "WAREHOUSE"
	ConnectionTypeUser      ConnectionType = "USER"
	ConnectionTypeRoot      ConnectionType = "ROOT"
)

func (r *ConnectionType) String() string {
	return string(*r)
}

func (r *ConnectionType) Valid() bool {
	switch *r {
	case ConnectionTypeWarehouse, ConnectionTypeUser, ConnectionTypeRoot:
		return true
	default:
		return false
	}
}

type JWTSession struct {
	jwt.RegisteredClaims
	Id   string         `json:"id"`
	Role ConnectionType `json:"role"`
}
