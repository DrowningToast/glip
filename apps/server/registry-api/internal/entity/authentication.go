package entity

type AuthenticationType string

const (
	AuthenticationTypeWarehouse AuthenticationType = "WAREHOUSE"
	AuthenticationTypeAdmin     AuthenticationType = "ADMIN"
)
