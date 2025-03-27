package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) SignJWT(ctx context.Context, identifier string, role entity.ConnectionType) (string, error) {
	claims := entity.JWTSession{
		Id:   identifier,
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

type warehouseConnectionStatus string

const (
	warehouseConnectionStatusActive   warehouseConnectionStatus = "ACTIVE"
	warehouseConnectionStatusInactive warehouseConnectionStatus = "INACTIVE"
	warehouseConnectionStatusRevoked  warehouseConnectionStatus = "REVOKED"
)

type WarehouseConnectionResponse struct {
	WarehouseConnection struct {
		Id string `json:"id"`
		// The warehouse id that the connection is for
		WarehouseId string                    `json:"warehouse_id"`
		ApiKey      string                    `json:"api_key"`
		Name        string                    `json:"name"`
		Status      warehouseConnectionStatus `json:"status"`
		CreatedAt   *time.Time                `json:"created_at"`
		UpdatedAt   *time.Time                `json:"updated_at"`
		LastUsedAt  *time.Time                `json:"last_used_at"`
	} `json:"warehouse_connection"`
}

func (u *Usecase) CreateUserConnectionSession(ctx context.Context, username string, password string) (*string, error) {
	account, err := u.AccountDg.GetAccountByUsername(ctx, username)
	if err != nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid username or password")
	}
	if account == nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid username or password")
	}

	tokenString, err := u.SignJWT(ctx, account.Username, entity.ConnectionTypeUser)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	return &tokenString, nil
}

// return jwt token string
func (u *Usecase) CreateWarehouseConnectionSession(ctx context.Context, apiKey string) (*string, error) {
	warehouseConn, err := u.WarehouseConnDg.GetWarehouseConnectByApiKey(ctx, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "cannot verify warehouse connection")
	}

	if warehouseConn.Status == entity.WarehouseConnectionStatusRevoked {
		return nil, errors.Wrap(errs.ErrUnauthorized, "warehouse connection revoked")
	}

	if warehouseConn.Status == entity.WarehouseConnectionStatusInactive {
		return nil, errors.Wrap(errs.ErrUnauthorized, "warehouse connection inactive")
	}

	tokenString, err := u.SignJWT(ctx, warehouseConn.Id, entity.ConnectionTypeWarehouse)
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

	tokenString, err := u.SignJWT(ctx, "root", entity.ConnectionTypeRoot)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	return &tokenString, nil
}
