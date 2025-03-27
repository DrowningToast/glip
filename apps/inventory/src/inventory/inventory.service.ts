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
      data,
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
          Status: status,
          DeliveryTime: null,
        },
      });
    } else if (status === ShipmentStatus.WAREHOUSE_RECEIVED) {
      const shipment = await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          Status: status,
          DeliveryTime: new Date(),
        },
      });
      await notifyWarehouse({
        shipmentId: shipment.shipmentId,
        Route: shipment.Route,
        LastWarehouseId: shipment.LastWarehouseId,
        DepartureWarehouseId: shipment.DepartureWarehouseId,
        DepartureAddress: shipment.DepartureAddress,
        DestinationWarehouseId: shipment.DestinationWarehouseId,
        DestinationAddress: shipment.DestinationAddress,
        CreatedBy: shipment.CreatedBy,
        OwnerId: shipment.OwnerId,
        Status: "ARRIVED_AT_WAREHOUSE",
        TotalWeight: shipment.TotalWeight,
        TotalVolume: shipment.TotalVolume,
        CreatedAt: shipment.CreatedAt.toISOString(),
        UpdatedAt: shipment.UpdatedAt.toISOString(),
        FromWarehouseId: shipment.FromWarehouseId,
        ToWarehouseId: shipment.ToWarehouseId,
        QueueType: "INBOUND",
        SpecialInstructions: shipment.SpecialInstructions,
      });
    } else if (status === ShipmentStatus.WAREHOUSE_DEPARTED) {
      const shipment = await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          Status: status,
        },
      });
      await notifyWarehouse({
        shipmentId: shipment.shipmentId,
        Route: shipment.Route,
        LastWarehouseId: shipment.LastWarehouseId,
        DepartureWarehouseId: shipment.DepartureWarehouseId,
        DepartureAddress: shipment.DepartureAddress,
        DestinationWarehouseId: shipment.DestinationWarehouseId,
        DestinationAddress: shipment.DestinationAddress,
        CreatedBy: shipment.CreatedBy,
        OwnerId: shipment.OwnerId,
        Status: "IN_TRANSIT_ON_THE_WAY",
        TotalWeight: shipment.TotalWeight,
        TotalVolume: shipment.TotalVolume,
        CreatedAt: shipment.CreatedAt.toISOString(),
        UpdatedAt: shipment.UpdatedAt.toISOString(),
        FromWarehouseId: shipment.FromWarehouseId,
        ToWarehouseId: shipment.ToWarehouseId,
        QueueType: "INBOUND",
        SpecialInstructions: shipment.SpecialInstructions,
      });
    } else {
      await this.prisma.shipments.update({
        where: {
          id,
        },
        data: {
          Status: status,
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
