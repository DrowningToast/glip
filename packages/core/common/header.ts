import { z } from "zod";

export const HeaderBearerSchema = z.object({
	authorization: z.string().startsWith("Bearer "),
});
