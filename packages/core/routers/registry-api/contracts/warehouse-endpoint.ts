import { initContract } from "@ts-rest/core";
import { z } from "zod";
import {
	HTTPErrorResponseSchema,
	HTTPSuccessResponseSchema,
} from "../../../common/http";
import { WarehouseEndpointSchema } from "../../../entity/warehouse-endpoints";
import { HeaderBearerSchema } from "../../../common/header";

const c = initContract();

export const WarehouseEndpointContract = c.router(
	{
		getEndpoint: {
			method: "GET",
			path: "/",
			query: z.object({
				id: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						warehouse_endpoint: WarehouseEndpointSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
		listEndpoint: {
			method: "GET",
			path: "/list",
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						warehouses_endpoints: z.array(WarehouseEndpointSchema),
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			headers: HeaderBearerSchema,
		},
		updateEndpoint: {
			method: "PUT",
			path: "/",
			responses: {
				200: HTTPSuccessResponseSchema(z.object({})),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
			body: z.undefined(),
			headers: HeaderBearerSchema,
		},
	},
	{
		pathPrefix: "/v1/warehouse-endpoint",
	}
);
