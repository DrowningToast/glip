import { useNavigate } from "react-router"
import { signOut } from "../usecase/auth/signout"
import { Button } from "./ui/button"

export const Unauthorized: React.FC = () => {
    const navigate = useNavigate()

    const handleSignOut = () => {
        signOut()
        navigate('/login')
    }

    return (
        <div className="flex flex-col items-center justify-center min-h-screen">
            <h1 className="text-2xl font-bold mb-4">You are not authorized to view this page</h1>
            <Button
                onClick={handleSignOut}
                className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
            >
                Sign Out
            </Button>
        </div>
    )
}