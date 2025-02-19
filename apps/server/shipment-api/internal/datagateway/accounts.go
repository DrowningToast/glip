package datagateway

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type AccountsDataGateway interface {
	CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error)
	GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error)
	GetAccountById(ctx context.Context, id int) (*entity.Account, error)
	ListAccounts(ctx context.Context, limit int, offset int) ([]*entity.Account, error)
	UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error)
	SoftDeleteAccount(ctx context.Context, id int) error
}
