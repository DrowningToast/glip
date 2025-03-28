export const signOut = () => {
	localStorage.removeItem("session");
	localStorage.removeItem("role");
};
