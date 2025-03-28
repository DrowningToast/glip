"use client";

import { LogOut, Menu, Package, Settings, Truck, User, X } from "lucide-react";
import { useState } from "react";
import { Link, useLocation } from "react-router";
import { Button } from "../../components/ui/button";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "../../components/ui/dropdown-menu";
import { useCustomerProfile } from "../../usecase/auth/useCustomerProfile";
import { useSession } from "../../usecase/auth/useSession";
import { useSignout } from "../../usecase/auth/useSignout";

export function CustomerHeader() {
	const location = useLocation();
	const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

	const { session: jwt } = useSession();
	const { data: profile } = useCustomerProfile({ jwt });
	const signOut = useSignout();

	const navigation = [
		{ name: "Shipments", href: "/customer", icon: Package },
		{ name: "Tracking", href: "/customer/tracking", icon: Truck },
	];

	return (
		<header className="fixed top-0 left-0 right-0 z-50 bg-background border-b">
			<nav
				className="mx-auto flex max-w-7xl items-center justify-between p-4 lg:px-8"
				aria-label="Global"
			>
				<div className="flex lg:flex-1">
					<Link
						to="/customer/dashboard"
						className="-m-1.5 p-1.5 flex items-center gap-2"
					>
						<span className="sr-only">
							Global Logistics Integration Platform
						</span>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							strokeWidth="2"
							strokeLinecap="round"
							strokeLinejoin="round"
							className="h-6 w-6"
						>
							<path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10" />
						</svg>
						<span className="font-semibold">Global Logistics</span>
					</Link>
				</div>

				<div className="flex lg:hidden">
					<button
						type="button"
						className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5"
						onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
					>
						<span className="sr-only">Open main menu</span>
						{mobileMenuOpen ? (
							<X className="h-6 w-6" aria-hidden="true" />
						) : (
							<Menu className="h-6 w-6" aria-hidden="true" />
						)}
					</button>
				</div>

				<div className="hidden lg:flex lg:gap-x-8">
					{navigation.map((item) => (
						<Link
							key={item.name}
							to={item.href}
							className={`flex items-center gap-1 text-sm font-medium leading-6 ${
								location.pathname.startsWith(item.href)
									? "text-primary"
									: "text-muted-foreground hover:text-foreground"
							}`}
						>
							<item.icon className="h-4 w-4" />
							{item.name}
						</Link>
					))}
				</div>

				<div className="hidden lg:flex lg:flex-1 lg:justify-end">
					<DropdownMenu>
						<DropdownMenuTrigger asChild>
							<Button variant="ghost" className="relative h-8 w-8 rounded-full">
								<User className="h-5 w-5" />
							</Button>
						</DropdownMenuTrigger>
						<DropdownMenuContent align="end">
							<DropdownMenuLabel>{profile?.name}</DropdownMenuLabel>
							<DropdownMenuSeparator />
							<DropdownMenuItem>
								<User className="mr-2 h-4 w-4" />
								<span>Profile</span>
							</DropdownMenuItem>
							<DropdownMenuItem>
								<Settings className="mr-2 h-4 w-4" />
								<span>Settings</span>
							</DropdownMenuItem>
							<DropdownMenuSeparator />
							<DropdownMenuItem>
								<LogOut className="mr-2 h-4 w-4" />
								<span>Log out</span>
							</DropdownMenuItem>
						</DropdownMenuContent>
					</DropdownMenu>
				</div>
			</nav>

			{/* Mobile menu */}
			{mobileMenuOpen && (
				<div className="lg:hidden">
					<div className="fixed inset-0 z-50"></div>
					<div className="fixed inset-y-0 right-0 z-50 w-full overflow-y-auto bg-background px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
						<div className="mt-6 flow-root">
							<div className="-my-6 divide-y">
								<div className="space-y-2 py-6">
									{navigation.map((item) => (
										<Link
											key={item.name}
											to={item.href}
											className={`-mx-3 flex items-center gap-2 rounded-lg px-3 py-2 text-base font-semibold leading-7 ${
												location.pathname.startsWith(item.href)
													? "bg-accent text-accent-foreground"
													: "hover:bg-muted"
											}`}
											onClick={() => setMobileMenuOpen(false)}
										>
											<item.icon className="h-5 w-5" />
											{item.name}
										</Link>
									))}
								</div>
								<div className="py-6">
									<Link
										to="/customer/profile"
										className="-mx-3 flex items-center gap-2 rounded-lg px-3 py-2 text-base font-semibold leading-7 hover:bg-muted"
										onClick={() => setMobileMenuOpen(false)}
									>
										<User className="h-5 w-5" />
										Profile
									</Link>
									<Link
										to="/customer/settings"
										className="-mx-3 flex items-center gap-2 rounded-lg px-3 py-2 text-base font-semibold leading-7 hover:bg-muted"
										onClick={() => setMobileMenuOpen(false)}
									>
										<Settings className="h-5 w-5" />
										Settings
									</Link>
									<Link
										to="/login"
										className="-mx-3 flex items-center gap-2 rounded-lg px-3 py-2 text-base font-semibold leading-7 hover:bg-muted"
										onClick={() => {
											setMobileMenuOpen(false);
											signOut();
										}}
									>
										<LogOut className="h-5 w-5" />
										Log out
									</Link>
								</div>
							</div>
						</div>
					</div>
				</div>
			)}
		</header>
	);
}
