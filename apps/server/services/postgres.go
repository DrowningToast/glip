package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PostgresConfig struct {
	Host     string `json:"host" env:"HOST"`
	Port     string `json:"port" env:"PORT"`
	User     string `json:"user" env:"USER"`
	Password string `json:"password" env:"PASSWORD"`
	DBName   string `json:"db_name" env:"DB_NAME"`
}

func (c *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBName)
}

func (c *PostgresConfig) NewConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, c.GetConnectionString())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
