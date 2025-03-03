package usecase

import (
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/config"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/repository/etcd"
)

type Usecase struct {
	Config           *config.RegistryConfig
	WarehouseRegions etcd.WarehouseRegions

	WarehouseConnectionDg datagateway.WarehouseConnectionDataGateway
	WarehouseEndpointDg   datagateway.WarehouseEndpointDataGateway
}

type UsecaseParams struct {
	Config           *config.RegistryConfig
	WarehouseRegions etcd.WarehouseRegions

	WarehouseConnectionDg datagateway.WarehouseConnectionDataGateway
	WarehouseEndpointDg   datagateway.WarehouseEndpointDataGateway
}

func NewUsecase(params *UsecaseParams) *Usecase {
	return &Usecase{
		Config: params.Config,

		WarehouseConnectionDg: params.WarehouseConnectionDg,
		WarehouseEndpointDg:   params.WarehouseEndpointDg,
		WarehouseRegions:      params.WarehouseRegions,
	}
}
