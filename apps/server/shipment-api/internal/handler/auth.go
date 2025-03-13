package handler

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) AuthenticateWarehouseConnection(ctx *fiber.Ctx) error {
	var body struct {
		Key string `json:"key" validate:"required"`
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return errors.Wrap(errs.ErrInvalidBody, err.Error())
	}

	if body.Key == "" {
		return errors.Wrap(errs.ErrInvalidBody, "key is required")
	}

	token, err := h.Uc.CreateWarehouseConnectionSession(ctx.Context(), body.Key)
	if err != nil {
		return errors.Wrap(err, "failed to create warehouse connection session")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			JWT string `json:"jwt"`
		}{
			JWT: *token,
		},
	})
}

func (h *Handler) AuthenticateAdmin(ctx *fiber.Ctx) error {
	var body struct {
		Key string `json:"key"`
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return errors.Wrap(errs.ErrInvalidBody, err.Error())
	}

	if body.Key == "" {
		return errors.Wrap(errs.ErrInvalidBody, "key is required")
	}

	token, err := h.Uc.CreateAdminApiSession(ctx.Context(), body.Key)
	if err != nil {
		return errors.Wrap(err, "failed to create admin api session")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			JWT string `json:"jwt"`
		}{
			JWT: *token,
		},
	})
}
