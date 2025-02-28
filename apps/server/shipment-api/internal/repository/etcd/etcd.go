package etcd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type WarehouseRegions struct {
	NA   map[string]entity.Warehouse `json:"zone-NA" validate:"required"`
	EU   map[string]entity.Warehouse `json:"zone-EU" validate:"required"`
	APAC map[string]entity.Warehouse `json:"zone-APAC" validate:"required"`
}

type EtcdRepository struct {
	Regions WarehouseRegions

	Client *clientv3.Client
}

func New(client *clientv3.Client) (*EtcdRepository, error) {
	r := &EtcdRepository{
		Client: client,
	}

	rawFile, err := os.Open("./config/warehouses.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open warehouses.json")
	}

	defer rawFile.Close()

	byteValue, err := io.ReadAll(rawFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read warehouses.json")
	}

	if err := json.Unmarshal(byteValue, &r.Regions); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal warehouses.json")
	}

	return r, nil
}
