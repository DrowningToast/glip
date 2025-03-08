package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/registry-api/internal/entity"
	"github.com/samber/lo"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var _ datagateway.WarehouseEndpointDataGateway = (*EtcdRepository)(nil)

func (r *EtcdRepository) GetEndpointByWarehouseId(ctx context.Context, warehouseId string) (*entity.WarehouseEndpoint, error) {
	resp, err := r.Client.Get(ctx, fmt.Sprintf("warehouse/%s", warehouseId))
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to get endpoint by warehouse id")
	}

	if resp.Count == 0 {
		return nil, nil
	}

	value := string(resp.Kvs[0].Value)

	return &entity.WarehouseEndpoint{
		WarehouseId: warehouseId,
		Endpoint:    value,
		UpdatedAt:   time.Unix(resp.Kvs[0].ModRevision, 0),
	}, nil
}

func (r *EtcdRepository) ListEndpoints(ctx context.Context) ([]*entity.WarehouseEndpoint, error) {
	resp, err := r.Client.Get(ctx, "warehouse/", clientv3.WithPrefix())
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

	warehouseEndpoints := lo.Map(resp.Kvs, func(kv *mvccpb.KeyValue, _ int) *entity.WarehouseEndpoint {
		return &entity.WarehouseEndpoint{
			WarehouseId: string(kv.Key),
			Endpoint:    string(kv.Value),
			UpdatedAt:   time.Unix(kv.ModRevision, 0),
		}
	})

	return warehouseEndpoints, nil
}

func (r *EtcdRepository) UpdateEndpoint(ctx context.Context, warehouseId string, endpoint string) error {
	_, err := r.Client.Put(ctx, fmt.Sprintf("warehouse/%s", warehouseId), endpoint)
	if err != nil {
		return errors.Wrap(errs.ErrInternal, "failed to update endpoint")
	}

	return nil
}

func (r *EtcdRepository) DeleteEndpoint(ctx context.Context, warehouseId string) error {
	_, err := r.Client.Delete(ctx, fmt.Sprintf("warehouse/%s", warehouseId))
	if err != nil {
		return errors.Wrap(errs.ErrInternal, "failed to delete endpoint")
	}

	return nil
}
