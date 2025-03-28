import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { Form, useForm } from "react-hook-form";
import * as z from "zod";

import { AlertCircle, Loader2 } from "lucide-react";
import { useNavigate } from "react-router";
import { Alert, AlertDescription } from "../../../components/ui/alert";
import { Button } from "../../../components/ui/button";
import {
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "../../../components/ui/form";
import { Input } from "../../../components/ui/input";
import { tryCatch } from "../../../lib/utils";
import { useSigninAdmin } from "../../../usecase/auth/useSigninAdmin";

const formSchema = z.object({
	apiKey: z.string().min(32, {
		message: "API key must be at least 32 characters.",
	}),
});

export function AdminLoginForm() {
	const navigate = useNavigate();

	const [error, setError] = useState<string | null>(null);

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			apiKey: "",
		},
	});

	const { mutateAsync: login, isPending: isLoading } = useSigninAdmin();
	async function onSubmit(values: z.infer<typeof formSchema>) {
		const [, err] = await tryCatch(() => login({ key: values.apiKey }));

		if (err) {
			setError(err.message);
			return;
		}

		// In a real app, you would validate the API key here
		// For demo purposes, we'll just redirect
		navigate("/admin/dashboard");
	}

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
				{error && (
					<Alert variant="destructive">
						<AlertCircle className="h-4 w-4" />
						<AlertDescription>{error}</AlertDescription>
					</Alert>
				)}
				<FormField
					control={form.control}
					name="apiKey"
					render={({ field }) => (
						<FormItem>
							<FormLabel>API Key</FormLabel>
							<FormControl>
								<Input
									placeholder="Enter your API key"
									{...field}
									type="password"
									autoComplete="off"
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<Button type="submit" className="w-full" disabled={isLoading}>
					{isLoading ? (
						<>
							<Loader2 className="mr-2 h-4 w-4 animate-spin" />
							Verifying...
						</>
					) : (
						"Access Admin Panel"
					)}
				</Button>
			</form>
		</Form>
	);
}
