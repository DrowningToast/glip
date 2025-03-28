import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import * as z from "zod";

import { Loader2 } from "lucide-react";
import { toast } from "sonner";
import { Button } from "../../../components/ui/button";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
} from "../../../components/ui/form";
import { Input } from "../../../components/ui/input";
import { tryCatch } from "../../../lib/utils";
import { useSigninWarehouse } from "../../../usecase/auth/useSigninWarehouse";

const formSchema = z.object({
	apiKey: z.string({
		required_error: "Please enter your API key.",
	}),
});

export function WarehouseLoginForm() {
	const navigate = useNavigate();

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			apiKey: "",
		},
	});

	const { mutateAsync: signinWarehouse, isPending: isLoading } =
		useSigninWarehouse();

	async function onSubmit(values: z.infer<typeof formSchema>) {
		const [, err] = await tryCatch(() =>
			signinWarehouse({
				key: values.apiKey,
			})
		);

		if (err) {
			toast.error(err.message);
			return;
		}

		// Redirect to warehouse dashboard after successful login
		navigate("/warehouse");
	}

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
				<FormField
					control={form.control}
					name="apiKey"
					render={({ field }) => (
						<FormItem>
							<FormLabel>API Key</FormLabel>
							<FormControl>
								<Input {...field} />
							</FormControl>
						</FormItem>
					)}
				/>
				<Button type="submit" className="w-full" disabled={isLoading}>
					{isLoading ? (
						<>
							<Loader2 className="mr-2 h-4 w-4 animate-spin" />
							Authenticating...
						</>
					) : (
						"Connect to Warehouse"
					)}
				</Button>
			</form>
		</Form>
	);
}
