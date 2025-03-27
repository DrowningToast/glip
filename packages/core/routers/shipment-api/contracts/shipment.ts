import { initContract } from "@ts-rest/core";
import { z } from "zod";
import { ShipmentSchema, ShipmentStatusSchema } from "../../../entity/shipment";
import {
	HTTPErrorResponseSchema,
	HTTPSuccessResponseSchema,
	PaginatedResultSchema,
} from "../../../common";

const c = initContract();

export const ShipmentApiContract = c.router(
	{
		listByCustomer: {
			method: "GET",
			path: "/customer/list",
			query: z.object({
				limit: z.number().optional(),
				offset: z.number().optional(),
				status: ShipmentStatusSchema.optional(),
				username: z.string().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(PaginatedResultSchema(ShipmentSchema)),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		track: {
			method: "POST",
			path: "/track",
			body: z.object({
				shipment_id: z.number(),
				email: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						shipment: ShipmentSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		create: {
			method: "POST",
			path: "/",
			body: z.object({
				departure_address: z.string(),
				departure_city: z.string(),
				destination_address: z.string(),
				destination_city: z.string(),
				owner_id: z.number().optional(),
				total_weight: z.number(),
				total_volume: z.number(),
				special_instructions: z.string().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						shipment: ShipmentSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		list: {
			method: "GET",
			path: "/list",
			query: z.object({
				limit: z.number().optional(),
				offset: z.number().optional(),
				status: ShipmentStatusSchema.optional(),
				last_warehouse_id: z.string().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(PaginatedResultSchema(ShipmentSchema)),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		get: {
			method: "GET",
			path: "/:shipment_id",
			pathParams: z.object({
				shipment_id: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						shipment: ShipmentSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
	},
	{
		pathPrefix: "/shipment",
	}
);
