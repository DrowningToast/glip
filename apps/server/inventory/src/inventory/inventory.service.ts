import { PrismaClient, ShipmentStatus } from "@prisma/client";
import { Prisma } from "@prisma/client";
import { notifyWarehouse } from "../rabbitmq/publisher";
import { InventoryRabbitMQType, ShipmentStatusRabbitMQType } from "../rabbitmq/inventoryType";

export class InventoryService {
  constructor(private prisma: PrismaClient) {}

  async getAllInventory() {
    return (await this.prisma.shipments.findMany()) || [];
  }

  async getInventoryById(id: string) {
    return await this.prisma.shipments.findUnique({
      where: {
        id,
      },
    });
  }

  async createInventory(data: Prisma.ShipmentsCreateInput) {
    return await this.prisma.shipments.create({
      data: {
        ...data,
        warehouse_id: Bun.env.INVENTORY_REGION as string,
      },
    });
  }

  async updateInventoryById(
    id: string,
    data: { name: string; weight: number; remarks?: string }
  ) {
    return await this.prisma.shipments.update({
      where: {
        id,
      },
      data,
    });
  }

  async updateInventoryStatus(id: string, status: ShipmentStatus) {
    if (status === ShipmentStatus.INCOMING_SHIPMENT) {
        await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          status: status,
          delivery_time: null,
        },
      });
    } else if (status === ShipmentStatus.WAREHOUSE_RECEIVED) {
      const shipment = await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          status: status,
          delivery_time: new Date(),
        },
      });
      console.log(shipment.route)
      await notifyWarehouse({
        id: shipment.shipmentId,
        route: shipment.route,
        last_warehouse_id: Bun.env.INVENTORY_REGION as string,
        departure_warehouse_id: shipment.departure_warehouse_id,
        departure_address: shipment.departure_address,
        destination_warehouse_id: shipment.destination_warehouse_id,
        destination_address: shipment.destination_address,
        created_by: shipment.created_by,
        owner_id: shipment.owner_id,
        status: "ARRIVED_AT_WAREHOUSE",
        total_weight: shipment.total_weight,
        total_volume: shipment.total_volume,
        created_at: shipment.created_at.toISOString(),
        updated_at: shipment.updated_at.toISOString(),
        from_warehouse_id: Bun.env.INVENTORY_REGION as string,
        type: "INBOUND",
        special_instructions: shipment.special_instructions,
      });
    } else if (status === ShipmentStatus.WAREHOUSE_DEPARTED) {
      const shipment = await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          status: status,
        },
      });
      console.log(shipment.route)
      console.log('IN TRANSIT ON THE WAY')
      await notifyWarehouse({
        id: shipment.shipmentId,
        route: shipment.route,
        last_warehouse_id: Bun.env.INVENTORY_REGION as string,
        departure_warehouse_id: shipment.departure_warehouse_id,
        departure_address: shipment.departure_address,
        destination_warehouse_id: shipment.destination_warehouse_id,
        destination_address: shipment.destination_address,
        created_by: shipment.created_by,
        owner_id: shipment.owner_id,
        status: "IN_TRANSIT_ON_THE_WAY",
        total_weight: shipment.total_weight,
        total_volume: shipment.total_volume,
        created_at: shipment.created_at.toISOString(),
        updated_at: shipment.updated_at.toISOString(),
        from_warehouse_id: Bun.env.INVENTORY_REGION as string,
        type: "INBOUND",
        special_instructions: shipment.special_instructions,
      });
    } else if (status === ShipmentStatus.DELIVERED) {
      const shipment = await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          status: status,
        },
      });
      console.log(shipment.route)
      console.log('DELIVERED')
      await notifyWarehouse({
        id: shipment.shipmentId,
        route: shipment.route,
        last_warehouse_id: Bun.env.INVENTORY_REGION as string,
        departure_warehouse_id: shipment.departure_warehouse_id,
        departure_address: shipment.departure_address,
        destination_warehouse_id: shipment.destination_warehouse_id,
        destination_address: shipment.destination_address,
        created_by: shipment.created_by,
        owner_id: shipment.owner_id,
        status: "DELIVERED",
        total_weight: shipment.total_weight,
        total_volume: shipment.total_volume,
        created_at: shipment.created_at.toISOString(),
        updated_at: shipment.updated_at.toISOString(),
        from_warehouse_id: Bun.env.INVENTORY_REGION as string,
        type: "INBOUND",
        special_instructions: shipment.special_instructions,
      });
    } else {
      await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          status: status,
        },
      });
    }
  }

  async removeInventoryById(id: string) {
    return await this.prisma.shipments.delete({
      where: {
        id,
      },
    });
  }
}
