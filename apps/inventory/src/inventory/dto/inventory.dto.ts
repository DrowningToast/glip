import { t } from "elysia";
import { ShipmentStatus } from "@prisma/client";

export const inventoryCreateDto = {
  body: t.Object({
    shipmentId: t.String(),
    Route: t.Array(t.String()),
    LastWarehouseId: t.String(),
    DepartureWarehouseId: t.String(),
    DepartureAddress: t.String(),
    DestinationWarehouseId: t.String(),
    DestinationAddress: t.String(),
    CreatedBy: t.Number(),
    OwnerId: t.Number(),
    Status: t.Enum(ShipmentStatus),
    TotalWeight: t.Number(),
    TotalVolume: t.Number(),
    SpecialInstructions: t.String(),
    CreatedAt: t.String(),
    UpdatedAt: t.String(),
    FromWarehouseId: t.String(),
    ToWarehouseId: t.String(),
    DeliveryTime: t.String(),
  }),
};

export const inventoryUpdateDto = {
    params: t.Object({
        id: t.String(),
    }),
    body: t.Object({
        name: t.String(),
        weight: t.Number(),
        remarks: t.String({ optional: true }),
    })
}