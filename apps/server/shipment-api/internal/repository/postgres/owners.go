package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

var _ datagateway.OwnersDataGateway = (*PostgresRepository)(nil)

func (r *PostgresRepository) CreateShipmentOwner(ctx context.Context, ownerPtr *entity.Customer) (*entity.Customer, error) {
	if ownerPtr == nil {
		return nil, errors.Wrap(errs.ErrInternal, "owner is nil")
	}

	owner, err := r.queries.CreateShipmentOwner(ctx, shipment_database.CreateShipmentOwnerParams{
		Name:    ownerPtr.Name,
		Email:   ownerPtr.Email,
		Phone:   mapStringPtrToPgText(ownerPtr.Phone),
		Address: mapStringPtrToPgText(ownerPtr.Address),
	})
	if err != nil {
		if checkPgErrCode(err, pgerrcode.UniqueViolation) {
			return nil, errors.Wrap(errs.ErrDuplicate, "owner already exists")
		}
		return nil, errors.Wrap(err, "failed to create shipment owner")
	}

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
		ID:      int32(ownerPtr.Id),
		Name:    ownerPtr.Name,
		Email:   ownerPtr.Email,
		Phone:   mapStringPtrToPgText(ownerPtr.Phone),
		Address: mapStringPtrToPgText(ownerPtr.Address),
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
