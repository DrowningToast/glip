import { useNavigate } from "react-router";

export const useSignout = () => {
	const navigate = useNavigate();

	return () => {
		localStorage.removeItem("session");
		localStorage.removeItem("role");
		navigate("/login");
	};
};
