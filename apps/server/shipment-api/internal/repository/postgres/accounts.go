package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/internal/utils/pgmapper"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

var _ datagateway.AccountsDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateAccount(ctx context.Context, accountPtr *entity.Account) (*entity.Account, error) {
	if accountPtr == nil {
		return nil, errors.Wrap(errs.ErrInternal, "account is nil")
	}
	account := *accountPtr

	createdAccount, err := r.queries.CreateAccount(ctx, shipment_database.CreateAccountParams{
		Username: account.Username,
		Password: account.Password,
		Email:    account.Email,
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, "account already exists")
		}
		return nil, errors.Wrap(err, "failed to create account")
	}

	return mapAccountModelToEntity(&createdAccount), nil
}

func (r *PostgresRepository) GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error) {
	account, err := r.queries.GetAccountByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get account by username")
	}

	return mapAccountModelToEntity(&account), nil
}

func (r *PostgresRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	account, err := r.queries.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get account by email")
	}

	return mapAccountModelToEntity(&account), nil
}

func (r *PostgresRepository) GetAccountById(ctx context.Context, id int) (*entity.Account, error) {
	account, err := r.queries.GetAccountById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get account by id")
	}

	return mapAccountModelToEntity(&account), nil
}

func (r *PostgresRepository) ListAccounts(ctx context.Context, limit int, offset int) ([]*entity.Account, error) {
	accounts, err := r.queries.ListAccounts(ctx, shipment_database.ListAccountsParams{
		ReturnOffset: pgtype.Int4{Int32: int32(offset)},
		ReturnLimit:  pgtype.Int4{Int32: int32(limit)},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list accounts")
	}

	return lo.Map(accounts, func(account shipment_database.Account, _ int) *entity.Account {
		return mapAccountModelToEntity(&account)
	}), nil
}

func (r *PostgresRepository) UpdateAccount(ctx context.Context, params *datagateway.UpdateAccountParams) (*entity.Account, error) {
	updatedAccount, err := r.queries.UpdateAccount(ctx, shipment_database.UpdateAccountParams{
		ID:       int32(params.Id),
		Username: pgmapper.MapStringPtrToPgText(params.Username),
		Password: pgmapper.MapStringPtrToPgText(params.Password),
		Email:    pgmapper.MapStringPtrToPgText(params.Email),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errs.ErrNotFound, "account not found")
		}
		return nil, errors.Wrap(err, "failed to update account")
	}

	return mapAccountModelToEntity(&updatedAccount), nil
}

func (r *PostgresRepository) SoftDeleteAccount(ctx context.Context, id int) error {
	if err := r.queries.SoftDeleteAccount(ctx, int32(id)); err != nil {
		return errors.Wrap(err, "failed to soft delete account")
	}

	return nil
}
