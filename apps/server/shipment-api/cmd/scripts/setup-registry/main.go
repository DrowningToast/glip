package main

import (
	"context"
	"log"

	globalConfig "github.com/drowningtoast/glip/apps/server/internal/config"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/config"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/repository/etcd"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/repository/postgres"
	"github.com/samber/lo"
)

type Permission struct {
	Path   string
	Type   string // "read", "write", or "readwrite"
	Prefix bool
}

type Role struct {
	Name        string
	Permissions []Permission
}

type User struct {
	Name     string
	Password string
	Roles    []string
}

// script for setting up the inventory service registry

func SetupServiceRegistry() {
	ctx, _ := context.WithCancel(context.Background())

	globalConfiguration, err := globalConfig.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	config, err := config.ExtendConfig(globalConfiguration, nil)
	if err != nil {
		log.Fatalf("Failed to load local config: %v\n", err)
	}

	// init etcd conn
	etcdConn, err := config.InventoryRegistryConfig.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v\n", err)
	}

	warehouseRepoPtr, err := etcd.New(etcdConn)
	if err != nil {
		log.Fatalf("Failed to create etcd repository: %v\n", err)
	}

	warehouseRepo := *warehouseRepoPtr

	// init pg conn
	shipmentPgConn, err := config.ShipmentPgConfig.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to shipment pg: %v\n", err)
	}
	pgRepo := postgres.New(shipmentPgConn)
	warehouseConnections, err := pgRepo.ListWarehouseConnections(ctx, 99999999, 0)
	if err != nil {
		log.Fatalf("Failed to list warehouse connections: %v\n", err)
	}

	// remove every users from the registry

	lo.MapValues(warehouseRepo.Regions.NA, func(w entity.Warehouse, k string) *User {
	})
}
