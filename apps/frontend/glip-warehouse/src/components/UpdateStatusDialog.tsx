import { useState } from "react"
import { toast } from "react-hot-toast"
import { useQueryClient, useMutation } from "@tanstack/react-query"

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
  } from "@/components/ui/dialog"
import { Button } from "./ui/button"
import { Select, SelectTrigger, SelectContent, SelectItem, SelectValue } from "./ui/select"

import { InventoryStatus } from "@/queries/inventory/type"
import { updateInventoryStatus } from "@/queries/inventory/query"

type UpdateStatusDialogProps = {
    id: string
    shipmentId: string
    status: InventoryStatus
}

export function UpdateStatusDialog({ id, shipmentId, status }: UpdateStatusDialogProps) {
    const [open, setOpen] = useState<boolean>(false)
    const [newStatus, setNewStatus] = useState<InventoryStatus>(status)
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const queryClient = useQueryClient()
    const mutation = useMutation({
        mutationFn: ({ id, status }: { id: string, status: InventoryStatus }) => 
            updateInventoryStatus(id, status),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["shipments"] })
        }
    })

    const handleUpdateStatus = async () => {
        setIsLoading(true)
        await mutation.mutateAsync({ id, status: newStatus })
        setIsLoading(false)
        setOpen(false)
        toast.success("Status updated successfully")
    }

    return (
        <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
                <Button variant="outline" size="sm">Update Status</Button>
            </DialogTrigger>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Update Status</DialogTitle>
                    <DialogDescription>
                        Update the status of the shipment
                    </DialogDescription>
                    <hr />
                    <div className="flex flex-col gap-2 text-sm">
                        <p>Shipment ID: {shipmentId}</p>
                        <p>Current Status: {status}</p>
                        <div className="flex items-center gap-2">
                            <p>New Status:</p>
                            <Select defaultValue={status} onValueChange={(value) => {
                                setNewStatus(value as InventoryStatus)
                            }}>
                            <SelectTrigger>
                                <SelectValue placeholder="Select Status" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="DELIVERED">Delivered</SelectItem>
                                <SelectItem value="CANCELLED">Cancelled</SelectItem>
                                <SelectItem value="INCOMING_SHIPMENT">Incoming Shipment</SelectItem>
                                <SelectItem value="WAREHOUSE_RECEIVED">Warehouse Received</SelectItem>
                                <SelectItem value="WAREHOUSE_DEPARTED">Warehouse Departed</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                    </div>
                </DialogHeader>
                <DialogFooter>
                    <Button variant="outline" size="sm" onClick={() => setOpen(false)}>Cancel</Button>
                    <Button variant="outline" size="sm" onClick={handleUpdateStatus} disabled={isLoading}>Update</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}