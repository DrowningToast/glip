package middlewares

import (
	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHeader struct {
	Authorization string `json:"Authorization" validate:"required"`
	AuthType      string `json:"AuthType" validate:"required"`
}

func NewAuthGuard(uc *usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := AuthHeader{
			Authorization: c.Get("Authorization"),
			AuthType:      c.Get("AuthType"),
		}

		if authHeader.AuthType != "ADMIN" && authHeader.AuthType != "WAREHOUSE" {
			return errors.Wrap(errs.ErrUnauthorized, "invalid auth type")
		}

		ctx, err := uc.Authenticate(c.Context(), entity.AuthenticationType(authHeader.AuthType), authHeader.Authorization)
		if err != nil {
			return errors.Wrap(errs.ErrUnauthorized, err.Error())
		}

		c.SetUserContext(ctx)

		return c.Next()
	}
}
