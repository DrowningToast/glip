package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

func (u *Usecase) SignJWT(ctx context.Context, id string, role entity.ConnectionType) (string, error) {
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

type warehouseConnectionStatus string

const (
	warehouseConnectionStatusActive   warehouseConnectionStatus = "ACTIVE"
	warehouseConnectionStatusInactive warehouseConnectionStatus = "INACTIVE"
	warehouseConnectionStatusRevoked  warehouseConnectionStatus = "REVOKED"
)

type warehouseConnectionResponse struct {
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

// return jwt token string
func (u *Usecase) CreateWarehouseConnectionSession(ctx context.Context, apiKey string) (*string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s/v1/warehouse-connection", u.Config.RegistryEndpoint, u.Config.RegistryPort), nil)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}
	req.Header.Set("Authorization", u.Config.RegistryApiKey)
	req.Header.Set("AuthType", "ADMIN")
	q := req.URL.Query()
	q.Add("api-key", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	defer resp.Body.Close()

	// check if found or not
	if resp.StatusCode == http.StatusNotFound {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(errs.ErrInternal, err.Error())
		}
		return nil, errors.Wrap(errs.ErrUnauthorized, string(body))
	}
	if resp.StatusCode != http.StatusOK {
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(errs.ErrInternal, err.Error())
		}
		return nil, errors.Wrap(errs.ErrInternal, "error while querying the database")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	respBody := struct {
		Result warehouseConnectionResponse `json:"result"`
	}{}
	if err := json.Unmarshal(body, &respBody); err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	warehouseConn := respBody.Result.WarehouseConnection

	if warehouseConn.Status == warehouseConnectionStatusRevoked {
		return nil, errors.Wrap(errs.ErrUnauthorized, "warehouse connection revoked")
	}

	if warehouseConn.Status == warehouseConnectionStatusInactive {
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
