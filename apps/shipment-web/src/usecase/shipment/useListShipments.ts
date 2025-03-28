import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

export const useListShipments = (lastWarehouseId: number) => {
	return useQuery({
		queryKey: ["shipments"],
		queryFn: async () => {
			const res = await shipmentApi.shipment.list({
				query: {
					limit: 10,
				},
			});
			return res.body.result;
		},
	});
};
