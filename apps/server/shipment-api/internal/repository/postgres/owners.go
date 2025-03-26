package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/internal/utils/pgmapper"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

var _ datagateway.OwnersDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateShipmentOwner(ctx context.Context, payload *entity.Customer) (*entity.Customer, error) {
	if payload == nil {
		return nil, errors.Wrap(errs.ErrInternal, "owner is nil")
	}

	log.Debug(payload)

	owner, err := r.queries.CreateShipmentOwner(ctx, shipment_database.CreateShipmentOwnerParams{
		Name:      payload.Name,
		Email:     payload.Email,
		Phone:     pgmapper.MapStringPtrToPgText(payload.Phone),
		Address:   pgmapper.MapStringPtrToPgText(payload.Address),
		AccountID: pgmapper.MapIntPtrToPgInt4(payload.AccountId),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, "owner already exists")
		}
		return nil, errors.Wrap(err, "failed to create shipment owner")
	}

	log.Debug(owner)

	return mapShipmentOwnerModelToEntity(&owner), nil
}

func (r *PostgresRepository) GetShipmentOwnerById(ctx context.Context, id int) (*entity.Customer, error) {
	owner, err := r.queries.GetShipmentOwnerById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get shipment owner by id")
	}

	return mapShipmentOwnerModelToEntity(&owner), nil
}

func (r *PostgresRepository) GetShipmentOwnerByEmail(ctx context.Context, email string) (*entity.Customer, error) {
	owner, err := r.queries.GetShipmentOwnerByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get shipment owner by email")
	}

	return mapShipmentOwnerModelToEntity(&owner), nil
}

func (r *PostgresRepository) GetShipmentOwnerByAccountId(ctx context.Context, id int) (*entity.Customer, error) {
	owner, err := r.queries.GetShipmentOwnerByAccountId(ctx, pgmapper.MapIntPtrToPgInt4(&id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get shipment owner by account id")
	}
	return mapShipmentOwnerModelToEntity(&owner), nil
}

func (r *PostgresRepository) ListShipmentOwners(ctx context.Context, limit int, offset int) ([]*entity.Customer, error) {
	owners, err := r.queries.ListShipmentOwners(
		ctx,
		shipment_database.ListShipmentOwnersParams{
			ReturnOffset: pgtype.Int4{Int32: int32(offset), Valid: true},
			ReturnLimit:  pgtype.Int4{Int32: int32(limit), Valid: true},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list shipment owners")
	}

	return lo.Map(owners, func(owner shipment_database.Owner, _ int) *entity.Customer {
		return mapShipmentOwnerModelToEntity(&owner)
	}), nil
}

func (r *PostgresRepository) UpdateShipmentOwner(ctx context.Context, ownerPtr *entity.Customer) (*entity.Customer, error) {
	owner, err := r.queries.UpdateShipmentOwner(ctx, shipment_database.UpdateShipmentOwnerParams{
		ID:        int32(ownerPtr.Id),
		Name:      ownerPtr.Name,
		Email:     ownerPtr.Email,
		Phone:     pgmapper.MapStringPtrToPgText(ownerPtr.Phone),
		Address:   pgmapper.MapStringPtrToPgText(ownerPtr.Address),
		AccountID: pgmapper.MapIntPtrToPgInt4(ownerPtr.AccountId),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update shipment owner")
	}

	return mapShipmentOwnerModelToEntity(&owner), nil
}

func (r *PostgresRepository) SoftDeleteShipmentOwner(ctx context.Context, id int) error {
	err := r.queries.SoftDeleteShipmentOwner(ctx, int32(id))
	if err != nil {
		return errors.Wrap(err, "failed to soft delete shipment owner")
	}

	return nil
}
