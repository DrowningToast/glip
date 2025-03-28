import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

export const useTrackShipments = (shipmentId: number) => {
	return useQuery({
		queryKey: ["track-shipments"],
		queryFn: async () => {
			const res = await shipmentApi.shipment.track({
				body: {
					email: "test@test.com",
					shipment_id: shipmentId,
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
