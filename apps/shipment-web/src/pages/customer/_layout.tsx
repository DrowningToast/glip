import type { ReactNode } from "react";
import { CustomerHeader } from "../../components/customer-header";
import { ThemeProvider } from "../../components/theme-provider";
import { ContentGuard } from "../../usecase/auth/ContentGuard";

export default function CustomerLayout({ children }: { children: ReactNode }) {
	return (
		<div className="min-h-screen bg-background">
			<ContentGuard requiredAuthentication roles={{ USER: true }}>
				<ThemeProvider attribute="class" defaultTheme="system" enableSystem>
					<CustomerHeader />
					<main className="pt-16">{children}</main>
				</ThemeProvider>
			</ContentGuard>
		</div>
	);
}
