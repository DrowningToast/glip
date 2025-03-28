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
		queryFn: async () => {
			const res = await shipmentApi.profile.getMyProfileAsCustomer({
				headers: {
					authorization: `Bearer ${jwt}`,
				},
			});

			if (res.status != 200) {
				throw new Error("Failed to fetch customer profile");
			}

			return res.body.result.customer;
		},
		enabled: !!jwt,
	});
};
