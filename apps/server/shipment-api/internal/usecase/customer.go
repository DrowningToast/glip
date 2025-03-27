package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type CreateCustomerParams struct {
	// Account info
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`

	// Customer info
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}

func (u *Usecase) CreateCustomer(ctx context.Context, params CreateCustomerParams) (*entity.Account, *entity.Customer, error) {
	// Create account
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	payloadAccount := &entity.Account{
		Username: params.Username,
		Password: string(hash),
		Email:    params.Email,
	}

	account, err := u.AccountDg.CreateAccount(ctx, payloadAccount)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create account in usecase level")
	}

	// create payloadCustomer
	payloadCustomer := &entity.Customer{
		Name:      params.Name,
		Email:     params.Email,
		Phone:     params.Phone,
		Address:   params.Address,
		AccountId: &account.Id,
	}
	customer, err := u.CustomerDg.CreateShipmentOwner(ctx, payloadCustomer)
	if err != nil {
		if err := u.AccountDg.SoftDeleteAccount(ctx, payloadAccount.Id); err != nil {
			return nil, nil, errors.Wrap(err, "creating customer record failed, and failed to revert the account entity creation!")
		}
	}

	return account, customer, nil
}

type GetCustomerQuery struct {
	Id        *int
	Email     *string
	AccountId *int
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
