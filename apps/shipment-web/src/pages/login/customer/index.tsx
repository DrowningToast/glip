import { ChevronLeft } from "lucide-react"
import { Link } from "react-router"
import { Button } from "../../../components/ui/button"
import { CustomerLoginForm } from "./form"


export default function CustomerLoginPage() {
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
                    <h1 className="text-2xl font-semibold tracking-tight">Customer Login</h1>
                    <p className="text-sm text-muted-foreground">Enter your credentials to access your account</p>
                </div>
                <CustomerLoginForm />
                <p className="px-8 text-center text-sm text-muted-foreground">
                    <Link to="/login/customer/register" className="hover:text-brand underline underline-offset-4">
                        Don&apos;t have an account? Sign up
                    </Link>
                </p>
            </div>
        </div>
    )
}

