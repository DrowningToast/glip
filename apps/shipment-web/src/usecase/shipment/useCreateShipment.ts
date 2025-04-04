import { useMutation } from "@tanstack/react-query";
import { CreateShipmentBody } from "core/routers/shipment-api/contracts/shipment";
import { toast } from "sonner";
import { shipmentApi } from "../../api/shipment";
import { useSession } from "../auth/useSession";

export const useCreateShipment = () => {
	const { session } = useSession();

	return useMutation({
		mutationFn: async (data: CreateShipmentBody) => {
			const response = await shipmentApi.shipment.create({
				body: {
					...data,
				},
				headers: {
					authorization: `Bearer ${session}`,
				},
			});

			switch (response.status) {
				case 200:
					return response.body.result.shipment;
				case 400:
					throw new Error("Failed to create shipment");
				case 500:
					throw new Error("Failed to create shipment");
			}
		},
		onSuccess: () => {
			toast.success("Shipment created successfully");
		},
		onError: () => {
			toast.error("Failed to create shipment");
		},
	});
};
