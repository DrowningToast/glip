import amqp from "amqplib";
import { InventoryService } from "../inventory/inventory.service";
import { prismaClient } from "../libs/prismaClient";
import { InventoryRabbitMQType } from "./inventoryType";
import { ShipmentStatus } from "@prisma/client";
const rabbitMQUrl = process.env.RABBITMQ_URL;

if (!rabbitMQUrl) {
    throw new Error("RABBITMQ_URL is not set");
}

const connection = await amqp.connect(rabbitMQUrl);
const channel = await connection.createChannel();

const inventoryService = new InventoryService(prismaClient);

await channel.assertQueue(`warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`);

console.log(`Waiting for messages in warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`);

channel.consume(`warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`, (msg) => {
    if (msg) {
        const shipment: InventoryRabbitMQType = JSON.parse(msg.content.toString());
        console.log(`Received message type ${shipment.QueueType}`);
        if (shipment.QueueType === "OUTBOUND") {
            inventoryService.createInventory({
                shipmentId: shipment.shipmentId,
                Route: shipment.Route,
                LastWarehouseId: shipment.LastWarehouseId,
                DepartureWarehouseId: shipment.DepartureWarehouseId,
                DepartureAddress: shipment.DepartureAddress,
                DestinationWarehouseId: shipment.DestinationWarehouseId,
                DestinationAddress: shipment.DestinationAddress,
                CreatedBy: shipment.CreatedBy,
                OwnerId: shipment.OwnerId,
                Status: ShipmentStatus.INCOMING_SHIPMENT,
                TotalWeight: shipment.TotalWeight,
                TotalVolume: shipment.TotalVolume,
                CreatedAt: shipment.CreatedAt,
                UpdatedAt: shipment.UpdatedAt,
                FromWarehouseId: shipment.FromWarehouseId,
                ToWarehouseId: shipment.ToWarehouseId,
            });
            console.log(`Created inventory for shipment ${shipment.shipmentId}`);
        }
    }
});

