import { ShipmentStatus, ShipmentStatuses } from "core/entity/shipment";
import { AlertCircle, CheckCircle2, Clock, Truck } from "lucide-react";

interface StatusBadgeProps {
	status: ShipmentStatus;
	size?: "sm" | "md" | "lg";
	showAnimation?: boolean;
}

export function StatusBadge({
	status,
	size = "md",
	showAnimation = true,
}: StatusBadgeProps) {
	// Size configurations
	const config = {
		sm: {
			container: "h-7 w-7",
			icon: "h-3.5 w-3.5",
			textPrimary: "text-xs",
			textSecondary: "text-[10px]",
			animation: "h-4 w-4",
			wrapper: "flex items-center",
		},
		md: {
			container: "h-9 w-9",
			icon: "h-5 w-5",
			textPrimary: "text-sm",
			textSecondary: "text-xs",
			animation: "h-6 w-6",
			wrapper: "flex items-center",
		},
		lg: {
			container: "h-12 w-12",
			icon: "h-6 w-6",
			textPrimary: "text-base font-medium",
			textSecondary: "text-sm",
			animation: "h-8 w-8",
			wrapper: "flex flex-col sm:flex-row items-center gap-3",
		},
	};

	const sizeConfig = config[size];

	switch (status) {
		case ShipmentStatuses.WAITING_FOR_PICKUP_TO_WAREHOUSE:
			return (
				<div className={sizeConfig.wrapper}>
					<div className="relative flex items-center justify-center">
						<div
							className={`${sizeConfig.container} rounded-full bg-amber-100 flex items-center justify-center`}
						>
							<Clock className={`${sizeConfig.icon} text-amber-600`} />
						</div>
						{showAnimation && (
							<span
								className={`animate-ping absolute ${sizeConfig.animation} rounded-full bg-amber-400 opacity-30`}
							></span>
						)}
					</div>
					<div className={size === "lg" ? "text-center sm:text-left" : "ml-3"}>
						<p className={`${sizeConfig.textPrimary} text-amber-700`}>
							Waiting for Pickup
						</p>
						<p className={`${sizeConfig.textSecondary} text-amber-500`}>
							Shipment registered
						</p>
					</div>
				</div>
			);
		case ShipmentStatuses.IN_TRANSIT_ON_THE_WAY:
			return (
				<div className={sizeConfig.wrapper}>
					<div className="relative flex items-center justify-center">
						<div
							className={`${sizeConfig.container} rounded-full bg-blue-100 flex items-center justify-center`}
						>
							<Truck className={`${sizeConfig.icon} text-blue-600`} />
						</div>
						{showAnimation && (
							<span
								className={`animate-pulse absolute ${sizeConfig.animation} rounded-full bg-blue-400 opacity-30`}
							></span>
						)}
					</div>
					<div className={size === "lg" ? "text-center sm:text-left" : "ml-3"}>
						<p className={`${sizeConfig.textPrimary} text-blue-700`}>
							In Transit
						</p>
						<p className={`${sizeConfig.textSecondary} text-blue-500`}>
							On the way
						</p>
					</div>
				</div>
			);
		case ShipmentStatuses.DELIVERED:
			return (
				<div className={sizeConfig.wrapper}>
					<div className="relative flex items-center justify-center">
						<div
							className={`${sizeConfig.container} rounded-full bg-green-100 flex items-center justify-center`}
						>
							<CheckCircle2 className={`${sizeConfig.icon} text-green-600`} />
						</div>
					</div>
					<div className={size === "lg" ? "text-center sm:text-left" : "ml-3"}>
						<p className={`${sizeConfig.textPrimary} text-green-700`}>
							Delivered
						</p>
						<p className={`${sizeConfig.textSecondary} text-green-500`}>
							Successfully delivered
						</p>
					</div>
				</div>
			);
		case ShipmentStatuses.CANCELLED:
			return (
				<div className={sizeConfig.wrapper}>
					<div className="relative flex items-center justify-center">
						<div
							className={`${sizeConfig.container} rounded-full bg-red-100 flex items-center justify-center`}
						>
							<AlertCircle className={`${sizeConfig.icon} text-red-600`} />
						</div>
					</div>
					<div className={size === "lg" ? "text-center sm:text-left" : "ml-3"}>
						<p className={`${sizeConfig.textPrimary} text-red-700`}>
							Cancelled
						</p>
						<p className={`${sizeConfig.textSecondary} text-red-500`}>
							Shipment cancelled
						</p>
					</div>
				</div>
			);
		default:
			return (
				<div className={sizeConfig.wrapper}>
					<div className="relative flex items-center justify-center">
						<div
							className={`${sizeConfig.container} rounded-full bg-gray-100 flex items-center justify-center`}
						>
							<AlertCircle className={`${sizeConfig.icon} text-gray-600`} />
						</div>
					</div>
					<div className={size === "lg" ? "text-center sm:text-left" : "ml-3"}>
						<p className={`${sizeConfig.textPrimary}`}>{status}</p>
					</div>
				</div>
			);
	}
}
