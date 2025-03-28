import { ChevronLeft } from "lucide-react";
import { Link, useParams, useSearchParams } from "react-router";
import { Button } from "../../../components/ui/button";
import { TrackingResult } from "./result";

export default function TrackingPage() {
	const { id } = useParams();
	const [searchParams] = useSearchParams();
	const email = searchParams.get("email");

	return (
		<div className="container mx-auto flex min-h-screen flex-col">
			<div className="flex items-center py-6">
				<Link to="/track">
					<Button variant="ghost" className="flex items-center gap-1">
						<ChevronLeft className="h-4 w-4" />
						Back to Tracking
					</Button>
				</Link>
			</div>
			<div className="flex flex-1 flex-col py-12">
				<div className="mx-auto w-full max-w-4xl space-y-6">
					<div className="space-y-2 text-center">
						<h1 className="text-3xl font-bold">Tracking Details</h1>
						<p className="text-muted-foreground">
							Tracking information for shipment{" "}
							<span className="font-medium">{id}</span>
						</p>
					</div>
					{email && id && <TrackingResult shipmentId={+id} email={email} />}
				</div>
			</div>
		</div>
	);
}
