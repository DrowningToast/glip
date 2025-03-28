import { Plus } from "lucide-react";
import { Link } from "react-router";
import { Button } from "../../components/ui/button";
import { ShipmentList } from "./shipment-list";

export default function ShipmentsPage() {
	return (
		<div className="container mx-auto py-10">
			<div className="flex justify-between items-center mb-6">
				<div>
					<h1 className="text-3xl font-bold tracking-tight">My Shipments</h1>
					<p className="text-muted-foreground mt-1">
						View and manage all your shipments in one place
					</p>
				</div>
				<Link to="/customer/shipments/create">
					<Button className="flex items-center gap-2">
						<Plus className="h-4 w-4" />
						Create Shipment
					</Button>
				</Link>
			</div>
			<ShipmentList />
		</div>
	);
}
