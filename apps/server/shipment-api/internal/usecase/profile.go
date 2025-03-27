package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

func (u Usecase) GetMyProfileAsCustomer(ctx context.Context, jwt *entity.JWTSession) (*entity.Account, *entity.Customer, error) {
	if jwt == nil {
		return nil, nil, errors.Wrap(errs.ErrUnauthorized, "empty session")
	}
	if jwt.Role != entity.ConnectionTypeUser {
		return nil, nil, errors.Wrap(errs.ErrForbidden, "not customer role")
	}
	account, err := u.AccountDg.GetAccountByUsername(ctx, jwt.Id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot get account info")
	}
	if account == nil {
		return nil, nil, errors.Wrap(errs.ErrNotFound, "account info not found from the username given")
	}
	customer, err := u.CustomerDg.GetShipmentOwnerByAccountId(ctx, account.Id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot get customer info")
	}
	if customer == nil {
		return nil, nil, errors.Wrap(errs.ErrNotFound, "customer info not found from given account id")
	}
	return account, customer, nil
}

func (u Usecase) GetMyProfileAsWarehouseConnection(ctx context.Context, jwt *entity.JWTSession, apiKey string) (*entity.WarehouseConnection, error) {
	warehouseconnection, err := u.WarehouseConnDg.GetWarehouseConnectByApiKey(ctx, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get warehouse connection")
	}

	if warehouseconnection.Id != jwt.Id {
		return nil, errors.Wrap(errs.ErrUnauthorized, "incorrect warehouse connection credentials")
	}

	return warehouseconnection, nil
}
