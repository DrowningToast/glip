import axiosClient from "@/libs/axiosClient";

import { Inventory, InventoryStatus } from "./type";

export const getAllShipment = async (): Promise<Inventory[]> => {
    const response = await axiosClient.get('/inventory/')
    return response.data;
}

export const getShipmentById = async (id: string): Promise<Inventory> => {
    const response = await axiosClient.get(`/inventory/${id}`)
    return response.data;
}

export const updateInventoryStatus = async (id: string, status: InventoryStatus): Promise<void> => {
    const response = await axiosClient.put(`/inventory/${id}/status`, { status })
    return response.data;
}