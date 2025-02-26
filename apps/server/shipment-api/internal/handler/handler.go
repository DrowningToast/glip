package handler

import (
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
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

type MiddlewareParameters struct {
	AuthGuard fiber.Handler
	RoleGuard func(permission entity.ConnectionType) fiber.Handler
}

func (h *Handler) Mount(r fiber.Router, middlewares MiddlewareParameters) {
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

	// Customer
	customerGroup := r.Group("/customer", middlewares.AuthGuard)
	customerGroup.Get("/", h.GetCustomer)
	customerGroup.Get("/list", h.ListCustomers)
	customerGroup.Post("/", h.CreateCustomer, middlewares.RoleGuard(entity.ConnectionTypeRoot))
	customerGroup.Put("/", h.UpdateCustomer, middlewares.RoleGuard(entity.ConnectionTypeRoot))
	customerGroup.Delete("/:id", h.DeleteCustomer, middlewares.RoleGuard(entity.ConnectionTypeRoot))

	// Warehouse Connections
	warehouseConnectionGroup := r.Group("/warehouse-connection", middlewares.AuthGuard, middlewares.RoleGuard(entity.ConnectionTypeRoot))
	warehouseConnectionGroup.Get("/", h.GetWarehouseConnection)
	warehouseConnectionGroup.Get("/list", h.ListWarehouseConnections)
	warehouseConnectionGroup.Post("/", h.CreateWarehouseConnection)
	warehouseConnectionGroup.Put("/", h.UpdateWarehouseConnection)
	warehouseConnectionGroup.Delete("/:id", h.RevokeWarehouseConnection)
}
