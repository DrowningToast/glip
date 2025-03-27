import { initContract } from "@ts-rest/core";
import { CustomerContract } from "./contracts/customer";
import { AuthContract } from "./contracts/auth";

const c = initContract();

export const ShipmentApiRouter = c.router({
	customer: CustomerContract,
	auth: AuthContract,
});
