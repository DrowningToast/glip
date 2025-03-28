import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

interface UseWarehouseProfileProps {
	jwt?: string;
}

export const useWarehouseProfile = ({ jwt }: UseWarehouseProfileProps) => {
	return useQuery({
		queryKey: ["warehouse-profile"],
		queryFn: () => {
			return shipmentApi.profile.getMyProfileAsWarehouseConnection();
		},
		enabled: !!jwt,
	});
};
