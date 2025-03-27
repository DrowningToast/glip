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

await channel.assertQueue(`warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`, {
    durable: true,
    autoDelete: true
});

console.log(`Waiting for messages in warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`);

channel.consume(`warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`, (msg) => {
    if (msg) {
        const shipment: InventoryRabbitMQType = JSON.parse(msg.content.toString());
        console.log(`Received message type ${shipment.type}`);
        console.log(shipment);
        if (shipment.type === "OUTBOUND") {
            inventoryService.createInventory({
                shipmentId: shipment.id,
                route: shipment.route,
                last_warehouse_id: shipment.last_warehouse_id,
                departure_warehouse_id: shipment.departure_warehouse_id,
                departure_address: shipment.departure_address,
                destination_warehouse_id: shipment.destination_warehouse_id,
                destination_address: shipment.destination_address,
                created_by: shipment.created_by,
                owner_id: shipment.owner_id,
                status: ShipmentStatus.INCOMING_SHIPMENT,
                total_weight: Number(shipment.total_weight),
                total_volume: Number(shipment.total_volume),
                created_at: shipment.created_at,
                updated_at: shipment.updated_at,
                from_warehouse_id: shipment.from_warehouse_id,
                to_warehouse_id: shipment.to_warehouse_id,
            });
            console.log(`Created inventory for shipment ${shipment.id}`);
        }
    }
});

