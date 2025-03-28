import { Link } from "react-router";
import { Button } from "../../../components/ui/button";
import { ChevronLeft } from "lucide-react";
import { WarehouseLoginForm } from "./form";

export default function WarehouseLoginPage() {
    return (
        <div className="container flex h-screen w-screen flex-col items-center justify-center">
            <Link to="/login" className="absolute left-4 top-4 md:left-8 md:top-8">
                <Button variant="ghost" className="flex items-center gap-1">
                    <ChevronLeft className="h-4 w-4" />
                    Back
                </Button>
            </Link>
            <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
                <div className="flex flex-col space-y-2 text-center">
                    <h1 className="text-2xl font-semibold tracking-tight">Warehouse Staff Login</h1>
                    <p className="text-sm text-muted-foreground">Enter your warehouse credentials to access the system</p>
                </div>
                <WarehouseLoginForm />
                <p className="px-8 text-center text-sm text-muted-foreground">
                    Contact your administrator if you need access or have forgotten your credentials.
                </p>
            </div>
        </div>
    )
}

