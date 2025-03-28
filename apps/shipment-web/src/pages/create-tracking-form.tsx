import { zodResolver } from "@hookform/resolvers/zod";
import { Search } from "lucide-react";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import * as z from "zod";
import { Button } from "../components/ui/button";
import { Card, CardContent } from "../components/ui/card";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "../components/ui/form";
import { Input } from "../components/ui/input";

const formSchema = z.object({
	shipmentId: z.string().min(1, {
		message: "Shipment ID is required",
	}),
	email: z.string().email({
		message: "Please enter a valid email address",
	}),
});

export function PublicTrackingForm() {
	const navigate = useNavigate();
	const [isSubmitting, setIsSubmitting] = useState(false);

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			shipmentId: "",
			email: "",
		},
	});

	function onSubmit(values: z.infer<typeof formSchema>) {
		setIsSubmitting(true);

		navigate(
			`/track/${values.shipmentId}?email=${encodeURIComponent(values.email)}`
		);
	}

	return (
		<Card className="w-full">
			<CardContent className="pt-6">
				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
						<div className="grid gap-4 sm:grid-cols-2">
							<FormField
								control={form.control}
								name="shipmentId"
								render={({ field }) => (
									<FormItem>
										<FormLabel>Shipment ID</FormLabel>
										<FormControl>
											<Input
												placeholder="Enter shipment ID (e.g. 1, 2, 3, ...)"
												{...field}
											/>
										</FormControl>
										<FormMessage />
									</FormItem>
								)}
							/>
							<FormField
								control={form.control}
								name="email"
								render={({ field }) => (
									<FormItem>
										<FormLabel>Email Address</FormLabel>
										<FormControl>
											<Input
												placeholder="Enter your email address"
												type="email"
												{...field}
											/>
										</FormControl>
										<FormMessage />
									</FormItem>
								)}
							/>
						</div>
						<Button type="submit" className="w-full" disabled={isSubmitting}>
							{isSubmitting ? (
								<>
									<Search className="mr-2 h-4 w-4 animate-pulse" />
									Tracking...
								</>
							) : (
								<>
									<Search className="mr-2 h-4 w-4" />
									Track Shipment
								</>
							)}
						</Button>
					</form>
				</Form>
			</CardContent>
		</Card>
	);
}
