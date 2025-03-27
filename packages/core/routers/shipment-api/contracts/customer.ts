import { initContract } from "@ts-rest/core";
import { z } from "zod";
import {
	HTTPErrorResponseSchema,
	HTTPSuccessResponseSchema,
	PaginatedResultSchema,
} from "../../../common/http";
import { CustomerSchema } from "../../../entity/customer";
import { AccountSchema } from "../../../entity/account";

const c = initContract();

export const CustomerContract = c.router(
	{
		getCustomer: {
			method: "GET",
			path: "/",
			query: z.object({
				id: z.string().optional(),
				email: z.string().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						customer: CustomerSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		listCustomers: {
			method: "GET",
			path: "/list",
			query: z.object({
				limit: z.number().optional(),
				offset: z.number().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(PaginatedResultSchema(CustomerSchema)),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		createCustomer: {
			method: "POST",
			path: "/",
			body: z.object({
				username: z.string(),
				password: z.string(),
				email: z.string(),
				name: z.string(),
				phone: z.string().optional(),
				address: z.string().optional(),
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
		updateCustomer: {
			method: "PUT",
			path: "/",
			body: z.object({
				id: z.string(),
				email: z.string().optional(),
				name: z.string().optional(),
				phone: z.string().optional(),
				address: z.string().optional(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						customer: CustomerSchema,
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		deleteCustomer: {
			method: "DELETE",
			path: "/:id",
			pathParams: z.object({
				id: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(z.object({})),
			},
		},
	},
	{
		pathPrefix: "/v1/customer",
	}
);
