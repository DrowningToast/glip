package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

func (u *Usecase) CreateCustomer(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	return u.CustomerDg.CreateShipmentOwner(ctx, customer)
}

type GetCustomerQuery struct {
	Id    *int
	Email *string
}

func (u *Usecase) GetCustomer(ctx context.Context, query GetCustomerQuery) (*entity.Customer, error) {
	if query.Id != nil {
		customer, err := u.CustomerDg.GetShipmentOwnerById(ctx, *query.Id)
		if err != nil {
			return nil, err
		}
		return customer, nil
	}
	if query.Email != nil {
		customer, err := u.CustomerDg.GetShipmentOwnerByEmail(ctx, *query.Email)
		if err != nil {
			return nil, err
		}
		return customer, nil
	}

	return nil, errors.Wrap(errs.ErrInvalidArgument, "id or email is required")
}

// If both field are provided, the id will be used
type ListCustomersQuery struct {
	Limit  int
	Offset int
}

func (u *Usecase) ListCustomers(ctx context.Context, query *ListCustomersQuery) ([]*entity.Customer, error) {
	return u.CustomerDg.ListShipmentOwners(ctx, query.Limit, query.Offset)
}

func (u *Usecase) UpdateCustomer(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	if customer == nil {
		return nil, errors.Wrap(errs.ErrInvalidArgument, "customer id is required")
	}

	return u.CustomerDg.UpdateShipmentOwner(ctx, customer)
}

func (u *Usecase) DeleteCustomer(ctx context.Context, id int) error {
	customer, err := u.CustomerDg.GetShipmentOwnerById(ctx, id)
	if err != nil {
		return errors.Wrap(errs.ErrNotFound, "failed to get customer")
	}

	if customer == nil {
		return errors.Wrap(errs.ErrNotFound, "customer not found")
	}

	return u.CustomerDg.SoftDeleteShipmentOwner(ctx, id)
}
