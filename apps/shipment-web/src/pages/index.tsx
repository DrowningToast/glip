import { useEffect } from "react";
import { useNavigate } from "react-router";
import { ContentGuard } from "../usecase/auth/ContentGuard";
import { useSession } from "../usecase/auth/useSession";

const Page = () => {
	const { session, role } = useSession();
	const navigate = useNavigate();

	useEffect(() => {
		if (session) {
			switch (role) {
				case "WAREHOUSE":
					navigate("/warehouse/dashboard");
					break;
				case "CUSTOMER":
					navigate("/customer/dashboard");
					break;
				default:
					navigate("/login");
			}
		} else {
			navigate("/login");
		}
	}, [navigate, session, role]);

	return (
		<ContentGuard requiredAuthentication={true}>
			<h1>Home Page</h1>
		</ContentGuard>
	);
};

export default Page;
