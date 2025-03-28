package services

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Host     string `json:"host" env:"HOST"`
	Port     string `json:"port" env:"PORT"`
	User     string `json:"user" env:"USER"`
	Password string `json:"password" env:"PASSWORD"`
	DBName   string `json:"db_name" env:"DB_NAME"`
}

func (c *PostgresConfig) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}

func (c *PostgresConfig) NewConnection(ctx context.Context) (*pgxpool.Pool, error) {
	connString := c.String()
	// connConfig, err := pgx.ParseConfig(connString)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to parse connection string")
	// }

	dbpool, err := pgxpool.New(ctx, connString)
	// conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	return dbpool, nil
}
