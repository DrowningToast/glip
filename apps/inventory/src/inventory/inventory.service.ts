import { PrismaClient } from "@prisma/client";
import { ParcelStatus } from "@prisma/client";

export class InventoryService {
  constructor(private prisma: PrismaClient) {}

  async getAllInventory() {
    return (await this.prisma.parcels.findMany()) || [];
  }

  async getInventoryById(id: string) {
    return await this.prisma.parcels.findUnique({
      where: {
        id,
      },
    });
  }

  async createInventory(data: {
    name: string;
    weight: number;
    remarks?: string;
  }) {
    return await this.prisma.parcels.create({
      data,
    });
  }

  async updateInventoryById(
    id: string,
    data: { name: string; weight: number; remarks?: string }
  ) {
    return await this.prisma.parcels.update({
      where: {
        id,
      },
      data,
    });
  }

  async updateInventoryStatus(id: string, status: ParcelStatus) {
    if (status === ParcelStatus.PENDING) {
      await this.prisma.parcels.update({
        where: {
          id,
        },
        data: {
          status,
          deliveryTime: null,
        },
      });
    } else if (status === ParcelStatus.DELIVERED) {
      await this.prisma.parcels.update({
        where: {
          id,
        },
        data: {
          status,
          deliveryTime: new Date(),
        },
      });
    } else {
      await this.prisma.parcels.update({
        where: {
          id,
        },
        data: {
          status,
        },
      });
    }
  }

  async removeInventoryById(id: string) {
    return await this.prisma.parcels.delete({
      where: {
        id,
      },
    });
  }
}
