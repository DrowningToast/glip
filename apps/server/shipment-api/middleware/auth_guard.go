package middleware

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewAuthGuard(u *usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session := u.GetUserContext(c)
		if session == nil {
			return errors.Wrap(errs.ErrUnauthorized, "User not authenticated")
		}

		return c.Next()
	}
}
