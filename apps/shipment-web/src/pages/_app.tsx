import { QueryClientProvider } from "@tanstack/react-query"
import { Outlet } from "react-router"
import { queryClient } from "../api/queryClient"

const RootLayout = () => {
    return (
        <QueryClientProvider client={queryClient}>
            <main className="max-w-screen flex flex-col items-center justify-center"><Outlet /></main>
        </QueryClientProvider>
    )
}

export default RootLayout