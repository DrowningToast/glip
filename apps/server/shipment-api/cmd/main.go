package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	globalConfig "github.com/drowningtoast/glip/apps/server/internal/config"
	"github.com/drowningtoast/glip/apps/server/internal/services"
	"github.com/drowningtoast/glip/apps/server/internal/utils"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/config"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/handler"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/repository/postgres"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/repository/rabbitmq"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/drowningtoast/glip/apps/server/shipment-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create background context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config
	globalConfiguration, err := globalConfig.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}
	if globalConfiguration == nil {
		log.Fatal("Failed to load config")
	}

	// Load local config
	configuration, err := config.ExtendConfig(globalConfiguration, nil)
	if err != nil {
		log.Fatalf("Failed to load local config: %v\n", err)
	}

	// shipment pg conn
	pgConn, err := configuration.ShipmentPgConfig.NewConnection(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to shipment pg: %v\n", err)
	}

	// init rabbitmq conn
	rabbitmqConn, err := configuration.RabbitMQConfig.NewConnection(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to rabbitmq: %v\n", err)
	}
	defer rabbitmqConn.Close()

	rabbitmqChannel, err := services.NewRabbmitMQChannel(ctx, rabbitmqConn)
	if err != nil {
		log.Fatalf("Failed to open rabbitmq channel: %v\n", err)
	}
	defer rabbitmqChannel.Close()

	// init repository
	pgRepo := postgres.New(pgConn)
	rabbitmqRepo := rabbitmq.NewRepository(&configuration.WarehouseRegions, &configuration.RabbitMQConfig, rabbitmqChannel)

	// init usecase
	uc := usecase.NewUsecase(&usecase.UsecaseParams{
		Config: configuration,

		ShipmentQueueDg: rabbitmqRepo,

		ShipmentDg: pgRepo,
		CustomerDg: pgRepo,
		CarrierDg:  pgRepo,
		AlertDg:    pgRepo,
	})

	// init h
	h := handler.New(handler.HandlerNewParams{
		Usecase: uc,
	})

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Shipment API",
		ErrorHandler: utils.FiberErrHandler,
	})

	// app.Use(logger.New())

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

	// Listen to rabbitmq
	go func() {
		errorChan := make(chan error)
		go h.Uc.WatchShipmentUpdates(ctx, errorChan)
		for err := range errorChan {
			log.Printf("Error watching shipment updates: %v\n", err)
		}
	}()

	// Start server with graceful shutdown
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", configuration.ShipmentPort)); err != nil {
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
