package handler

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetMyProfileAsCustomer(ctx *fiber.Ctx) error {
	userContext := ctx.UserContext().Value(usecase.UserContextKey{})
	jwt, ok := userContext.(*entity.JWTSession)
	if !ok {
		return errors.Wrap(errs.ErrInternal, "failed to parse your user context session")
	}
	if jwt.Role != entity.ConnectionTypeUser {
		return errors.Wrap(errs.ErrUnauthorized, "session is not customer role")
	}

	account, customer, err := h.Uc.GetMyProfileAsCustomer(ctx.Context(), jwt)
	if err != nil {
		return errors.Wrap(err, "failed to get customer profile")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Account  entity.Account  `json:"account"`
			Customer entity.Customer `json:"customer"`
		}{
			Account:  *account,
			Customer: *customer,
		},
	})
}

func (h *Handler) GetMyProfileAsWarehouseConnection(ctx *fiber.Ctx) error {
	userContext := ctx.UserContext().Value(usecase.UserContextKey{})
	jwt, ok := userContext.(*entity.JWTSession)
	if !ok {
		return errors.Wrap(errs.ErrInternal, "failed to parse your user context session")
	}
	if jwt.Role != entity.ConnectionTypeWarehouse {
		return errors.Wrap(errs.ErrUnauthorized, "session is not warehouse connection role")
	}

	apiKey := ctx.Query("api-key")
	if apiKey == "" {
		return errors.Wrap(errs.ErrForbidden, "empty api key")
	}

	warehouseConn, err := h.Uc.GetMyProfileAsWarehouseConnection(ctx.Context(), jwt, apiKey)
	if err != nil {
		return errors.Wrap(err, "unable to get warehouse connection info")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			WarehouseConnection entity.WarehouseConnection `json:"warehouse_connection"`
		}{
			WarehouseConnection: *warehouseConn,
		},
	})
}
