package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/common"
	"github.com/drowningtoast/glip/apps/server/config"
	"github.com/drowningtoast/glip/apps/server/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/handler"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/repository/postgres"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/drowningtoast/glip/apps/server/shipment-api/middleware"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create background context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config
	configuration := config.Load()
	if configuration == nil {
		log.Fatal("Failed to load config")
	}

	// shipment pg conn
	pgConn, err := configuration.ShipmentPgConfig.NewConnection(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to shipment pg: %v\n", err)
	}

	// init repository
	pgRepo := postgres.New(pgConn)

	// init usecase
	uc := usecase.NewUsecase(&usecase.UsecaseParams{
		Config: configuration,

		ShipmentDg:            pgRepo,
		WarehouseDg:           pgRepo,
		WarehouseConnectionDg: pgRepo,
		CustomerDg:            pgRepo,
		CarrierDg:             pgRepo,
		AlertDg:               pgRepo,
	})

	// init h
	h := handler.New(handler.HandlerNewParams{
		Usecase: uc,
	})

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Shipment API",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			customErr := errs.Err{}
			if errors.As(err, &customErr) {
				code = customErr.StatusCode
				return c.Status(code).JSON(common.HTTPResponse{
					Code:    customErr.Error(),
					Message: err.Error(),
				})
			}

			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				code = fiberErr.Code
				return c.Status(code).JSON(common.HTTPResponse{
					Code:    fiberErr.Error(),
					Message: fiberErr.Message,
				})
			}

			return c.Status(code).JSON(common.HTTPResponse{
				Code:    errs.ErrInternal.Code,
				Message: err.Error(),
			})
		},
	})

	// Middleware
	initContextMiddleware := middleware.NewInitContextMiddleware(uc)
	authGuard := middleware.NewAuthGuard(uc)
	roleGuard := middleware.NewRoleGuard(uc)

	// Mount
	v1Router := app.Group("/v1", initContextMiddleware)
	h.Mount(v1Router, handler.MiddlewareParameters{
		AuthGuard: authGuard,
		RoleGuard: roleGuard,
	})

	// Start server with graceful shutdown
	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}

	// Cancel context
	cancel()
}
