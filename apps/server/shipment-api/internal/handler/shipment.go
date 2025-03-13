package handler

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateShipment(ctx *fiber.Ctx) error {
	var body struct {
		Shipment *usecase.CreateShipmentParams `json:"shipment" validate:"required"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(errs.ErrInvalidBody, err.Error())
	}

	shipment, err := h.Uc.CreateShipment(ctx.Context(), *body.Shipment)
	if err != nil {
		return errors.Wrap(err, "failed to create shipment")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Shipment *entity.Shipment `json:"shipment"`
		}{
			Shipment: shipment,
		},
	})
}

// request query: ?status=status&last_warehouse_id=warehouse_id&limit=limit&offset=offset
func (h *Handler) ListShipments(ctx *fiber.Ctx) error {
	var query struct {
		Status          *entity.ShipmentStatus `query:"status"`
		LastWarehouseId *string                `query:"last_warehouse_id"`
		Limit           *int                   `query:"limit"`
		Offset          *int                   `query:"offset"`
	}

	if err := ctx.QueryParser(&query); err != nil {
		return errors.Wrap(errs.ErrInvalidQueryString, err.Error())
	}

	if query.Limit == nil {
		limit := 100
		query.Limit = &limit
	}

	if query.Offset == nil {
		offset := 0
		query.Offset = &offset
	}

	shipments, err := h.Uc.ListShipments(ctx.Context(), *&usecase.ListShipmentsParams{
		Status:          query.Status,
		LastWarehouseId: query.LastWarehouseId,
		Limit:           *query.Limit,
		Offset:          *query.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "failed to list shipments")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Shipments []*entity.Shipment `json:"shipments"`
		}{
			Shipments: shipments,
		},
	})
}

// expect request param: shipment_id
func (h *Handler) GetShipment(ctx *fiber.Ctx) error {
	shipmentId := ctx.Params("shipment_id")
	if shipmentId == "" {
		return errors.Wrap(errs.ErrInvalidArgument, "shipment_id is required")
	}

	shipmentIdInt, err := strconv.Atoi(shipmentId)
	if err != nil {
		return errors.Wrap(errs.ErrInvalidArgument, "shipment_id is not a valid integer")
	}

	shipment, err := h.Uc.GetShipmentById(ctx.Context(), shipmentIdInt)
	if err != nil {
		return errors.Wrap(err, "failed to get shipment")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Shipment *entity.Shipment `json:"shipment"`
		}{
			Shipment: shipment,
		},
	})
}
