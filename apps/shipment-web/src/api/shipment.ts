import { NewAPIClient } from "core";
import { ShipmentApiRouter } from "core/routers/shipment-api/router";
import { axiosClient } from "./axios";

export const shipmentApi = NewAPIClient(
	axiosClient,
	ShipmentApiRouter,
	import.meta.env.VITE_SHIPMENT_API_URL
);
