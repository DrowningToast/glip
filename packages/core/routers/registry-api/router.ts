import { initContract } from "@ts-rest/core";
import { WarehouseConnectionContract } from "./contracts/warehouse-connection";
import { WarehouseEndpointContract } from "./contracts/warehouse-endpoint";

const c = initContract();

export const RegistryApiRouter = c.router({
	warehouseConnection: WarehouseConnectionContract,
	warehouseEndpoint: WarehouseEndpointContract,
});
