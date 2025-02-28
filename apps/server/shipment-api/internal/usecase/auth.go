package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

func (u *Usecase) SignJWT(ctx context.Context, id int, role entity.ConnectionType) (string, error) {
	claims := entity.JWTSession{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(u.Config.ShipmentAuthConfig.JWTExpirationTime) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.Config.ShipmentAuthConfig.JWTSecret))
	if err != nil {
		return "", errors.Wrap(errs.ErrInternal, err.Error())
	}

	return tokenString, nil
}

func (u *Usecase) VerifyJWT(ctx context.Context, tokenString string) (*entity.JWTSession, error) {
	claims := entity.JWTSession{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.Config.ShipmentAuthConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, err.Error())
	}

	if !token.Valid {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid token")
	}

	return &claims, nil
}

// return jwt token string
func (u *Usecase) CreateWarehouseConnectionSession(ctx context.Context, apiKey string) (*string, error) {
	warehouse, err := u.WarehouseConnectionDg.GetWarehouseConnectionByApiKey(ctx, apiKey)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	if warehouse == nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid api key")
	}

	tokenString, err := u.SignJWT(ctx, warehouse.Id, entity.ConnectionTypeWarehouse)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	return &tokenString, nil
}

// return jwt token string
func (u *Usecase) CreateAdminApiSession(ctx context.Context, apiKey string) (*string, error) {
	configuredRootKeys := make([]string, 0)
	configuredRootKeys = append(configuredRootKeys, u.Config.ShipmentAuthConfig.JWTAdminApiSecret)

	if !lo.Contains(configuredRootKeys, apiKey) {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid api key")
	}

	tokenString, err := u.SignJWT(ctx, 0, entity.ConnectionTypeRoot)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	return &tokenString, nil
}
