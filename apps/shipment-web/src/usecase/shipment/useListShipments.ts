import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";
import { useSession } from "../auth/useSession";

export const useListShipments = (lastWarehouseId?: string) => {
	const { session } = useSession();

	return useQuery({
		queryKey: ["shipments"],
		queryFn: async () => {
			const res = await shipmentApi.shipment.list({
				query: {
					last_warehouse_id: lastWarehouseId
						? lastWarehouseId?.toString()
						: undefined,
				},
				headers: {
					authorization: `Bearer ${session}`,
				},
			});
			switch (res.status) {
				case 200:
					return res.body.result;
				case 400:
					throw new Error(res.body.message);
				case 500:
					throw new Error(res.body.message);
				default:
					throw new Error("Unknown error");
			}
		},
	});
};
