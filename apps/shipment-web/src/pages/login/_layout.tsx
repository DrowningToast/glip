import { Navigate, Outlet } from "react-router";
import { useSession } from "../../usecase/auth/useSession";

const LoginLayout = () => {
	const { role } = useSession();

	switch (role) {
		case "USER":
			return <Navigate to="/customer" />;
		case "WAREHOUSE_CONNECTION":
			return <Navigate to="/warehouse" />;
		case "ADMIN":
			return <Navigate to="/admin" />;
	}

	return <Outlet />;
};

export default LoginLayout;
