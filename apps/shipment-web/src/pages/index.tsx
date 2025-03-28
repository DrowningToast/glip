import { Link } from "react-router";
import { ThemeProvider } from "../components/theme-provider";
import { Button } from "../components/ui/button";

const Page = () => {
	return (
		<ThemeProvider attribute="class" defaultTheme="system" enableSystem>
			<div className="flex min-h-screen flex-col">
				<header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
					<div className="container flex h-16 items-center justify-between">
						<div className="flex items-center gap-2">
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
							<span className="font-bold">
								Global Logistics Integration Platform
							</span>
						</div>
						<nav className="hidden gap-6 md:flex">
							<Link
								to="/track"
								className="text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
							>
								Track Shipment
							</Link>
							<Link
								to="/about"
								className="text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
							>
								About
							</Link>
							<Link
								to="/contact"
								className="text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
							>
								Contact
							</Link>
						</nav>
						<div>
							<Link to="/login">
								<Button>Login</Button>
							</Link>
						</div>
					</div>
				</header>
				<main className="flex-1">
					<section className="w-full py-12 md:py-24 lg:py-32 xl:py-48">
						<div className="container px-4 md:px-6">
							<div className="grid gap-6 lg:grid-cols-2 lg:gap-12 xl:grid-cols-2">
								<div className="flex flex-col justify-center space-y-4">
									<div className="space-y-2">
										<h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none">
											Global Logistics Integration Platform
										</h1>
										<p className="max-w-[600px] text-muted-foreground md:text-xl">
											Streamlining international logistics with real-time
											tracking and efficient route management.
										</p>
									</div>
									<div className="flex flex-col gap-2 min-[400px]:flex-row">
										<Link to="/login">
											<Button size="lg" className="w-full min-[400px]:w-auto">
												Get Started
											</Button>
										</Link>
										<Link to="/track">
											<Button
												size="lg"
												variant="outline"
												className="w-full min-[400px]:w-auto"
											>
												Track Your Shipment
											</Button>
										</Link>
									</div>
								</div>
								<div className="flex items-center justify-center">
									<div className="relative h-[350px] w-[350px] sm:h-[450px] sm:w-[450px]">
										<img
											src="/placeholder.svg?height=450&width=450"
											width={450}
											height={450}
											alt="Global Logistics"
											className="rounded-lg object-cover"
										/>
									</div>
								</div>
							</div>
						</div>
					</section>
					<section className="w-full bg-muted py-12 md:py-24 lg:py-32">
						<div className="container px-4 md:px-6">
							<div className="mx-auto flex max-w-[58rem] flex-col items-center justify-center space-y-4 text-center">
								<h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl">
									Track Your Shipment
								</h2>
								<p className="max-w-[85%] text-muted-foreground md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed">
									Enter your shipment ID and email to track your package in
									real-time.
								</p>
							</div>
							<div className="mx-auto mt-8 grid w-full max-w-3xl gap-6">
								<PublicTrackingForm />
							</div>
						</div>
					</section>
				</main>
				<footer className="border-t py-6 md:py-0">
					<div className="container flex flex-col items-center justify-between gap-4 md:h-24 md:flex-row">
						<p className="text-center text-sm leading-loose text-muted-foreground md:text-left">
							© 2023 Global Logistics Integration Platform. All rights
							reserved.
						</p>
						<div className="flex gap-4">
							<Link
								to="/terms"
								className="text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
							>
								Terms of Service
							</Link>
							<Link
								to="/privacy"
								className="text-sm font-medium text-muted-foreground transition-colors hover:text-foreground"
							>
								Privacy Policy
							</Link>
						</div>
					</div>
				</footer>
			</div>
		</ThemeProvider>
	);
};

export default Page;
