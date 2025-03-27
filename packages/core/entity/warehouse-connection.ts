import { z } from "zod";

export const WarehouseConnectionStatusSchema = z.enum([
	"ACTIVE",
	"INACTIVE",
	"REVOKED",
]);

export type WarehouseConnectionStatus = z.infer<
	typeof WarehouseConnectionStatusSchema
>;

export const WarehouseConnectionSchema = z.object({
	id: z.number().int(),
	warehouse_id: z.string(),
	api_key: z.string(),
	name: z.string(),
	status: WarehouseConnectionStatusSchema,
	created_at: z.string().optional(),
	updated_at: z.string().optional(),
	last_used_at: z.string().optional(),
});

export type WarehouseConnection = z.infer<typeof WarehouseConnectionSchema>;
