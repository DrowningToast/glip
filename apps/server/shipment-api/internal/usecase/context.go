package usecase

import (
	"context"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/gofiber/fiber/v3"
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

func (u *Usecase) InjectSessionContext(ctx context.Context, c fiber.Ctx) (*entity.JWTSession, error) {
	bearerString := c.Get("Authorization")
	if bearerString == "" {
		return nil, nil
	}

	splitedTokenString := strings.Split(bearerString, " ")
	if len(splitedTokenString) != 2 {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid authorization header")
	}

	tokenString := splitedTokenString[1]

	session, err := u.VerifyJWT(c.Context(), tokenString)
	if err != nil {
		return nil, errors.Wrap(errs.ErrUnauthorized, err.Error())
	}

	if !session.Role.Valid() {
		return nil, errors.Wrap(errs.ErrUnauthorized, "invalid role")
	}

	context := u.InitUserContext(c.Context(), session)
	c.SetContext(context)

	return session, nil
}
