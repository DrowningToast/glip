package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
)

type (
	RequestContextAuthTypeKey struct{}
	RequestContextSessionKey  struct{}
)

func (uc *Usecase) SetRequestContext(ctx context.Context, authType entity.AuthenticationType, session *entity.WarehouseConnection) (context.Context, error) {
	switch authType {
	case entity.AuthenticationTypeWarehouse:
		ctx = context.WithValue(ctx, RequestContextAuthTypeKey{}, authType)
		ctx = context.WithValue(ctx, RequestContextSessionKey{}, session)
		return ctx, nil

	case entity.AuthenticationTypeAdmin:
		return context.WithValue(ctx, RequestContextAuthTypeKey{}, authType), nil
	}

	return nil, errors.Wrap(errs.ErrUnauthorized, "invalid authentication type")
}

func (uc *Usecase) GetRequestContext(ctx context.Context) (entity.AuthenticationType, *entity.WarehouseConnection, error) {
	// Get auth type
	authTypeVal := ctx.Value(RequestContextAuthTypeKey{})
	if authTypeVal == nil {
		return "", nil, errors.Wrap(errs.ErrUnauthorized, "authentication type not found in context")
	}

	authType, ok := authTypeVal.(entity.AuthenticationType)
	if !ok {
		return "", nil, errors.Wrap(errs.ErrUnauthorized, "invalid authentication type in context")
	}

	// Get session (only required for warehouse auth type)
	if authType == entity.AuthenticationTypeWarehouse {
		sessionVal := ctx.Value(RequestContextSessionKey{})
		if sessionVal == nil {
			return "", nil, errors.Wrap(errs.ErrUnauthorized, "warehouse session not found in context")
		}

		session, ok := sessionVal.(*entity.WarehouseConnection)
		if !ok {
			return "", nil, errors.Wrap(errs.ErrUnauthorized, "invalid warehouse session in context")
		}
		return authType, session, nil
	}

	// For admin auth type, session is not required
	return authType, nil, nil
}

func (uc *Usecase) AuthenticateWarehouseConnection(ctx context.Context, apiKey string) (*entity.WarehouseConnection, error) {
	warehouseConn, err := uc.WarehouseConnectionDg.GetWarehouseConnectionByApiKey(ctx, apiKey)
	if err != nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, "error while authenticating warehouse connection")
	}

	if warehouseConn == nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, "warehouse connection not found")
	}

	if warehouseConn.Status != entity.WarehouseConnectionStatusActive {
		return nil, errors.Wrap(errs.ErrUnauthorized, "warehouse connection is not active")
	}

	return warehouseConn, nil
}

func (uc *Usecase) AuthenticateAdmin(ctx context.Context, apiSecretKey string) error {
	if len(apiSecretKey) == 0 {
		return errors.Wrap(errs.ErrUnauthorized, "api secret key is required")
	}

	if apiSecretKey != uc.Config.RegistryAuthConfig.APISecretKey {
		return errors.Wrap(errs.ErrUnauthorized, "invalid api secret key")
	}

	return nil
}

func (uc *Usecase) Authenticate(ctx context.Context, authType entity.AuthenticationType, apiSecretKey string) (context.Context, error) {
	switch authType {
	case entity.AuthenticationTypeWarehouse:
		warehouseConn, err := uc.AuthenticateWarehouseConnection(ctx, apiSecretKey)
		if err != nil {
			return nil, err
		}

		ctx, err = uc.SetRequestContext(ctx, authType, warehouseConn)
		if err != nil {
			return nil, errors.Wrap(errs.ErrInternal, err.Error())
		}

		return ctx, nil
	case entity.AuthenticationTypeAdmin:
		err := uc.AuthenticateAdmin(ctx, apiSecretKey)
		if err != nil {
			return nil, errors.Wrap(err, "error while authenticating admin")
		}

		ctx, err = uc.SetRequestContext(ctx, authType, nil)
		if err != nil {
			return nil, errors.Wrap(err, "error while authenticating admin")
		}

		return ctx, nil
	default:
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid authentication type")
	}
}
