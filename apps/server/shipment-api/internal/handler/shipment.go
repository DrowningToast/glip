package handler

import (
	"fmt"
	"log"
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

	// Check who made this shipment
	userContext := ctx.UserContext().Value(usecase.UserContextKey{})
	if userContext == nil {
		return errors.Wrap(errs.ErrForbidden, "User is not authenticated")
	}
	session, ok := userContext.(*entity.JWTSession)
	if !ok {
		return errors.Wrap(errs.ErrInternal, "Invalid context data type")
	}

	switch session.Role {
	case entity.ConnectionTypeUser:
		username := session.Id
		s, err := h.Uc.CreateShipment(ctx.Context(), *body.Shipment, &username)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to create shipment by customer: %s", username))
		}
		return ctx.JSON(common.HTTPResponse{
			Result: struct {
				Shipment *entity.Shipment `json:"shipment"`
			}{
				Shipment: s,
			},
		})
	case entity.ConnectionTypeRoot:
		s, err := h.Uc.CreateShipment(ctx.Context(), *body.Shipment, nil)
		if err != nil {
			return errors.Wrap(err, "failed to create shipment as root")
		}
		return ctx.JSON(common.HTTPResponse{
			Result: struct {
				Shipment *entity.Shipment `json:"shipment"`
			}{
				Shipment: s,
			},
		})

	default:
		return errors.Wrap(errs.ErrUnauthorized, "invalid role to create shipment")
	}
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

	shipments, err := h.Uc.ListShipments(ctx.Context(), usecase.ListShipmentsParams{
		Status:          query.Status,
		LastWarehouseId: query.LastWarehouseId,
		Limit:           *query.Limit,
		Offset:          *query.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "failed to list shipments")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: common.PaginatedResult[*entity.Shipment]{
			Items:  shipments,
			Count:  len(shipments),
			Offset: *query.Offset,
			Limit:  *query.Limit,
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

// sent by account
func (h *Handler) ListShipmentsByAccountUser(ctx *fiber.Ctx) error {
	var query struct {
		Limit  *int                   `query:"limit"`
		Offset *int                   `query:"offset"`
		Status *entity.ShipmentStatus `query:"status"`
	}
	if query.Limit == nil {
		limit := 100
		query.Limit = &limit
	}

	if query.Offset == nil {
		offset := 0
		query.Offset = &offset
	}

	session := ctx.UserContext().Value(usecase.UserContextKey{}).(*entity.JWTSession)
	log.Println(session)
	if session == nil {
		return errors.Wrap(errs.ErrUnauthorized, "account not found")
	}

	shipments, err := h.Uc.ListShipmentsByAccountUser(ctx.Context(), *&usecase.ListShipmentsByAccountUser{
		Limit:    *query.Limit,
		Offset:   *query.Offset,
		Status:   query.Status,
		Username: session.Id,
	})
	if err != nil {
		return errors.Wrap(err, "failed to list shipments")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: common.PaginatedResult[*entity.Shipment]{
			Items:  shipments,
			Count:  len(shipments),
			Offset: *query.Offset,
			Limit:  *query.Limit,
		},
	})
}

// sent by shipment owner
// expect request body: {shipment_id: shipment_id, email: email}
func (h *Handler) TrackShipment(ctx *fiber.Ctx) error {
	var body struct {
		ShipmentId int    `json:"shipment_id" validate:"required"`
		Email      string `json:"email" validate:"required"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(errs.ErrInvalidBody, err.Error())
	}

	userContext := ctx.UserContext().Value(usecase.UserContextKey{})
	if userContext != nil {
		// search by admin
		session, ok := userContext.(*entity.JWTSession)
		if !ok {
			return errors.Wrap(errs.ErrInternal, "failed to parse user context value")
		}
		if session.Role == entity.ConnectionTypeRoot {
			shipment, err := h.Uc.GetShipmentByOwner(ctx.Context(), usecase.GetShipmentByOwnerParams{
				ShipmentId: body.ShipmentId,
				Email:      body.Email,
			}, true)
			if err != nil {
				return errors.Wrap(err, "failed to track shipment")
			}
			return ctx.JSON(common.HTTPResponse{
				Result: struct {
					Shipment *entity.Shipment `json:"shipment"`
				}{
					Shipment: shipment,
				},
			})
		}
	}

	// search by customer
	shipment, err := h.Uc.GetShipmentByOwner(ctx.Context(), usecase.GetShipmentByOwnerParams{
		ShipmentId: body.ShipmentId,
		Email:      body.Email,
	}, false)
	if err != nil {
		return errors.Wrap(err, "failed to track shipment")
	}
	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Shipment *entity.Shipment `json:"shipment"`
		}{
			Shipment: shipment,
		},
	})
}
