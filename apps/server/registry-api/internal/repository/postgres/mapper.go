package postgres

import (
	registry_database "github.com/drowningtoast/glip/apps/server/registry-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
)

func mapWarehouseConnectionModelToEntity(connection *registry_database.WarehouseConnection) *entity.WarehouseConnection {
	return &entity.WarehouseConnection{
		Id:          int(connection.ID),
		WarehouseId: connection.WarehouseID,
		ApiKey:      connection.ApiKey,
		Name:        connection.Name,
		Status:      entity.WarehouseConnectionStatus(connection.Status),
		CreatedAt:   &connection.CreatedAt.Time,
		UpdatedAt:   &connection.UpdatedAt.Time,
		LastUsedAt:  &connection.LastUsedAt.Time,
	}
}
