import { FileText, Info, MapPin } from "lucide-react";
import { useEffect } from "react";
import { StatusBadge } from "../../../components/status-badge";
import { Button } from "../../../components/ui/button";
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from "../../../components/ui/card";
import { Separator } from "../../../components/ui/separator";
import {
	Tabs,
	TabsContent,
	TabsList,
	TabsTrigger,
} from "../../../components/ui/tabs";
import { useTrackShipment } from "../../../usecase/shipment/useTrackShipment";

interface Props {
	shipmentId: number;
	email: string;
}

// Helper function to format dates
function formatDate(dateString: string) {
	return new Date(dateString).toLocaleDateString("en-US", {
		year: "numeric",
		month: "long",
		day: "numeric",
	});
}

export const TrackingResult: React.FC<Props> = ({ shipmentId, email }) => {
	const {
		mutateAsync: fetch,
		data: shipment,
		isPending: isLoading,
		error,
	} = useTrackShipment();

	useEffect(() => {
		if (shipmentId && email) {
			fetch({ shipmentId: shipmentId.toString(), email });
		}
	}, [shipmentId, email]);

	if (isLoading) {
		return (
			<Card>
				<CardContent className="flex flex-col items-center justify-center p-6">
					<div className="h-16 w-16 animate-pulse rounded-full bg-muted"></div>
					<p className="mt-4 text-center text-muted-foreground">
						Loading tracking information...
					</p>
				</CardContent>
			</Card>
		);
	}

	if (error) {
		return (
			<Card>
				<CardContent className="flex flex-col items-center justify-center p-6">
					<div className="flex h-16 w-16 items-center justify-center rounded-full bg-red-100">
						<Info className="h-8 w-8 text-red-600" />
					</div>
					<h3 className="mt-4 text-center text-lg font-medium">
						Tracking Error
					</h3>
					<p className="mt-2 text-center text-muted-foreground">
						{error.message}
					</p>
				</CardContent>
			</Card>
		);
	}

	if (!shipment) {
		return null;
	}

	return (
		<div className="space-y-6">
			<Card>
				<CardContent className="p-6">
					<div className="flex flex-col items-center justify-center space-y-4 sm:flex-row sm:justify-between sm:space-y-0">
						<div className="flex flex-col items-center sm:items-start">
							<h2 className="text-2xl font-bold">{shipment.id}</h2>
						</div>
						<StatusBadge status={shipment.status} size="lg" />
					</div>
				</CardContent>
			</Card>

			<Tabs defaultValue="tracking">
				<TabsList className="grid w-full grid-cols-3">
					<TabsTrigger value="tracking">Tracking History</TabsTrigger>
					<TabsTrigger value="details">Shipment Details</TabsTrigger>
					<TabsTrigger value="documents">Documents</TabsTrigger>
				</TabsList>

				<TabsContent value="tracking" className="mt-4">
					<Card>
						<CardHeader>
							<CardTitle>Tracking History</CardTitle>
							<CardDescription>
								Follow the journey of your shipment from origin to destination
							</CardDescription>
						</CardHeader>
					</Card>
				</TabsContent>

				<TabsContent value="details" className="mt-4">
					<Card>
						<CardHeader>
							<CardTitle>Shipment Details</CardTitle>
							<CardDescription>
								Detailed information about your shipment
							</CardDescription>
						</CardHeader>
						<CardContent>
							<div className="grid gap-6 md:grid-cols-2">
								<div className="space-y-4">
									<div>
										<h4 className="mb-2 font-medium">Origin</h4>
										<div className="rounded-md border p-3">
											<p className="font-medium">
												{shipment.departure_warehouse_id}
											</p>
											<p className="text-sm text-muted-foreground">
												{shipment.departure_address}
											</p>
										</div>
									</div>
									<div>
										<h4 className="mb-2 font-medium">Package Information</h4>
										<div className="rounded-md border p-3">
											<div className="flex items-center justify-between py-1">
												<span className="text-sm text-muted-foreground">
													Weight:
												</span>
												<span className="font-medium">
													{shipment.total_weight.toFixed(2)} kg
												</span>
											</div>
											<Separator className="my-1" />
											<div className="flex items-center justify-between py-1">
												<span className="text-sm text-muted-foreground">
													Volume:
												</span>
												<span className="font-medium">
													{shipment.total_volume.toFixed(2)} m³
												</span>
											</div>
										</div>
									</div>
								</div>
								<div className="space-y-4">
									<div>
										<h4 className="mb-2 font-medium">Destination</h4>
										<div className="rounded-md border p-3">
											<div className="flex items-start gap-2">
												<MapPin className="mt-0.5 h-4 w-4 text-muted-foreground" />
												<div>
													<p className="font-medium">
														{shipment.destination_warehouse_id}
													</p>
													<p className="text-sm text-muted-foreground">
														{shipment.destination_address}
													</p>
												</div>
											</div>
										</div>
									</div>
									<div>
										<h4 className="mb-2 font-medium">Dates</h4>
										<div className="rounded-md border p-3">
											<div className="flex items-center justify-between py-1">
												<span className="text-sm text-muted-foreground">
													Created:
												</span>
												<span className="font-medium">
													{formatDate(shipment.created_at)}
												</span>
											</div>
											<Separator className="my-1" />
										</div>
									</div>
								</div>
							</div>
						</CardContent>
					</Card>
				</TabsContent>

				<TabsContent value="documents" className="mt-4">
					<Card>
						<CardHeader>
							<CardTitle>Shipping Documents</CardTitle>
							<CardDescription>
								Access and download shipping documents
							</CardDescription>
						</CardHeader>
						<CardContent>
							<div className="space-y-4">
								<div className="rounded-md border p-4">
									<div className="flex items-center justify-between">
										<div className="flex items-center gap-2">
											<FileText className="h-5 w-5 text-muted-foreground" />
											<div>
												<p className="font-medium">Commercial Invoice</p>
												<p className="text-sm text-muted-foreground">
													PDF document - 245 KB
												</p>
											</div>
										</div>
										<Button variant="outline" size="sm">
											Download
										</Button>
									</div>
								</div>
								<div className="rounded-md border p-4">
									<div className="flex items-center justify-between">
										<div className="flex items-center gap-2">
											<FileText className="h-5 w-5 text-muted-foreground" />
											<div>
												<p className="font-medium">Packing List</p>
												<p className="text-sm text-muted-foreground">
													PDF document - 180 KB
												</p>
											</div>
										</div>
										<Button variant="outline" size="sm">
											Download
										</Button>
									</div>
								</div>
								<div className="rounded-md border p-4">
									<div className="flex items-center justify-between">
										<div className="flex items-center gap-2">
											<FileText className="h-5 w-5 text-muted-foreground" />
											<div>
												<p className="font-medium">Bill of Lading</p>
												<p className="text-sm text-muted-foreground">
													PDF document - 320 KB
												</p>
											</div>
										</div>
										<Button variant="outline" size="sm">
											Download
										</Button>
									</div>
								</div>
							</div>
						</CardContent>
					</Card>
				</TabsContent>
			</Tabs>
		</div>
	);
};
