export type InventoryStatus = "DELIVERED" | "CANCELLED" | "INCOMING_SHIPMENT" | "WAREHOUSE_RECEIVED" | "WAREHOUSE_DEPARTED"

export type Inventory = {
    id: string;
    shipmentId: string;
    route: string[];
    last_warehouse_id: string | null;
    departure_warehouse_id: string;
    departure_address: string | null;
    destination_warehouse_id: string;
    destination_address: string;
    created_by: number;
    owner_id: number;
    status: InventoryStatus;
    total_weight: number;
    total_volume: number;
    special_instructions: string | null;
    created_at: string;
    updated_at: string;
    from_warehouse_id: string;
    to_warehouse_id: string;
    delivery_time: string | null;
    warehouse_id: string;
}