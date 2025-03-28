import { NewAPIClient } from "core";
import { RegistryApiRouter } from "core/routers/registry-api/router";
import { axiosClient } from "./axios";

export const registryApi = NewAPIClient(
	axiosClient,
	RegistryApiRouter,
	import.meta.env.VITE_REGISTRY_API_URL
);
