import { z } from "zod";

export const HeaderBearerSchema = z.object({
	authorization: z.string().startsWith("Bearer "),
});

export const HeaderAuthorizationSchema = z.object({
	Authorization: z.string().startsWith("Bearer "),
});

export const HeaderAuthTypeSchema = z.object({
	AuthType: z.string(),
});
