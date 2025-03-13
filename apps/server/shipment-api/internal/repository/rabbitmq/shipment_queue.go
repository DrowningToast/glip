package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/datagateway"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
)

var _ datagateway.ShipmentQueueDataGateway = (*RabbitMQRepository)(nil)

func (r *RabbitMQRepository) CreateToReceivedShipment(ctx context.Context, shipment *entity.Shipment, warehouseId string) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	queue := r.GetWarehouseOutboundQueue(warehouseId)
	if queue == nil {
		return errors.Wrap(errs.ErrInternal, "failed to get warehouse queue")
	}

	shipmentQueue := entity.ShipmentQueue{
		Shipment: *shipment,

		ToWarehouseId: &warehouseId,
		QueueType:     entity.ShipmentQueueTypeOutbound,
	}

	json, err := json.Marshal(shipmentQueue)
	if err != nil {
		return errors.Wrap(errs.ErrInternal, err.Error())
	}

	err = r.Channel.PublishWithContext(
		timeoutCtx,
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        json,
		},
	)
	if err != nil {
		return errors.Wrap(errs.ErrInternal, err.Error())
	}

	return nil
}

func (r *RabbitMQRepository) ListInboundShipments(ctx context.Context) (map[string][]entity.ShipmentQueue, error) {
	queue := r.InboundQueue
	if queue == nil {
		return nil, errors.Wrap(errs.ErrInternal, "failed to get warehouse queue")
	}

	queueByWarehouseId := make(map[string][]entity.ShipmentQueue)

	msgs, err := r.Channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrap(errs.ErrInternal, err.Error())
	}

messageLoop:
	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				break messageLoop
			}

			defer msg.Nack(false, true)

			var shipmentQueue entity.ShipmentQueue
			if err := json.Unmarshal(msg.Body, &shipmentQueue); err != nil {
				return nil, errors.Wrap(errs.ErrInternal, err.Error())
			}

			if shipmentQueue.FromWarehouseId == nil {
				return nil, errors.Wrap(errs.ErrInternal, "from warehouse id is nil")
			}

			queueByWarehouseId[*shipmentQueue.FromWarehouseId] = append(queueByWarehouseId[*shipmentQueue.FromWarehouseId], shipmentQueue)
		}
	}

	return queueByWarehouseId, nil
}

func (r *RabbitMQRepository) ListOutboundShipments(ctx context.Context) (map[string][]entity.ShipmentQueue, error) {
	warehouseMap := r.Warehouses.ToMap()
	warehouseQueueNames := lo.Map(lo.Values(warehouseMap), func(warehouse *entity.Warehouse, _ int) amqp091.Queue {
		return *r.GetWarehouseOutboundQueue(warehouse.Id)
	})

	queueByWarehouseId := make(map[string][]entity.ShipmentQueue)

	for _, queueName := range warehouseQueueNames {
		queueByWarehouseId[queueName.Name] = make([]entity.ShipmentQueue, 0)

		queue, err := r.Channel.QueueDeclare(queueName.Name, true, true, false, false, nil)
		if err != nil {
			// Queue does not exist, skip
			continue
		}

		msgs, err := r.Channel.Consume(queue.Name, "", false, false, false, false, nil)
		if err != nil {
			return nil, errors.Wrap(errs.ErrInternal, err.Error())
		}

	messageLoop:
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					break messageLoop
				}

				defer msg.Nack(false, true)

				var shipmentQueue entity.ShipmentQueue
				if err := json.Unmarshal(msg.Body, &shipmentQueue); err != nil {
					return nil, errors.Wrap(errs.ErrInternal, err.Error())
				}

				if shipmentQueue.ToWarehouseId == nil {
					return nil, errors.Wrap(errs.ErrInternal, "to warehouse id is nil")
				}

				queueByWarehouseId[*shipmentQueue.ToWarehouseId] = append(queueByWarehouseId[*shipmentQueue.ToWarehouseId], shipmentQueue)
			}
		}
	}

	return queueByWarehouseId, nil
}

func (r *RabbitMQRepository) WatchReceivedShipment(ctx context.Context, shipmentChan chan<- entity.ShipmentQueue, errorChan chan error, terminateChan <-chan struct{}) error {
	queue := r.InboundQueue
	if queue == nil {
		return errors.Wrap(errs.ErrInternal, "failed to get warehouse queue")
	}

	msgs, err := r.Channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return errors.Wrap(errs.ErrInternal, err.Error())
	}

	go func() {
		for msg := range msgs {
			var shipmentQueue entity.ShipmentQueue
			if err := json.Unmarshal(msg.Body, &shipmentQueue); err != nil {
				msg.Nack(false, true)
				errorChan <- errors.Wrap(errs.ErrInternal, err.Error())
			}
			shipmentQueue.Msg = &msg

			shipmentChan <- shipmentQueue
		}
	}()

	<-terminateChan
	close(shipmentChan)
	close(errorChan)

	return nil
}
