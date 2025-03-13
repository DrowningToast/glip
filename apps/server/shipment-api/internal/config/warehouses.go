package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type WarehouseRegion string

const (
	WarehouseRegionNA   WarehouseRegion = "zoneNA"
	WarehouseRegionEU   WarehouseRegion = "zoneEU"
	WarehouseRegionAPAC WarehouseRegion = "zoneAPAC"
)

type WarehouseRegions struct {
	NA   map[string]*entity.Warehouse `json:"zoneNA" validate:"required"`
	EU   map[string]*entity.Warehouse `json:"zoneEU" validate:"required"`
	APAC map[string]*entity.Warehouse `json:"zoneAPAC" validate:"required"`
}

// Key is warehouseId, value is warehouse
func (r WarehouseRegions) ToMap() map[string]*entity.Warehouse {
	warehouses := make(map[string]*entity.Warehouse)

	for _, warehouse := range r.NA {
		warehouses[warehouse.Id] = warehouse
	}

	for _, warehouse := range r.EU {
		warehouses[warehouse.Id] = warehouse
	}

	for _, warehouse := range r.APAC {
		warehouses[warehouse.Id] = warehouse
	}

	return warehouses
}

func (r WarehouseRegions) GetWarehouseRegion(warehouseId string) (WarehouseRegion, error) {
	wInNA := r.NA[warehouseId]
	if wInNA != nil {
		return WarehouseRegionNA, nil
	}

	wInEU := r.EU[warehouseId]
	if wInEU != nil {
		return WarehouseRegionEU, nil
	}

	wInAPAC := r.APAC[warehouseId]
	if wInAPAC != nil {
		return WarehouseRegionAPAC, nil
	}

	return "", errors.Wrap(errs.ErrInvalidArgument, "warehouse not found")
}

func (r WarehouseRegions) GetWarehouseByCity(city string) (*entity.Warehouse, error) {
	warehouses := r.ToMap()

	warehouse, ok := lo.Find(lo.Values(warehouses), func(warehouse *entity.Warehouse) bool {
		return warehouse.City == city
	})
	if !ok {
		return nil, errors.Wrap(errs.ErrInvalidArgument, "warehouse not found")
	}

	return warehouse, nil
}

func (r WarehouseRegions) GetWarehouseById(warehouseId string) (*entity.Warehouse, error) {
	// search through every region
	wInNA := r.NA[warehouseId]
	if wInNA != nil {
		warehouse := *wInNA
		return &warehouse, nil
	}

	wInEU := r.EU[warehouseId]
	if wInEU != nil {
		warehouse := *wInEU
		return &warehouse, nil
	}

	wInAPAC := r.APAC[warehouseId]
	if wInAPAC != nil {
		warehouse := *wInAPAC
		return &warehouse, nil
	}

	return nil, errors.Wrap(errs.ErrInvalidArgument, "warehouse not found")
}

func LoadWarehouseConfig() (*WarehouseRegions, error) {
	var regions WarehouseRegions

	rawFile, err := os.Open("./config/warehouses.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open warehouses.json")
	}

	defer rawFile.Close()

	byteValue, err := io.ReadAll(rawFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read warehouses.json")
	}

	if err := json.Unmarshal(byteValue, &regions); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal warehouses.json")
	}

	return &regions, nil
}
