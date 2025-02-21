package middleware

import (
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

func NewAuthGuard(u *usecase.Usecase) fiber.Handler {
	return func(c fiber.Ctx) error {
		bearerString := c.Get("Authorization")
		if bearerString == "" {
			return errors.Wrap(errs.ErrUnauthorized, "missing authorization header")
		}

		splitedTokenString := strings.Split(bearerString, " ")
		if len(splitedTokenString) != 2 {
			return errors.Wrap(errs.ErrUnauthorized, "invalid authorization header")
		}

		tokenString := splitedTokenString[1]

		session, err := u.VerifyJWT(c.Context(), tokenString)
		if err != nil {
			return errors.Wrap(errs.ErrUnauthorized, err.Error())
		}

		if !session.Role.Valid() {
			return errors.Wrap(errs.ErrUnauthorized, "invalid role")
		}

		context := u.InitUserContext(c.Context(), session)
		c.SetContext(context)

		return c.Next()
	}
}
