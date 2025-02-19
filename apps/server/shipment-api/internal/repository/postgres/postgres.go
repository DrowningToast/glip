package postgres

import (
	"errors"

	database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresRepository struct {
	db      database.DBTX
	queries *database.Queries
}

func New(db database.DBTX) *PostgresRepository {
	return &PostgresRepository{
		db:      db,
		queries: database.New(db),
	}
}

func checkPgErrCode(err error, code string) bool {
	var pgErr *pgconn.PgError = &pgconn.PgError{}

	if errors.As(err, &pgErr) {
		return pgErr.Code == code
	}

	return false
}
