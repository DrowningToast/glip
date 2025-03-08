package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/pkg/errors"
)

type WarehouseRegion string

const (
	WarehouseRegionNA   WarehouseRegion = "NA"
	WarehouseRegionEU   WarehouseRegion = "EU"
	WarehouseRegionAPAC WarehouseRegion = "APAC"
)

type WarehouseRegions struct {
	NA   map[string]*WarehouseRegion `json:"zone-NA" validate:"required"`
	EU   map[string]*WarehouseRegion `json:"zone-EU" validate:"required"`
	APAC map[string]*WarehouseRegion `json:"zone-APAC" validate:"required"`
}

func (r WarehouseRegions) GetWarehouseRegion(warehouseId string) (*WarehouseRegion, error) {
	wInNA := r.NA[warehouseId]
	if wInNA != nil {
		region := *wInNA
		return &region, nil
	}

	wInEU := r.EU[warehouseId]
	if wInEU != nil {
		region := *wInEU
		return &region, nil
	}

	wInAPAC := r.APAC[warehouseId]
	if wInAPAC != nil {
		region := *wInAPAC
		return &region, nil
	}

	return nil, errors.Wrap(errs.ErrInvalidArgument, "warehouse not found")
}

func GetWarehouseRegions() (*WarehouseRegions, error) {
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
