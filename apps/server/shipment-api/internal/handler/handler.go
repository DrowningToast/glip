package handler

import (
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/middleware"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Uc *usecase.Usecase
}

type HandlerNewParams struct {
	Usecase *usecase.Usecase
}

func New(params HandlerNewParams) Handler {
	return Handler{
		Uc: params.Usecase,
	}
}

type MiddlewareParameters struct {
	AuthGuard fiber.Handler
	RoleGuard func(params middleware.RoleGuardParams) fiber.Handler
}

func (h *Handler) Mount(r fiber.Router, middlewares MiddlewareParameters) {
	if r == nil {
		panic("router is nil")
	}

	// Health check
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Shipment API!")
	})

	// Auth
	authGroup := r.Group("/auth")
	authGroup.Post("/warehouse", h.AuthenticateWarehouseConnection)
	authGroup.Post("/admin", h.AuthenticateAdmin)
	authGroup.Post("/customer", h.AuthenticateCustomerConnection)

	// Customer
	customerGroup := r.Group("/customer")
	customerGroup.Get("/", h.GetCustomer, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Root: true}))
	customerGroup.Get("/list", h.ListCustomers, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Root: true}))
	customerGroup.Post("/", h.CreateCustomer)
	customerGroup.Put("/", h.UpdateCustomer, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Root: true}))
	customerGroup.Delete("/:id", h.DeleteCustomer, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Root: true}))

	// Shipment
	shipmentGroup := r.Group("/shipment")

	// list by customer
	shipmentGroup.Get("/customer/list", h.ListShipmentsByAccountUser, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{User: true}))
	// track by shipment owner
	shipmentGroup.Post("/track", h.TrackShipment)
	// Create shipment
	shipmentGroup.Post("/", h.CreateShipment, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Warehouse: true, Root: true}))
	// List shipments
	shipmentGroup.Get("/list", h.ListShipments, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Warehouse: true}))
	// Get shipment by id
	shipmentGroup.Get("/:shipment_id", h.GetShipment, middlewares.AuthGuard, middlewares.RoleGuard(middleware.RoleGuardParams{Warehouse: true}))
}
