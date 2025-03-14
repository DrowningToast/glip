package middleware

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type RoleGuardParams struct {
	Root      bool
	User      bool
	Warehouse bool
}

func NewRoleGuard(u *usecase.Usecase) func(params RoleGuardParams) fiber.Handler {
	return func(params RoleGuardParams) fiber.Handler {
		return func(c *fiber.Ctx) error {
			session := c.UserContext().Value(usecase.UserContextKey{}).(*entity.JWTSession)
			if session == nil {
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			isAllowed := func() bool {
				switch session.Role {
				case entity.ConnectionTypeRoot:
					return params.Root
				case entity.ConnectionTypeUser:
					return params.User
				case entity.ConnectionTypeWarehouse:
					return params.Warehouse
				default:
					return false
				}
			}()

			if !isAllowed {
				return errors.Wrap(errs.ErrForbidden, "invalid role")
			}

			return c.Next()
		}
	}
}
