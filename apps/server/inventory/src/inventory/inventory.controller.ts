import { Elysia, t } from "elysia";
import { ShipmentStatus } from "@prisma/client";

import { prismaClient } from "../libs/prismaClient";
import { InventoryService } from "./inventory.service";
import { inventoryCreateDto, inventoryUpdateDto } from "./dto/inventory.dto";

export const inventoryController = new Elysia({
  prefix: "/inventory",
  detail: {
    tags: ["Inventory"],
  },
})
  .decorate("prisma", prismaClient)
  .decorate("inventoryService", new InventoryService(prismaClient))
  
inventoryController.get("/", async ({ inventoryService }) => {
    return inventoryService.getAllInventory();
  })
  
inventoryController.get(
    "/:id",
    async ({ inventoryService, params, error }) => {
      const parcel = await inventoryService.getInventoryById(params.id);
      if (!parcel) {
        return error(404, { message: "Inventory not found" });
      }
      return parcel;
    },
    {
      params: t.Object({
        id: t.String(),
      }),
    }
  )

inventoryController.post(
    "/",
    async ({ inventoryService, body }) => {
      return inventoryService.createInventory(body);
    },
    inventoryCreateDto
  )

inventoryController.put(
    "/:id",
    async ({ inventoryService, body, params, error }) => {
      const parcelExists = await inventoryService.getInventoryById(params.id);
      if (!parcelExists) {
        return error(404, { message: "Inventory not found" });
      }
      return inventoryService.updateInventoryById(params.id, body);
    },
    inventoryUpdateDto
  )

inventoryController.put(
    "/:id/status",
    async ({ inventoryService, params, body, error }) => {
      const parcelExists = await inventoryService.getInventoryById(params.id);
      if (!parcelExists) {
        return error(404, { message: "Inventory not found" });
      }
      return inventoryService.updateInventoryStatus(params.id, body.status);
    },
    {
      params: t.Object({
        id: t.String(),
      }),
      body: t.Object({
        status: t.Enum(ShipmentStatus),
      }),
    }
  )

inventoryController.delete(
    "/:id",
    async ({ inventoryService, params, error }) => {
      const parcelExists = await inventoryService.getInventoryById(params.id);
      if (!parcelExists) {
        return error(404, { message: "Inventory not found" });
      }
      return inventoryService.removeInventoryById(params.id);
    },
    {
      params: t.Object({
        id: t.String(),
      }),
    }
  );
