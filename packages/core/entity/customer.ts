import { z } from "zod";

export const CustomerSchema = z.object({
	id: z.string(),
	name: z.string(),
	email: z.string(),
	phone: z.string(),
	address: z.string(),
});

export type Customer = z.infer<typeof CustomerSchema>;
