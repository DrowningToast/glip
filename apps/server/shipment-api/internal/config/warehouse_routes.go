package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

// Region -> Destination -> Departure -> Route details
type WarehouseRoutes map[string]map[string]map[string]struct {
	Route []string `json:"route" validate:"required"`
}

func GetWarehouseRoutes() (*WarehouseRoutes, error) {
	var routesConfig WarehouseRoutes

	rawFile, err := os.Open("./config/warehouse_routes.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open warehouse_routes.json")
	}

	defer rawFile.Close()

	byteValue, err := io.ReadAll(rawFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read warehouse_routes.json")
	}

	if err := json.Unmarshal(byteValue, &routesConfig); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal warehouse_routes.json")
	}

	return &routesConfig, nil
}
