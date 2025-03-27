import { initContract } from "@ts-rest/core";
import { z } from "zod";
import {
	HTTPErrorResponseSchema,
	HTTPSuccessResponseSchema,
} from "../../../common/http";

const c = initContract();

export const AuthContract = c.router(
	{
		AuthWarehouseConnection: {
			method: "POST",
			path: "/warehouse",
			body: z.object({
				key: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						jwt: z.string(),
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		AuthCustomerConnection: {
			method: "POST",
			path: "/admin",
			body: z.object({
				username: z.string(),
				password: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						jwt: z.string(),
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
		AuthAdminConnection: {
			method: "POST",
			path: "/customer",
			body: z.object({
				key: z.string(),
			}),
			responses: {
				200: HTTPSuccessResponseSchema(
					z.object({
						jwt: z.string(),
					})
				),
				400: HTTPErrorResponseSchema(),
				500: HTTPErrorResponseSchema(),
			},
		},
	},
	{
		pathPrefix: "/v1/auth",
	}
);
