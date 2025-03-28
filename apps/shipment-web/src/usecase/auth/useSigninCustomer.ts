import { useMutation } from "@tanstack/react-query";
import { queryClient } from "../../api/queryClient";
import { shipmentApi } from "../../api/shipment";
import { useSession } from "./useSession";

export const useSigninCustomer = () => {
	const { setSession, setRole } = useSession();
	const mutation = useMutation({
		mutationFn: async (data: { username: string; password: string }) => {
			const res = await shipmentApi.auth.AuthCustomerConnection({
				body: {
					username: data.username,
					password: data.password,
				},
			});
			switch (res.status) {
				case 200:
					return res.body.result.jwt;
				case 400:
					throw new Error(res.body.message);
				case 500:
					throw new Error("Internal server error");
				default:
					throw new Error("Unknown error");
			}
		},
		onSuccess: (data) => {
			queryClient.setQueryData(["customer"], data);
			setSession(data);
			setRole("customer");
		},
		onError: (error) => {
			console.error(error);
		},
	});

	return mutation;
};
