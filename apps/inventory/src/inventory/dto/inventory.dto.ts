import { t } from "elysia";
import { ShipmentStatus } from "@prisma/client";

export const inventoryCreateDto = {
  body: t.Object({
    shipmentId: t.String(),
    route: t.Array(t.String()),
    last_warehouse_id: t.String(),
    departure_warehouse_id: t.String(),
    departure_address: t.String(),
    destination_warehouse_id: t.String(),
    destination_address: t.String(),
    created_by: t.Number(),
    owner_id: t.Number(),
    status: t.Enum(ShipmentStatus),
    total_weight: t.Number(),
    total_volume: t.Number(),
    created_at: t.String(),
    updated_at: t.String(),
    from_warehouse_id: t.String(),
    to_warehouse_id: t.String(),
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