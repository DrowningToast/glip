package middleware

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewRoleGuard(u *usecase.Usecase) func(role entity.ConnectionType) fiber.Handler {
	return func(role entity.ConnectionType) fiber.Handler {
		return func(c *fiber.Ctx) error {
			session := u.GetUserContext(c)
			if session == nil {
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			if session.Role != role {
				return errors.Wrap(errs.ErrForbidden, "invalid role")
			}

			return c.Next()
		}
	}
}
