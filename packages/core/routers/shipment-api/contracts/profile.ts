import { initContract } from "@ts-rest/core";
import { z } from "zod";
import {
	HTTPErrorResponseSchema,
	HTTPSuccessResponseSchema,
} from "../../../common";
import { AccountSchema } from "../../../entity/account";
import { CustomerSchema } from "../../../entity/customer";
import { WarehouseConnectionSchema } from "../../../entity/warehouse-connection";

const c = initContract();

export const ProfileContract = c.router(
	{
		getMyProfileAsCustomer: {
			method: "GET",
			path: "/customer/me",
			headers: z.object({
				Authorization: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						account: AccountSchema,
						customer: CustomerSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		getMyProfileAsWarehouseConnection: {
			method: "GET",
			path: "/warehouse-connection/me",
			headers: z.object({
				Authorization: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						account: AccountSchema,
						warehouseConnection: WarehouseConnectionSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
	},
	{
		pathPrefix: "/v1/profile",
	}
);
