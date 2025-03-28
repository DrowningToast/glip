import { Outlet } from "react-router";
import { CustomerHeader } from "../../components/customer-header";
import { ThemeProvider } from "../../components/theme-provider";
import { ContentGuard } from "../../usecase/auth/ContentGuard";

export default function CustomerLayout() {
	return (
		<div className="min-h-screen bg-background">
			<ContentGuard requiredAuthentication roles={{ USER: true }}>
				<ThemeProvider attribute="class" defaultTheme="system" enableSystem>
					<CustomerHeader />
					<div className="pt-16">
						<Outlet />
					</div>
				</ThemeProvider>
			</ContentGuard>
		</div>
	);
}
