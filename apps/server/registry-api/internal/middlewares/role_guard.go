package middlewares

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

// Requires a valid session and a specific role
func NewRoleGuard(uc *usecase.Usecase, authType entity.AuthenticationType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authType, _, err := uc.GetRequestContext(c)
		if err != nil {
			return errors.Wrap(errs.ErrUnauthorized, err.Error())
		}

		if authType != authType {
			return errors.Wrap(errs.ErrUnauthorized, "invalid authentication type")
		}

		return c.Next()
	}
}
