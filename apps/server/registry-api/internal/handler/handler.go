package handler

import (
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
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

type MiddlewareParameters struct {
	AuthGuard      fiber.Handler
	AdminGuard     fiber.Handler
	WarehouseGuard fiber.Handler
}

func (h *Handler) Mount(r fiber.Router, middlewares MiddlewareParameters) {
	if r == nil {
		panic("router is nil")
	}

	// Health check
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Shipment API!")
	})

	// CRUD Warehouse Connections
	warehouseConnectionGroup := r.Group("/warehouse-connection", middlewares.AuthGuard, middlewares.AdminGuard)
	warehouseConnectionGroup.Get("/", h.GetWarehouseConnection)
	warehouseConnectionGroup.Get("/list", h.ListWarehouseConnections)
	warehouseConnectionGroup.Post("/", h.CreateWarehouseConnection)
	warehouseConnectionGroup.Put("/", h.UpdateWarehouseConnection)
	warehouseConnectionGroup.Delete("/:id", h.RevokeWarehouseConnection)

	// Updating Warehouse Endpoint
	warehouseEndpointGroup := r.Group("/warehouse-endpoint", middlewares.AuthGuard, middlewares.WarehouseGuard)
	warehouseEndpointGroup.Get("/", h.GetWarehouseEndpoint)
	warehouseEndpointGroup.Get("/list", h.ListWarehouseEndpoints)
	warehouseEndpointGroup.Put("/", h.UpdateWarehouseEndpoint)
}
