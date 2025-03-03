package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	globalConfig "github.com/drowningtoast/glip/apps/server/internal/config"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/config"
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handle interrupting signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		cancel()
	}()

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

	// create root user
	etcdConn.UserAdd(ctx, config.InventoryRegistryConfig.Username, config.InventoryRegistryConfig.RootPassword)
	etcdConn.UserGrantRole(ctx, config.InventoryRegistryConfig.Username, "root")

	etcdConn, err = config.InventoryRegistryConfig.NewConnectionWithRootUser()
	if err != nil {
		log.Fatalf("Failed to connect to etcd with root user: %v\n", err)
	}

	etcdConn, err = config.InventoryRegistryConfig.NewConnectionWithRootUser()
	if err != nil {
		log.Fatalf("Failed to connect to etcd with root user: %v\n", err)
	}

	_, err = etcdConn.Auth.AuthEnable(ctx)
	if err != nil {
		log.Fatalf("Failed to enable auth: %v\n", err)
	}

	// warehouseRepoPtr, err := etcd.New(etcdConn)
	// if err != nil {
	// 	log.Fatalf("Failed to create etcd repository: %v\n", err)
	// }

	// warehouseRepo := *warehouseRepoPtr

	// // init pg conn
	// registryPgConn, err := config.InventoryPgConfig.NewConnection(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to registry pg: %v\n", err)
	// }
	// pgRepo := postgres.New(registryPgConn)
	// warehouseConnections, err := pgRepo.ListWarehouseConnections(ctx, 99999999, 0)
	// if err != nil {
	// 	log.Fatalf("Failed to list warehouse connections: %v\n", err)
	// }

	// bytes, err := json.Marshal(warehouseRepo.Regions)
	// if err != nil {
	// 	log.Fatalf("Failed to marshal warehouse regions: %v\n", err)
	// }

	// // setup timeout ctx
	// timeoutCtx, _ := context.WithTimeout(ctx, 10*time.Second)

	// etcdConnWg := sync.WaitGroup{}

	// // add region users and roles
	// var regions map[string]map[string]entity.Warehouse
	// if err := json.Unmarshal(bytes, &regions); err != nil {
	// 	log.Fatalf("Failed to unmarshal warehouse regions: %v\n", err)
	// }
	// for _, region := range regions {
	// 	for _, w := range region {
	// 		warehouseConn, found := lo.Find(warehouseConnections, func(wc *entity.WarehouseConnection) bool {
	// 			return wc.WarehouseId == w.Id
	// 		})
	// 		if !found {
	// 			continue
	// 		}

	// 		etcdConnWg.Add(1)
	// 		go func() {
	// 			defer etcdConnWg.Done()

	// 			// create role for warehouse
	// 			roleName := w.Id
	// 			etcdConn.RoleAdd(ctx, roleName)
	// 			key := "/warehouse/" + w.Id
	// 			etcdConn.RoleGrantPermission(ctx, roleName, key+"/", key+"0", clientv3.PermissionType(clientv3.PermReadWrite))

	// 			if warehouseConn == nil {
	// 				log.Fatalf("Warehouse connection not found for warehouse: %v\n", w.Id)
	// 			}
	// 			user := &User{
	// 				Name:     w.Name,
	// 				Password: warehouseConn.ApiKey,
	// 				Roles:    []string{"warehouse"},
	// 			}

	// 			etcdConn.UserAdd(ctx, user.Name, user.Password)
	// 			etcdConn.UserGrantRole(ctx, user.Name, roleName)
	// 		}()
	// 	}
	// }

	// done := make(chan struct{})
	// go func() {
	// 	etcdConnWg.Wait()
	// 	close(done)
	// }()

	// // wait for setup to complete or timeout
	// select {
	// case <-done:
	// 	return
	// case <-ctx.Done():
	// 	log.Println("Setup interrupted")
	// 	os.Exit(1)
	// case <-timeoutCtx.Done():
	// 	log.Println("Setup timed out")
	// 	cancel()
	// 	os.Exit(1)
	// }

	log.Println("Setup complete")
}
