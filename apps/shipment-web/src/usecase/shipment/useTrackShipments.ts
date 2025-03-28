import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

export const useTrackShipments = (shipmentId: number) => {
	return useQuery({
		queryKey: ["track-shipments"],
		queryFn: async () => {
			const res = await shipmentApi.shipment.track({
				query: {
					shipment_id: shipmentId,
				},
			});
			return res.body.result;
		},
	});
};
