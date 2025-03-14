package usecase

import (
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/config"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
)

type Usecase struct {
	Config *config.ShipmentConfig

	ShipmentQueueDg datagateway.ShipmentQueueDataGateway

	AccountDg  datagateway.AccountsDataGateway
	ShipmentDg datagateway.ShipmentDataGateway
	CustomerDg datagateway.OwnersDataGateway
	CarrierDg  datagateway.CarrierDataGateway
	AlertDg    datagateway.AlertDataGateway
}

type UsecaseParams struct {
	Config *config.ShipmentConfig

	ShipmentQueueDg datagateway.ShipmentQueueDataGateway

	ShipmentDg datagateway.ShipmentDataGateway
	CustomerDg datagateway.OwnersDataGateway
	CarrierDg  datagateway.CarrierDataGateway
	AlertDg    datagateway.AlertDataGateway
}

func NewUsecase(params *UsecaseParams) *Usecase {
	return &Usecase{
		Config: params.Config,

		ShipmentQueueDg: params.ShipmentQueueDg,

		ShipmentDg: params.ShipmentDg,
		CustomerDg: params.CustomerDg,
		CarrierDg:  params.CarrierDg,
		AlertDg:    params.AlertDg,
	}
}
