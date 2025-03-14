package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type CreateAccountParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (u *Usecase) CreateAccount(ctx context.Context, params CreateAccountParams) (*entity.Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := &entity.Account{
		Username: params.Username,
		Password: string(hash),
		Email:    params.Email,
	}

	return u.AccountDg.CreateAccount(ctx, account)
}

type GetAccountParams struct {
	Id       *int
	Email    *string
	Username *string
}

// GetAccount returns an account by id, email, or username
func (u *Usecase) GetAccount(ctx context.Context, params GetAccountParams) (*entity.Account, error) {
	if params.Id != nil {
		return u.AccountDg.GetAccountById(ctx, *params.Id)
	}
	if params.Email != nil {
		return u.AccountDg.GetAccountByEmail(ctx, *params.Email)
	}
	if params.Username != nil {
		return u.AccountDg.GetAccountByUsername(ctx, *params.Username)
	}

	return nil, errors.Wrap(errs.ErrInvalidArgument, "no account id or email or username provided")
}

type ListAccountsParams struct {
	Limit  int
	Offset int
}

func (u *Usecase) ListAccounts(ctx context.Context, params ListAccountsParams) ([]*entity.Account, error) {
	return u.AccountDg.ListAccounts(ctx, params.Limit, params.Offset)
}

type UpdateAccountParams struct {
	Id       int
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

func (u *Usecase) UpdateAccount(ctx context.Context, params UpdateAccountParams) (*entity.Account, error) {
	if params.Username != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*params.Username), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		hashStr := string(hash)
		params.Password = &hashStr
	}

	return u.AccountDg.UpdateAccount(ctx, &datagateway.UpdateAccountParams{
		Id:       params.Id,
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
	})
}
