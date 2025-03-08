package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	globalConfig "github.com/drowningtoast/glip/apps/server/internal/config"
	"github.com/drowningtoast/glip/apps/server/internal/utils"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/config"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/handler"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/middlewares"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/repository/etcd"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/repository/postgres"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/usecase"
	"github.com/gofiber/fiber/v3"
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

	// registry pg conn
	pgConn, err := configuration.RegistryPgConfig.NewConnection(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to registry pg: %v\n", err)
	}

	// registry etcd conn
	etcdConn, err := configuration.EtcdConfig.NewConnectionWithRootUser()
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v\n", err)
	}

	// init repository
	pgRepo := postgres.New(pgConn)
	etcdRepo, err := etcd.New(etcdConn)
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v\n", err)
	}

	// init usecase
	uc := usecase.NewUsecase(&usecase.UsecaseParams{
		Config:           configuration,
		WarehouseRegions: etcdRepo.Regions,

		WarehouseConnectionDg: pgRepo,
		WarehouseEndpointDg:   etcdRepo,
	})

	// init h
	h := handler.New(handler.HandlerNewParams{
		Usecase: uc,
	})

	// init server
	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Registry Name",
		ErrorHandler: utils.FiberErrHandler,
	})

	// app.Use(logger.New())

	// middlewares
	authGuard := middlewares.NewAuthGuard(uc)
	adminGuard := middlewares.NewRoleGuard(uc, entity.AuthenticationTypeAdmin)
	warehouseGuard := middlewares.NewRoleGuard(uc, entity.AuthenticationTypeWarehouse)

	// mount
	v1Router := app.Group("/v1")
	h.Mount(v1Router, handler.MiddlewareParameters{
		AuthGuard:      authGuard,
		AdminGuard:     adminGuard,
		WarehouseGuard: warehouseGuard,
	})

	// Start server with graceful shutdown
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", configuration.RegistryPort)); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
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
