import { useMutation } from "@tanstack/react-query";
import { queryClient } from "../../api/queryClient";
import { shipmentApi } from "../../api/shipment";

export const useSigninAdmin = () => {
	const mutation = useMutation({
		mutationFn: async (data: { key: string }) => {
			const res = await shipmentApi.auth.AuthAdminConnection({
				body: {
					key: data.key,
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
			queryClient.setQueryData(["admin"], data);
		},
		onError: (error) => {
			console.error(error);
		},
	});

	return mutation;
};
