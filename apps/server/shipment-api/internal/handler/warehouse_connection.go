package handler

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) CreateWarehouseConnection(c fiber.Ctx) error {
	var body struct {
		WarehouseConnection struct {
			WarehouseId int                              `json:"warehouse_id" validate:"required"`
			ApiKey      string                           `json:"api" validate:"required"`
			Name        string                           `json:"name" validate:"required"`
			Status      entity.WarehouseConnectionStatus `json:"status" validate:"required"`
		} `json:"warehouse_connection" validate:"required"`
	}

	if err := c.Bind().Body(&body); err != nil {
		return errors.Wrap(errs.ErrInternal, "error while binding body")
	}

	payload := entity.WarehouseConnection{
		WarehouseId: body.WarehouseConnection.WarehouseId,
		ApiKey:      body.WarehouseConnection.ApiKey,
		Name:        body.WarehouseConnection.Name,
		Status:      body.WarehouseConnection.Status,
	}

	warehouseConn, err := h.uc.CreateWarehouseConnection(c.Context(), payload)
	if err != nil {
		return errors.Wrap(err, "error while creating warehouse connection")
	}

	if warehouseConn == nil {
		return errors.Wrap(errs.ErrInternal, "error while creating warehouse connection")
	}

	return c.Status(fiber.StatusCreated).JSON(common.HTTPResponse{
		Result: struct {
			WarehouseConnection entity.WarehouseConnection `json:"warehouse_connection"`
		}{
			WarehouseConnection: *warehouseConn,
		},
	})
}

type GetWarehouseConnectionRequestQuery struct {
	Id     *int    `json:"id,omitempty"`
	ApiKey *string `json:"api_key,omitempty"`
}

// accepts a warehouse connection request and returns a warehouse connection
func (h *Handler) GetWarehouseConnection(c fiber.Ctx) error {
	var query GetWarehouseConnectionRequestQuery

	if err := c.Bind().Query(&query); err != nil {
		return errors.Wrap(errs.ErrInternal, "error while binding query")
	}

	if query.Id == nil && query.ApiKey == nil {
		return errors.Wrap(errs.ErrInvalidQueryString, "id or api key is required")
	}

	warehouseConn, err := h.uc.GetWarehouseConnection(c.Context(), usecase.GetWarehouseConnectionQuery{
		Id:     query.Id,
		ApiKey: query.ApiKey,
	})
	if err != nil {
		return errors.Wrap(err, "error while querying the database")
	}

	if warehouseConn == nil {
		return errors.Wrap(errs.ErrNotFound, "warehouse connection not found")
	}

	return c.JSON(common.HTTPResponse{
		Result: struct {
			Warehouse entity.WarehouseConnection `json:"warehouse"`
		}{
			Warehouse: *warehouseConn,
		},
	})
}

type ListWarehouseConnectionsQuery struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`

	Status *entity.WarehouseConnectionStatus `json:"status,omitempty"`
}

func (h *Handler) ListWarehouseConnections(c fiber.Ctx) error {
	var query ListWarehouseConnectionsQuery

	if err := c.Bind().Query(&query); err != nil {
		return errors.Wrap(errs.ErrInternal, err.Error())
	}

	limit, offset := 100, 0
	if query.Limit != nil {
		limit = *query.Limit
	}
	if query.Offset != nil {
		offset = *query.Offset
	}

	warehouseConnPtrs, err := h.uc.ListWarehouseConnections(c.Context(), usecase.ListWarehouseConnectionsQuery{
		Offset: offset,
		Limit:  limit,
		Status: query.Status,
	})
	if err != nil {
		return errors.Wrap(err, "error while querying the database")
	}

	return c.JSON(common.HTTPResponse{
		Result: common.PaginatedResult[*entity.WarehouseConnection]{
			Count:  len(warehouseConnPtrs),
			Items:  warehouseConnPtrs,
			Offset: offset,
			Limit:  limit,
		},
	})
}

func (h *Handler) UpdateWarehouseConnection(c fiber.Ctx) error {
	var body struct {
		WarehouseConnection struct {
			WarehouseId int                              `json:"warehouse_id" validate:"required"`
			ApiKey      string                           `json:"api" validate:"required"`
			Name        string                           `json:"name" validate:"required"`
			Status      entity.WarehouseConnectionStatus `json:"status" validate:"required"`
		} `json:"warehouse_connection" validate:"required"`
	}

	if err := c.Bind().Body(&body); err != nil {
		return errors.Wrap(errs.ErrInternal, err.Error())
	}

	payload := entity.WarehouseConnection{
		WarehouseId: body.WarehouseConnection.WarehouseId,
		ApiKey:      body.WarehouseConnection.ApiKey,
		Name:        body.WarehouseConnection.Name,
		Status:      body.WarehouseConnection.Status,
	}

	warehouseConn, err := h.uc.UpdateWarehouseConnection(c.Context(), payload)
	if err != nil {
		return errors.Wrap(err, "error while updating warehouse connection")
	}

	if warehouseConn == nil {
		return errors.Wrap(errs.ErrNotFound, "warehouse connection not found")
	}

	return c.JSON(common.HTTPResponse{
		Result: struct {
			WarehouseConnection entity.WarehouseConnection `json:"warehouse_connection"`
		}{
			WarehouseConnection: *warehouseConn,
		},
	})
}

// accepts params "id" and returns a warehouse connection
func (h *Handler) RevokeWarehouseConnection(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.Wrap(errs.ErrInvalidArgument, "invalid id")
	}

	err = h.uc.RevokeWarehouseConnection(c.Context(), id)
	if err != nil {
		return errors.Wrap(err, "error while revoking warehouse connection")
	}

	return c.JSON(common.HTTPResponse{
		Result: common.EmptyHTTPResponse,
	})
}
