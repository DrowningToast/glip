import { initContract } from "@ts-rest/core";
import { z } from "zod";
import {
	HTTPErrorResponseSchema,
	HTTPSuccessResponseSchema,
	PaginatedResultSchema,
} from "../../../common/http";
import {
	WarehouseConnectionSchema,
	WarehouseConnectionStatusSchema,
} from "../../../entity/warehouse-connection";
import { HeaderBearerSchema } from "../../../common/header";

const c = initContract();

export const WarehouseConnectionContract = c.router(
	{
		getConnection: {
			method: "GET",
			path: "/",
			query: z.object({
				id: z.string().optional(),
				"api-key": z.string().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						warehouse_connection: WarehouseConnectionSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
		listConnections: {
			method: "GET",
			path: "/list",
			query: z.object({
				limit: z.number().optional(),
				offset: z.number().optional(),

				status: WarehouseConnectionStatusSchema.optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					PaginatedResultSchema(WarehouseConnectionSchema)
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
		createConnection: {
			method: "POST",
			path: "/",
			body: z.object({
				warehouse_connection: z.object({
					warehouse_id: z.string(),
					api_key: z.string(),
					status: WarehouseConnectionStatusSchema,
					name: z.string(),
				}),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						warehjouse_connection: WarehouseConnectionSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
		updateConnection: {
			method: "PUT",
			path: "/",
			body: z.object({
				warehouse_connection: z.object({
					id: z.string(),
					api: z.string().optional(),
					status: WarehouseConnectionStatusSchema.optional(),
					name: z.string().optional(),
					warehouse_id: z.string(),
				}),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						warehouse_connection: WarehouseConnectionSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
		deleteConnection: {
			method: "DELETE",
			path: "/:id",
			responses: {
				200: HTTPSuccessResponseSchema(z.object({})),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
	},
	{
		pathPrefix: "/v1/warehouse-connection",
	}
);
