import { z } from "zod";

export const ShipmentStatusSchema = z.enum([
	"WAITING_FOR_PICKUP_TO_WAREHOUSE",
	"ARRIVED_AT_WAREHOUSE",
	"IN_TRANSIT_ON_THE_WAY",
	"DELIVERED",
	"CANCELLED",
	"LOST",
]);

export type ShipmentStatus = z.infer<typeof ShipmentStatusSchema>;

export const ShipmentSchema = z.object({
	id: z.number(),
	route: z.array(z.string()),
	last_warehouse_id: z.string().nullable(),
	departure_warehouse_id: z.string(),
	departure_address: z.string().nullable(),
	destination_warehouse_id: z.string(),
	destination_address: z.string(),
	created_by: z.number(),
	owner_id: z.number().nullable(),
	status: ShipmentStatusSchema,
	total_weight: z.number(),
	total_volume: z.number(),
	special_instructions: z.string().nullable(),
	created_at: z.string(),
	updated_at: z.string(),
});

export type Shipment = z.infer<typeof ShipmentSchema>;
