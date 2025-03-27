package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type OwnersDataGateway interface {
	CreateShipmentOwner(ctx context.Context, owner *entity.Customer) (*entity.Customer, error)
	GetShipmentOwnerByAccountId(ctx context.Context, id int) (*entity.Customer, error)
	GetShipmentOwnerById(ctx context.Context, id int) (*entity.Customer, error)
	GetShipmentOwnerByEmail(ctx context.Context, email string) (*entity.Customer, error)
	ListShipmentOwners(ctx context.Context, limit int, offset int) ([]*entity.Customer, error)
	UpdateShipmentOwner(ctx context.Context, owner *entity.Customer) (*entity.Customer, error)
	SoftDeleteShipmentOwner(ctx context.Context, id int) error
}
