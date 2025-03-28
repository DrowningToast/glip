import { initContract } from "@ts-rest/core";
import { CustomerContract } from "./contracts/customer";
import { AuthContract } from "./contracts/auth";
import { ShipmentContract } from "./contracts/shipment";
import { ProfileContract } from "./contracts/profile";

const c = initContract();

export const ShipmentApiRouter = c.router({
	customer: CustomerContract,
	auth: AuthContract,
	shipment: ShipmentContract,
	profile: ProfileContract,
});
