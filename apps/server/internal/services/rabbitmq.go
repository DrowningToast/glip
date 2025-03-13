package services

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	DefaultUser  string `json:"default_user" env:"DEFAULT_USER"`
	DefaultPass  string `json:"default_pass" env:"DEFAULT_PASS"`
	DefaultHost  string `json:"default_host" env:"DEFAULT_HOST"`
	DefaultPort  string `json:"default_port" env:"DEFAULT_PORT"`
	DefaultVhost string `json:"default_vhost" env:"DEFAULT_VHOST"`
}

func (c *RabbitMQConfig) String() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.DefaultUser, c.DefaultPass, c.DefaultHost, c.DefaultPort, c.DefaultVhost)
}

func (c *RabbitMQConfig) NewConnection(ctx context.Context) (*amqp.Connection, error) {
	conn, err := amqp.Dial(c.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to rabbitmq")
	}

	return conn, nil
}

func NewRabbmitMQChannel(ctx context.Context, conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open a channel")
	}

	return ch, nil
}
