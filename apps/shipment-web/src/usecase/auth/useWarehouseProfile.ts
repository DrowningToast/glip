import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

interface UseWarehouseProfileProps {
	jwt?: string;
}

export const useWarehouseProfile = ({ jwt }: UseWarehouseProfileProps) => {
	return useQuery({
		queryKey: ["warehouse-profile"],
		queryFn: async () => {
			const res = await shipmentApi.profile.getMyProfileAsWarehouseConnection({
				headers: {
					authorization: `Bearer ${jwt}`,
				},
			});

			if (res.status != 200) {
				throw new Error("Failed to fetch warehouse profile");
			}

			return res.body.result.warehouseConnection;
		},
		enabled: !!jwt,
	});
};
