package middleware

import (
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

func NewInitContextMiddleware(uc *usecase.Usecase) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, err := uc.InjectSessionContext(c.Context(), c)
		if err != nil {
			return err
		}

		return c.Next()
	}
}
