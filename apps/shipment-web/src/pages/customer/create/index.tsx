import { ChevronLeft } from "lucide-react";
import { Link } from "react-router";
import { Button } from "../../../components/ui/button";
import { CreateShipmentForm } from "./form";

export default function CreateShipmentPage() {
	return (
		<div className="container mx-auto py-10">
			<div className="flex items-center mb-6">
				<Link to="/customer/shipments">
					<Button variant="ghost" className="flex items-center gap-1">
						<ChevronLeft className="h-4 w-4" />
						Back to Shipments
					</Button>
				</Link>
			</div>
			<div className="flex flex-col gap-2 mb-8">
				<h1 className="text-3xl font-bold tracking-tight">
					Create New Shipment
				</h1>
				<p className="text-muted-foreground">
					Fill in the details below to create a new shipment
				</p>
			</div>
			<CreateShipmentForm />
		</div>
	);
}
