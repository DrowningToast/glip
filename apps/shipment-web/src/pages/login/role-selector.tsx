"use client"

import { ShieldCheck, UserCircle, Warehouse } from "lucide-react"
import { useState } from "react"
import { useNavigate } from "react-router"
import { Card, CardDescription, CardHeader, CardTitle } from "../../components/ui/card"

export function RoleSelector() {
    const navigate = useNavigate()
    const [selectedRole, setSelectedRole] = useState<string | null>(null)

    const handleRoleSelect = (role: string) => {
        setSelectedRole(role)
        navigate(`/login/${role.toLowerCase()}`)
    }

    return (
        <div className="grid gap-4">
            <Card
                className={`cursor-pointer hover:border-primary transition-colors ${selectedRole === "customer" ? "border-primary" : ""}`}
                onClick={() => handleRoleSelect("customer")}
            >
                <CardHeader className="flex flex-row items-center gap-4 p-4">
                    <UserCircle className="h-8 w-8 text-primary" />
                    <div>
                        <CardTitle className="text-xl">Customer</CardTitle>
                        <CardDescription>For shippers and receivers</CardDescription>
                    </div>
                </CardHeader>
            </Card>

            <Card
                className={`cursor-pointer hover:border-primary transition-colors ${selectedRole === "warehouse" ? "border-primary" : ""}`}
                onClick={() => handleRoleSelect("warehouse")}
            >
                <CardHeader className="flex flex-row items-center gap-4 p-4">
                    <Warehouse className="h-8 w-8 text-primary" />
                    <div>
                        <CardTitle className="text-xl">Warehouse Staff</CardTitle>
                        <CardDescription>For warehouse and distribution center staff</CardDescription>
                    </div>
                </CardHeader>
            </Card>

            <Card
                className={`cursor-pointer hover:border-primary transition-colors ${selectedRole === "admin" ? "border-primary" : ""}`}
                onClick={() => handleRoleSelect("admin")}
            >
                <CardHeader className="flex flex-row items-center gap-4 p-4">
                    <ShieldCheck className="h-8 w-8 text-primary" />
                    <div>
                        <CardTitle className="text-xl">Admin</CardTitle>
                        <CardDescription>For system administrators</CardDescription>
                    </div>
                </CardHeader>
            </Card>
        </div>
    )
}

