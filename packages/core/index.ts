import {
	initClient,
	type AppRouter,
	type ClientArgs,
	type InitClientReturn,
} from "@ts-rest/core";
import { isAxiosError, type AxiosInstance, type Method } from "axios";
import axiosClient from "../../apps/frontend/glip-warehouse/src/libs/axiosClient";

// Router
export * from "./routers/shipment-api/router";

// Contracts
export * from "./routers/shipment-api/contracts/customer";
export * from "./routers/shipment-api/contracts/auth";
export * from "./routers/shipment-api/contracts/shipment";
export * from "./routers/shipment-api/contracts/profile";

// Router
export * from "./routers/registry-api/router";

// Contracts
export * from "./routers/registry-api/contracts/warehouse-connection";
export * from "./routers/registry-api/contracts/warehouse-endpoint";

// Entity
export * from "./entity/index";

// Common
export * from "./common";

// Client
export type RawAPIClient<T extends AppRouter> = InitClientReturn<T, ClientArgs>;

export const NewAPIClient = <T extends AppRouter>(
	axiosClient: AxiosInstance,
	contract: T,
	baseUrl: string
): RawAPIClient<T> => {
	const client = initClient(contract, {
		baseUrl: baseUrl,
		baseHeaders: {},
		api: async (args) => {
			try {
				const result = await axiosClient.request({
					method: args.method as Method,
					url: encodeURI(args.path),
					headers: {
						...args.headers,
					},
					data: args.body,
				});
				// eslint-disable-next-line @typescript-eslint/no-explicit-any
				return {
					status: result.status,
					body: result.data,
					headers: result.headers as any,
				};
			} catch (e) {
				if (isAxiosError(e)) {
					const error = e;
					if (error.response) {
						return {
							status: error.response.status,
							body: error.response.data,
							headers: error.response.headers,
						};
					}
				}
				throw e;
			}
		},
	});

	return client;
};
