import { z } from "zod";

export const WarehouseEndpointSchema = z.object({
	warehouse_id: z.string(),
	endpoint: z.string(),
	updated_at: z.coerce.date(),
});

export type WarehouseEndpoint = z.infer<typeof WarehouseEndpointSchema>;
