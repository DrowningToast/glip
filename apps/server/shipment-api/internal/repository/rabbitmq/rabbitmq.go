package rabbitmq

import (
	"fmt"
	"log"

	"github.com/drowningtoast/glip/apps/server/internal/services"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQRepository struct {
	Warehouses *config.WarehouseRegions
	Config     *services.RabbitMQConfig

	Channel      *amqp091.Channel
	InboundQueue *amqp091.Queue
}

func NewRepository(warehouses *config.WarehouseRegions, config *services.RabbitMQConfig, channel *amqp091.Channel) *RabbitMQRepository {
	queueName := "warehouse_queue/to_shipment_api"

	inboundQueue, err := channel.QueueDeclare(queueName, true, true, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	return &RabbitMQRepository{
		Warehouses:   warehouses,
		Config:       config,
		Channel:      channel,
		InboundQueue: &inboundQueue,
	}
}

func (r *RabbitMQRepository) GetWarehouseOutboundQueue(warehouseId string) *amqp091.Queue {
	queueName := fmt.Sprintf("warehouse_queue/to_warehouse/%s", warehouseId)

	q, err := r.Channel.QueueDeclare(queueName, true, true, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	log.Printf("declared queue: %v", q)

	return &q
}
