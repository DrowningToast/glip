import {
	AlertCircle,
	CheckCircle2,
	ChevronDown,
	ChevronUp,
	Clock,
	Eye,
	FileText,
	Filter,
	MapPin,
	MoreHorizontal,
	Search,
	Table,
	Truck,
} from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router";
import { Badge } from "../../components/ui/badge";
import { Button } from "../../components/ui/button";
import { Card, CardContent } from "../../components/ui/card";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "../../components/ui/dropdown-menu";
import { Input } from "../../components/ui/input";
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "../../components/ui/select";
import {
	TableBody,
	TableCell,
	TableHead,
	TableHeader,
	TableRow,
} from "../../components/ui/table";

// Mock data for shipments

// Helper function to format dates
function formatDate(dateString: string) {
	return new Date(dateString).toLocaleDateString("en-US", {
		year: "numeric",
		month: "short",
		day: "numeric",
	});
}

// Helper function to get status badge
function getStatusBadge(status: string) {
	switch (status) {
		case "WAITING_FOR_PICKUP":
			return (
				<Badge
					variant="outline"
					className="flex items-center gap-1 bg-yellow-50 text-yellow-700 border-yellow-200"
				>
					<Clock className="h-3 w-3" />
					Waiting for Pickup
				</Badge>
			);
		case "IN_TRANSIT_ON_THE_WAY":
			return (
				<Badge
					variant="outline"
					className="flex items-center gap-1 bg-blue-50 text-blue-700 border-blue-200"
				>
					<Truck className="h-3 w-3" />
					In Transit
				</Badge>
			);
		case "DELIVERED":
			return (
				<Badge
					variant="outline"
					className="flex items-center gap-1 bg-green-50 text-green-700 border-green-200"
				>
					<CheckCircle2 className="h-3 w-3" />
					Delivered
				</Badge>
			);
		case "CANCELLED":
			return (
				<Badge
					variant="outline"
					className="flex items-center gap-1 bg-red-50 text-red-700 border-red-200"
				>
					<AlertCircle className="h-3 w-3" />
					Cancelled
				</Badge>
			);
		default:
			return <Badge variant="outline">{status}</Badge>;
	}
}

