import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";
import { shipmentApi } from "../../api/shipment";

export const useTrackShipment = () => {
	const mutation = useMutation({
		mutationFn: async (data: { shipmentId: string; email: string }) => {
			const res = await shipmentApi.shipment.track({
				body: {
					email: data.email,
					shipment_id: Number(data.shipmentId),
				},
			});
			switch (res.status) {
				case 200:
					return res.body.result.shipment;
				case 400:
					throw new Error("Shipment not found");
				case 500:
					throw new Error("Internal server error");
				default:
					throw new Error("Failed to fetch shipment");
			}
		},
		onSuccess: () => {
			toast.success("Shipment tracked successfully");
		},
		onError: () => {
			toast.error("Failed to track shipment");
		},
	});

	return mutation;
};
