import { z } from "zod";

export const HeaderBearerSchema = z.object({
	authorization: z.string().startsWith("Bearer "),
});

export const RegistryAuthHeadersSchema = z.object({
	authorization: z.string(),
	authtype: z.string(),
});
