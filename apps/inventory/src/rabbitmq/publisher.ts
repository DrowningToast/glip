import amqp from "amqplib";
import { InventoryRabbitMQType } from "./inventoryType";

export const notifyWarehouse = async (shipment: InventoryRabbitMQType) => {
    try {
        const connection = await amqp.connect(process.env.RABBITMQ_URL as string);
        const channel = await connection.createChannel();
        
        await channel.assertQueue(`warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`);
        channel.sendToQueue(`warehouse_queue/to_warehouse/${process.env.INVENTORY_REGION}`, Buffer.from(JSON.stringify(shipment)));

        setTimeout(async () => {
            await connection.close();
        }, 1000);
    } catch (error) {
        console.error(error);
    }
}