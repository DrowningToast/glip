import { Role } from "core";
import useLocalStorage from "react-use-localstorage";

export const useSession = () => {
	const [session, setSession] = useLocalStorage("session", "");
	const [role, setRole] = useLocalStorage("role", "");

	return {
		session,
		role: role as Role | string,
		setSession,
		setRole,
	};
};
