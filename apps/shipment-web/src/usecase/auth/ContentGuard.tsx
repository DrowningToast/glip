import { Role } from "core";
import React from "react";
import { useNavigate } from "react-router";
import { Skeleton } from "../../components/ui/skeleton";
import { Unauthorized } from "../../components/unauthorized";
import { useCustomerProfile } from "./useCustomerProfile";
import { useSession } from "./useSession";
import { useWarehouseProfile } from "./useWarehouseProfile";

interface ContentGuardProps {
	children: React.ReactNode;
	jwt?: string;
	requiredAuthentication: boolean;
	roles?: Partial<Record<Role, boolean>>;
}

export const ContentGuard: React.FC<ContentGuardProps> = ({
	children,
	requiredAuthentication,
	roles,
	jwt,
}) => {
	const { role, session } = useSession();
	const navigate = useNavigate();

	const { data: customer, isLoading: isCustomerProfileLoading } =
		useCustomerProfile({
			jwt: role === "USER" ? jwt : undefined,
		});

	const { data: warehouse, isLoading: isWarehouseProfileLoading } =
		useWarehouseProfile({
			jwt: role === "WAREHOUSE_CONNECTION" ? jwt : undefined,
		});

	if (!requiredAuthentication) {
		return <>{children}</>;
	}

	if (role === "" || session == "") {
		navigate("/login");
	}

	if (isCustomerProfileLoading && role === "USER") {
		return <Skeleton className="w-full h-full" />;
	}

	if (isWarehouseProfileLoading && role === "WAREHOUSE_CONNECTION") {
		return <Skeleton className="w-full h-full" />;
	}

	if (roles && !roles[role as Role]) {
		return <Unauthorized />;
	}

	return <>{children}</>;
};
