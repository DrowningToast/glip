package registry_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/pkg/errors"
)

var _ datagateway.WarehouseConnectionDataGateway = (*RegistryApiRepository)(nil)

type WarehouseConnectionResponse struct {
	WarehouseConnection entity.WarehouseConnection `json:"warehouse_connection"`
}

func (r *RegistryApiRepository) GetWarehouseConnectByApiKey(ctx context.Context, apiKey string) (*entity.WarehouseConnection, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s/v1/warehouse-connection", r.Url, r.Port), nil)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}
	req.Header.Set("Authorization", r.ApiKey)
	req.Header.Set("AuthType", "ADMIN")
	q := req.URL.Query()
	q.Add("api-key", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	defer resp.Body.Close()

	// check if found or not
	if resp.StatusCode == http.StatusNotFound {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(errs.ErrInternal, err.Error())
		}
		return nil, errors.Wrap(errs.ErrUnauthorized, string(body))
	}
	if resp.StatusCode != http.StatusOK {
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(errs.ErrInternal, err.Error())
		}
		return nil, errors.Wrap(errs.ErrInternal, "error while querying the database")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	respBody := struct {
		Result WarehouseConnectionResponse `json:"result"`
	}{}
	if err := json.Unmarshal(body, &respBody); err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	warehouseConn := respBody.Result.WarehouseConnection

	return &warehouseConn, nil
}