export function ShipmentList() {
	const navigate = useNavigate();
	const [searchTerm, setSearchTerm] = useState("");
	const [statusFilter, setStatusFilter] = useState("all");
	const [sortField, setSortField] = useState("created_at");
	const [sortDirection, setSortDirection] = useState("desc");
	const { data: shipments } = useShipments();

	// Filter and sort shipments
	const filteredShipments = shipments
		.filter((shipment) => {
			// Apply search filter
			if (searchTerm) {
				const searchLower = searchTerm.toLowerCase();
				return (
					shipment.id.toLowerCase().includes(searchLower) ||
					shipment.departure_address.toLowerCase().includes(searchLower) ||
					shipment.destination_address.toLowerCase().includes(searchLower)
				);
			}
			return true;
		})
		.filter((shipment) => {
			// Apply status filter
			if (statusFilter === "all") return true;
			return shipment.status === statusFilter;
		})
		.sort((a, b) => {
			// Apply sorting
			const fieldA = a[sortField as keyof typeof a];
			const fieldB = b[sortField as keyof typeof b];

			if (typeof fieldA === "string" && typeof fieldB === "string") {
				return sortDirection === "asc"
					? fieldA.localeCompare(fieldB)
					: fieldB.localeCompare(fieldA);
			}

			// For dates and numbers
			if (fieldA < fieldB) return sortDirection === "asc" ? -1 : 1;
			if (fieldA > fieldB) return sortDirection === "asc" ? 1 : -1;
			return 0;
		});

	// Toggle sort direction
	const toggleSort = (field: string) => {
		if (sortField === field) {
			setSortDirection(sortDirection === "asc" ? "desc" : "asc");
		} else {
			setSortField(field);
			setSortDirection("asc");
		}
	};

	// Render sort indicator
	const renderSortIndicator = (field: string) => {
		if (sortField !== field) return null;
		return sortDirection === "asc" ? (
			<ChevronUp className="h-4 w-4 ml-1" />
		) : (
			<ChevronDown className="h-4 w-4 ml-1" />
		);
	};

	return (
		<Card>
			<CardContent className="p-6">
				<div className="flex flex-col md:flex-row gap-4 mb-6 justify-between">
					<div className="relative w-full md:w-64">
						<Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
						<Input
							placeholder="Search shipments..."
							className="pl-8"
							value={searchTerm}
							onChange={(e) => setSearchTerm(e.target.value)}
						/>
					</div>
					<div className="flex gap-2">
						<Select value={statusFilter} onValueChange={setStatusFilter}>
							<SelectTrigger className="w-[180px]">
								<div className="flex items-center gap-2">
									<Filter className="h-4 w-4" />
									<SelectValue placeholder="Filter by status" />
								</div>
							</SelectTrigger>
							<SelectContent>
								<SelectItem value="all">All Statuses</SelectItem>
								<SelectItem value="WAITING_FOR_PICKUP">
									Waiting for Pickup
								</SelectItem>
								<SelectItem value="IN_TRANSIT_ON_THE_WAY">
									In Transit
								</SelectItem>
								<SelectItem value="DELIVERED">Delivered</SelectItem>
								<SelectItem value="CANCELLED">Cancelled</SelectItem>
							</SelectContent>
						</Select>
					</div>
				</div>

				<div className="rounded-md border">
					<Table>
						<TableHeader>
							<TableRow>
								<TableHead className="w-[100px]">
									<Button
										variant="ghost"
										className="flex items-center p-0 font-medium"
										onClick={() => toggleSort("id")}
									>
										ID {renderSortIndicator("id")}
									</Button>
								</TableHead>
								<TableHead>
									<Button
										variant="ghost"
										className="flex items-center p-0 font-medium"
										onClick={() => toggleSort("departure_warehouse_id")}
									>
										From {renderSortIndicator("departure_warehouse_id")}
									</Button>
								</TableHead>
								<TableHead>
									<Button
										variant="ghost"
										className="flex items-center p-0 font-medium"
										onClick={() => toggleSort("destination_warehouse_id")}
									>
										To {renderSortIndicator("destination_warehouse_id")}
									</Button>
								</TableHead>
								<TableHead>
									<Button
										variant="ghost"
										className="flex items-center p-0 font-medium"
										onClick={() => toggleSort("status")}
									>
										Status {renderSortIndicator("status")}
									</Button>
								</TableHead>
								<TableHead className="hidden md:table-cell">
									<Button
										variant="ghost"
										className="flex items-center p-0 font-medium"
										onClick={() => toggleSort("total_weight")}
									>
										Weight (kg) {renderSortIndicator("total_weight")}
									</Button>
								</TableHead>
								<TableHead className="hidden md:table-cell">
									<Button
										variant="ghost"
										className="flex items-center p-0 font-medium"
										onClick={() => toggleSort("created_at")}
									>
										Created {renderSortIndicator("created_at")}
									</Button>
								</TableHead>
								<TableHead className="text-right">Actions</TableHead>
							</TableRow>
						</TableHeader>
						<TableBody>
							{filteredShipments.length === 0 ? (
								<TableRow>
									<TableCell colSpan={7} className="h-24 text-center">
										No shipments found.
									</TableCell>
								</TableRow>
							) : (
								filteredShipments.map((shipment) => (
									<TableRow key={shipment.id}>
										<TableCell className="font-medium">{shipment.id}</TableCell>
										<TableCell>
											<div className="flex flex-col">
												<span className="font-medium">
													{shipment.departure_warehouse_id}
												</span>
												<span className="text-xs text-muted-foreground truncate max-w-[150px]">
													{shipment.departure_address}
												</span>
											</div>
										</TableCell>
										<TableCell>
											<div className="flex flex-col">
												<span className="font-medium">
													{shipment.destination_warehouse_id}
												</span>
												<span className="text-xs text-muted-foreground truncate max-w-[150px]">
													{shipment.destination_address}
												</span>
											</div>
										</TableCell>
										<TableCell>{getStatusBadge(shipment.status)}</TableCell>
										<TableCell className="hidden md:table-cell">
											{shipment.total_weight.toFixed(2)}
										</TableCell>
										<TableCell className="hidden md:table-cell">
											{formatDate(shipment.created_at)}
										</TableCell>
										<TableCell className="text-right">
											<DropdownMenu>
												<DropdownMenuTrigger asChild>
													<Button variant="ghost" className="h-8 w-8 p-0">
														<span className="sr-only">Open menu</span>
														<MoreHorizontal className="h-4 w-4" />
													</Button>
												</DropdownMenuTrigger>
												<DropdownMenuContent align="end">
													<DropdownMenuLabel>Actions</DropdownMenuLabel>
													<DropdownMenuItem
														onClick={() =>
															navigate(`/customer/shipments/${shipment.id}`)
														}
														className="cursor-pointer"
													>
														<Eye className="h-4 w-4 mr-2" />
														View Details
													</DropdownMenuItem>
													<DropdownMenuItem
														onClick={() =>
															navigate(
																`/customer/shipments/${shipment.id}/track`
															)
														}
														className="cursor-pointer"
													>
														<MapPin className="h-4 w-4 mr-2" />
														Track Shipment
													</DropdownMenuItem>
													<DropdownMenuSeparator />
													<DropdownMenuItem
														onClick={() =>
															navigate(
																`/customer/shipments/${shipment.id}/documents`
															)
														}
														className="cursor-pointer"
													>
														<FileText className="h-4 w-4 mr-2" />
														View Documents
													</DropdownMenuItem>
												</DropdownMenuContent>
											</DropdownMenu>
										</TableCell>
									</TableRow>
								))
							)}
						</TableBody>
					</Table>
				</div>
			</CardContent>
		</Card>
	);
}
