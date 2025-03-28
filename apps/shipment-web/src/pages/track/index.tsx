import { ChevronLeft } from "lucide-react";
import { Link } from "react-router";
import { Button } from "../../components/ui/button";
import { PublicTrackingForm } from "../create-tracking-form";

export default function TrackPage() {
	return (
		<div className="container mx-auto flex min-h-screen flex-col">
			<div className="flex items-center py-6">
				<Link to="/">
					<Button variant="ghost" className="flex items-center gap-1">
						<ChevronLeft className="h-4 w-4" />
						Back to Home
					</Button>
				</Link>
			</div>
			<div className="flex flex-1 flex-col items-center justify-center py-12">
				<div className="mx-auto w-full max-w-md space-y-6 text-center">
					<div className="space-y-2">
						<h1 className="text-3xl font-bold">Track Your Shipment</h1>
						<p className="text-muted-foreground">
							Enter your shipment ID and email to track your package in
							real-time.
						</p>
					</div>
					<PublicTrackingForm />
				</div>
			</div>
		</div>
	);
}
