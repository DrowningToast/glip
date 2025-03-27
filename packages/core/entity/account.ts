import { z } from "zod";

export const AccountSchema = z.object({
	id: z.string(),
	username: z.string(),
	password: z.string(),
	email: z.string(),
});

export type Account = z.infer<typeof AccountSchema>;
