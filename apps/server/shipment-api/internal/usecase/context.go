package usecase

import (
	"context"

	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
)

type UserContextKey struct{}

func (u *Usecase) InitUserContext(ctx context.Context, session *entity.JWTSession) context.Context {
	return context.WithValue(ctx, UserContextKey{}, session)
}

func (u *Usecase) GetUserContext(ctx context.Context) *entity.JWTSession {
	session, ok := ctx.Value(UserContextKey{}).(*entity.JWTSession)
	if !ok {
		return nil
	}

	return session
}
