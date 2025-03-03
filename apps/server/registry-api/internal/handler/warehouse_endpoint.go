package handler

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
	"github.com/gofiber/fiber/v3"
)

// req query: warehouse_id
func (h *Handler) GetWarehouseEndpoint(c fiber.Ctx) error {
	warehouseId := c.Query("warehouse_id")
	if warehouseId == "" {
		return errors.Wrap(errs.ErrInvalidArgument, "warehouse_id is required")
	}

	warehouseEndpoint, err := h.uc.GetWarehouseEndpoint(c.Context(), warehouseId)
	if err != nil {
		return errors.Wrap(err, "failed to get warehouse endpoint")
	}

	return c.JSON(warehouseEndpoint)
}

func (h *Handler) ListWarehouseEndpoints(c fiber.Ctx) error {
	warehouseEndpoints, err := h.uc.ListWarehouseEndpoints(c.Context())
	if err != nil {
		return errors.Wrap(err, "failed to list warehouse endpoints")
	}

	return c.JSON(warehouseEndpoints)
}

func (h *Handler) UpdateWarehouseEndpoint(c fiber.Ctx) error {
	authType, warehousePtr, err := h.uc.GetRequestContext(c.Context())
	if err != nil {
		return errors.Wrap(err, "failed to get request context")
	}
	if authType != entity.AuthenticationTypeWarehouse {
		return errors.Wrap(errs.ErrUnauthorized, "unauthorized")
	}
	if warehousePtr == nil {
		return errors.Wrap(errs.ErrUnauthorized, "unauthorized")
	}

	source := c.RequestCtx().RemoteAddr()

	if err = h.uc.UpdateWarehouseEndpoint(c.Context(), warehousePtr.WarehouseId, source.String()); err != nil {
		return errors.Wrap(err, "failed to update warehouse endpoint")
	}

	return c.JSON(common.EmptyHTTPResponse)
}
