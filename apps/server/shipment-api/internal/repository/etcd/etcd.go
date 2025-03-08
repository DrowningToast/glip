package etcd

import (
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type WarehouseRegions struct {
	NA   map[string]entity.Warehouse `json:"zone-NA" validate:"required"`
	EU   map[string]entity.Warehouse `json:"zone-EU" validate:"required"`
	APAC map[string]entity.Warehouse `json:"zone-APAC" validate:"required"`
}

type EtcdRepository struct {
	Client *clientv3.Client
}

func New(client *clientv3.Client) (*EtcdRepository, error) {
	r := &EtcdRepository{
		Client: client,
	}

	return r, nil
}
