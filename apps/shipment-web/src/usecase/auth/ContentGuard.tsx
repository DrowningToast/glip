import { Role } from "core"
import React from "react"
import { useNavigate } from "react-router"
import { Skeleton } from "../../components/ui/skeleton"
import { Unauthorized } from "../../components/unauthorized"
import { useCustomerProfile } from "./useCustomerProfile"
import { useSession } from "./useSession"
import { useWarehouseProfile } from "./useWarehouseProfile"

interface ContentGuardProps {
    children: React.ReactNode
    jwt?: string
    requiredAuthentication: boolean
    roles?: Record<Role, boolean>
}

export const ContentGuard: React.FC<ContentGuardProps> = ({ children, requiredAuthentication, roles, jwt }) => {

    const { role, session } = useSession()
    const navigate = useNavigate()


    const { isLoading: isCustomerProfileLoading } = useCustomerProfile({
        jwt: requiredAuthentication ? jwt : undefined
    })

    const { isLoading: isWarehouseProfileLoading } = useWarehouseProfile({
        jwt: requiredAuthentication ? jwt : undefined
    })

    if (!requiredAuthentication) {
        return <>{children}</>
    }

    if (role === '' || session == '') {
        navigate('/login')
    }

    if (isCustomerProfileLoading || isWarehouseProfileLoading) {
        return <Skeleton className="w-full h-full" />
    }

    if (roles && !roles[role as Role]) {
        return <Unauthorized />
    }


    return (
        <>{children}</>
    )
}