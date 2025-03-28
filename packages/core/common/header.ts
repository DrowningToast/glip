import { z } from "zod";

export const HeaderBearerSchema = z.object({
	authorization: z.string().startsWith("Bearer "),
});

export const RegistryAuthHeadersSchema = z.object({
	Authorization: z.string(),
	AuthType: z.string(),
});
