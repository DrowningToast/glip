"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";

import { Loader2 } from "lucide-react";
import { useNavigate } from "react-router";
import { toast } from "sonner";
import { Button } from "../../../components/ui/button";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "../../../components/ui/form";
import { Input } from "../../../components/ui/input";
import { tryCatch } from "../../../lib/utils";
import { useSigninCustomer } from "../../../usecase/auth/useSigninCustomer";

const formSchema = z.object({
	username: z.string().min(1, {
		message: "Username is required.",
	}),
	password: z.string().min(8, {
		message: "Password must be at least 8 characters.",
	}),
	rememberMe: z.boolean().default(false),
});

export function CustomerLoginForm() {
	const navigate = useNavigate();

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			username: "",
			password: "",
			rememberMe: false,
		},
	});

	const { mutateAsync: signinCustomer, isPending: isLoading } =
		useSigninCustomer();

	async function onSubmit(values: z.infer<typeof formSchema>) {
		const [, error] = await tryCatch(() =>
			signinCustomer({
				username: values.username,
				password: values.password,
			})
		);

		if (error) {
			toast.error(error.message);
		} else {
			navigate("/customer");
		}
	}

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
				<FormField
					control={form.control}
					name="username"
					render={({ field }) => (
						<FormItem>
							<FormLabel>Username</FormLabel>
							<FormControl>
								<Input placeholder="username" {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="password"
					render={({ field }) => (
						<FormItem>
							<FormLabel>Password</FormLabel>
							<FormControl>
								<Input type="password" placeholder="••••••••" {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<Button type="submit" className="w-full" disabled={isLoading}>
					{isLoading ? (
						<>
							<Loader2 className="mr-2 h-4 w-4 animate-spin" />
							Signing in...
						</>
					) : (
						"Sign in"
					)}
				</Button>
			</form>
		</Form>
	);
}
