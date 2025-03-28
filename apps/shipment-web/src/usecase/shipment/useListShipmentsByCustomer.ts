import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

interface UseListShipmentsByCustomerProps {
	jwt?: string;
}

export const useListShipmentsByCustomer = ({
	jwt,
}: UseListShipmentsByCustomerProps) => {
	return useQuery({
		queryKey: ["shipments-by-customer"],
		queryFn: async () => {
			const res = await shipmentApi.shipment.listByCustomer({
				headers: {
					authorization: `Bearer ${jwt}`,
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
		enabled: !!jwt,
	});
};
