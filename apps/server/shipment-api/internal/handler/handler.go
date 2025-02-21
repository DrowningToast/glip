package handler

import (
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	uc *usecase.Usecase
}

type HandlerNewParams struct {
	Usecase *usecase.Usecase
}

func New(params HandlerNewParams) Handler {
	return Handler{
		uc: params.Usecase,
	}
}

func (h *Handler) Mount(r fiber.Router) {
	if r == nil {
		panic("router is nil")
	}

	// Health check
	r.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Shipment API!")
	})

	// Auth
	authGroup := r.Group("/auth")
	authGroup.Post("/warehouse", h.AuthenticateWarehouseConnection)
	authGroup.Post("/admin", h.AuthenticateAdmin)
}
