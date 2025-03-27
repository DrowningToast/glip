import { z } from "zod";

export const HTTPSuccessResponseSchema = <T extends z.ZodType>(schema: T) =>
	z.object({
		result: schema,
	});

export const HTTPErrorResponseSchema = <
	codes extends [string, ...string[]],
	T extends z.ZodEnum<codes>,
>(
	codes?: T
) =>
	z.object({
		code: codes ? codes : z.string(),
		message: z.string(),
	});

export const PaginatedResultSchema = <T extends z.ZodType>(schema: T) =>
	z.object({
		count: z.number(),
		items: z.array(schema),
		offset: z.number(),
		limit: z.number(),
	});
