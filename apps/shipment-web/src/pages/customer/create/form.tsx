import { zodResolver } from "@hookform/resolvers/zod";
import { Loader2 } from "lucide-react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { toast } from "sonner";
import * as z from "zod";
import { Button } from "../../../components/ui/button";
import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from "../../../components/ui/card";
import {
	Form,
	FormControl,
	FormDescription,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "../../../components/ui/form";
import { Input } from "../../../components/ui/input";
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "../../../components/ui/select";
import { Separator } from "../../../components/ui/separator";
import { Textarea } from "../../../components/ui/textarea";
import { tryCatch } from "../../../lib/utils";
import { useCustomerProfile } from "../../../usecase/auth/useCustomerProfile";
import { useSession } from "../../../usecase/auth/useSession";
import { useCreateShipment } from "../../../usecase/shipment/useCreateShipment";
//
// Add this near the top of the file, after imports
const WAREHOUSES = [
	{ id: "USA1", name: "USA1", region: "North America" },
	{ id: "USA2", name: "USA2", region: "North America" },
	{ id: "USA3", name: "USA3", region: "North America" },
	{ id: "EU1", name: "EU1", region: "Europe" },
	{ id: "EU2", name: "EU2", region: "Europe" },
	{ id: "EU3", name: "EU3", region: "Europe" },
	{ id: "APAC1", name: "APAC1", region: "Asia-Pacific" },
	{ id: "APAC2", name: "APAC2", region: "Asia-Pacific" },
	{ id: "APAC3", name: "APAC3", region: "Asia-Pacific" },
] as const;

// Form schema for shipment creation
const formSchema = z.object({
	departure_warehouse_id: z.string({
		required_error: "Please select a departure warehouse",
	}),
	departure_address: z.string().min(5, {
		message: "Departure address must be at least 5 characters",
	}),
	destination_warehouse_id: z.string({
		required_error: "Please select a destination warehouse",
	}),
	destination_address: z.string().min(5, {
		message: "Destination address must be at least 5 characters",
	}),
	total_weight: z.coerce.number().positive({
		message: "Weight must be a positive number",
	}),
	total_volume: z.coerce.number().positive({
		message: "Volume must be a positive number",
	}),
	special_instructions: z.string().optional(),
});

export function CreateShipmentForm() {
	const navigate = useNavigate();

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			departure_address: "",
			destination_address: "",
			special_instructions: "",
		},
	});

	const { session } = useSession();
	const { data: customerProfile, isLoading: isProfileLoading } =
		useCustomerProfile({
			jwt: session,
		});
	const { mutateAsync: createShipment, isPending: isLoading } =
		useCreateShipment();
	async function onSubmit(values: z.infer<typeof formSchema>) {
		const [, err] = await tryCatch(() => {
			if (isProfileLoading) {
				throw new Error("Failed to fetch customer profile");
			}

			return createShipment({
				shipment: {
					...values,
					owner_id: Number(customerProfile?.id),
					departure_city: values.departure_address.split(",")[0],
					destination_city: values.destination_address.split(",")[0],
				},
			});
		});
		if (err) {
			toast.error(err.message);
			return;
		}

		// Redirect to shipments list with success message
		navigate("/customer");
	}

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
				<Card>
					<CardHeader>
						<CardTitle>Shipment Details</CardTitle>
						<CardDescription>
							Enter the details of your shipment including origin, destination,
							and package information.
						</CardDescription>
					</CardHeader>
					<CardContent className="space-y-6">
						<div>
							<h3 className="text-lg font-medium mb-4">Origin Information</h3>
							<div className="grid gap-6 sm:grid-cols-2">
								<FormField
									control={form.control}
									name="departure_warehouse_id"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Departure Warehouse</FormLabel>
											<Select
												onValueChange={field.onChange}
												defaultValue={field.value}
											>
												<FormControl>
													<SelectTrigger>
														<SelectValue placeholder="Select departure warehouse" />
													</SelectTrigger>
												</FormControl>
												<SelectContent>
													{WAREHOUSES.map((warehouse) => (
														<SelectItem key={warehouse.id} value={warehouse.id}>
															{warehouse.name} ({warehouse.region})
														</SelectItem>
													))}
												</SelectContent>
											</Select>
											<FormDescription>
												Select the warehouse where your shipment will depart
												from.
											</FormDescription>
											<FormMessage />
										</FormItem>
									)}
								/>
								<FormField
									control={form.control}
									name="departure_address"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Departure Address</FormLabel>
											<FormControl>
												<Textarea
													placeholder="Enter the full departure address"
													className="resize-none"
													{...field}
												/>
											</FormControl>
											<FormDescription>
												Provide the complete address for pickup.
											</FormDescription>
											<FormMessage />
										</FormItem>
									)}
								/>
							</div>
						</div>

						<Separator />

						<div>
							<h3 className="text-lg font-medium mb-4">
								Destination Information
							</h3>
							<div className="grid gap-6 sm:grid-cols-2">
								<FormField
									control={form.control}
									name="destination_warehouse_id"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Destination Warehouse</FormLabel>
											<Select
												onValueChange={field.onChange}
												defaultValue={field.value}
											>
												<FormControl>
													<SelectTrigger>
														<SelectValue placeholder="Select destination warehouse" />
													</SelectTrigger>
												</FormControl>
												<SelectContent>
													{WAREHOUSES.map((warehouse) => (
														<SelectItem key={warehouse.id} value={warehouse.id}>
															{warehouse.name} ({warehouse.region})
														</SelectItem>
													))}
												</SelectContent>
											</Select>
											<FormDescription>
												Select the warehouse where your shipment will be
												delivered.
											</FormDescription>
											<FormMessage />
										</FormItem>
									)}
								/>
								<FormField
									control={form.control}
									name="destination_address"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Destination Address</FormLabel>
											<FormControl>
												<Textarea
													placeholder="Enter the full destination address"
													className="resize-none"
													{...field}
												/>
											</FormControl>
											<FormDescription>
												Provide the complete address for delivery.
											</FormDescription>
											<FormMessage />
										</FormItem>
									)}
								/>
							</div>
						</div>

						<Separator />

						<div>
							<h3 className="text-lg font-medium mb-4">Package Information</h3>
							<div className="grid gap-6 sm:grid-cols-2">
								<FormField
									control={form.control}
									name="total_weight"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Total Weight (kg)</FormLabel>
											<FormControl>
												<Input
													type="number"
													step="0.01"
													placeholder="Enter weight in kilograms"
													{...field}
												/>
											</FormControl>
											<FormDescription>
												Enter the total weight of your shipment in kilograms.
											</FormDescription>
											<FormMessage />
										</FormItem>
									)}
								/>
								<FormField
									control={form.control}
									name="total_volume"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Total Volume (m³)</FormLabel>
											<FormControl>
												<Input
													type="number"
													step="0.01"
													placeholder="Enter volume in cubic meters"
													{...field}
												/>
											</FormControl>
											<FormDescription>
												Enter the total volume of your shipment in cubic meters.
											</FormDescription>
											<FormMessage />
										</FormItem>
									)}
								/>
							</div>
						</div>

						<FormField
							control={form.control}
							name="special_instructions"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Special Instructions (Optional)</FormLabel>
									<FormControl>
										<Textarea
											placeholder="Enter any special handling instructions or notes"
											className="resize-none"
											{...field}
										/>
									</FormControl>
									<FormDescription>
										Provide any special handling instructions, fragile items, or
										other important notes.
									</FormDescription>
									<FormMessage />
								</FormItem>
							)}
						/>
					</CardContent>
					<CardFooter className="flex justify-between">
						<Button
							type="button"
							variant="outline"
							onClick={() => navigate("/customer/shipments")}
						>
							Cancel
						</Button>
						<Button type="submit" disabled={isLoading}>
							{isLoading ? (
								<>
									<Loader2 className="mr-2 h-4 w-4 animate-spin" />
									Creating Shipment...
								</>
							) : (
								"Create Shipment"
							)}
						</Button>
					</CardFooter>
				</Card>
			</form>
		</Form>
	);
}
