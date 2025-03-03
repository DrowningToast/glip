package utils

import (
	"errors"

	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/gofiber/fiber/v3"
)

func FiberErrHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	customErr := errs.Err{}
	if errors.As(err, &customErr) {
		code = customErr.StatusCode
		return c.Status(code).JSON(common.HTTPResponse{
			Code:    customErr.Error(),
			Message: err.Error(),
		})
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		return c.Status(code).JSON(common.HTTPResponse{
			Code:    fiberErr.Error(),
			Message: fiberErr.Message,
		})
	}

	return c.Status(code).JSON(common.HTTPResponse{
		Code:    errs.ErrInternal.Code,
		Message: err.Error(),
	})
}
