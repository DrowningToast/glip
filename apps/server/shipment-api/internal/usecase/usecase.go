package usecase

import (
	"github.com/drowningtoast/glip/apps/server/config"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
)

type Usecase struct {
	Config *config.Config

	WarehouseDg           datagateway.WarehouseDataGateway
	WarehouseConnectionDg datagateway.WarehouseConnectionDataGateway
	ShipmentDg            datagateway.ShipmentDataGateway
	CustomerDg            datagateway.OwnersDataGateway
	CarrierDg             datagateway.CarrierDataGateway
	AlertDg               datagateway.AlertDataGateway
}

type UsecaseParams struct {
	Config *config.Config

	WarehouseDg           datagateway.WarehouseDataGateway
	WarehouseConnectionDg datagateway.WarehouseConnectionDataGateway
	ShipmentDg            datagateway.ShipmentDataGateway
	CustomerDg            datagateway.OwnersDataGateway
	CarrierDg             datagateway.CarrierDataGateway
	AlertDg               datagateway.AlertDataGateway
}

func NewUsecase(params *UsecaseParams) *Usecase {
	return &Usecase{
		Config: params.Config,

		WarehouseDg:           params.WarehouseDg,
		WarehouseConnectionDg: params.WarehouseConnectionDg,
		ShipmentDg:            params.ShipmentDg,
		CustomerDg:            params.CustomerDg,
		CarrierDg:             params.CarrierDg,
		AlertDg:               params.AlertDg,
	}
}
