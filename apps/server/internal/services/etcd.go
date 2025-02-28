package services

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfig struct {
	Endpoint     string `env:"ETCD_ENDPOINT,required"`
	Username     string `env:"ETCD_USERNAME,required"`
	RootPassword string `env:"ETCD_ROOT_PASSWORD,required"`
	Port         string `env:"ETCD_PORT,required"`
}

func (c *EtcdConfig) NewConnection() (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", c.Endpoint, c.Port)},
		DialTimeout: 2 * time.Second,
	})

	if err == context.DeadlineExceeded {
		return nil, errors.Wrap(err, "etcd connection timeout")
	}

	return client, nil
}
