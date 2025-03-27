import { initContract } from "@ts-rest/core";
import { WarehouseConnectionContract } from "./contracts/warehouse-connection";
import { WarehouseEndpointContract } from "./contracts/warehouse-endpoint";

const c = initContract();

export const RouterApiRouter = c.router({
	warehouseConnection: WarehouseConnectionContract,
	WarehouseEndpoint: WarehouseEndpointContract,
});
