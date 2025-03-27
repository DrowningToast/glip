import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import Generouted from "@generouted/react-router/plugin";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
	plugins: [react(), Generouted(), tailwindcss()],
	server: {
		port: 5173,
	},
});
