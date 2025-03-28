import { useQuery } from "@tanstack/react-query";
import { shipmentApi } from "../../api/shipment";

interface UseCustomerProfileProps {
	jwt?: string;
}

export const useCustomerProfile = ({ jwt }: UseCustomerProfileProps) => {
	if (!jwt) {
		console.warn("jwt is not provided");
	}

	return useQuery({
		queryKey: ["customer-profile"],
		queryFn: () => {
			return shipmentApi.profile.getMyProfileAsCustomer();
		},
		enabled: !!jwt,
	});
};
