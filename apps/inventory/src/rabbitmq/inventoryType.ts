export type ShipmentQueueType = "INBOUND" | "OUTBOUND";

export type ShipmentStatusRabbitMQType = "WAITING_FOR_PICKUP_TO_WAREHOUSE" | "IN_TRANSIT_ON_THE_WAY" | "ARRIVED_AT_WAREHOUSE" | "DELIVERED" | "CANCELLED";

export type InventoryRabbitMQType = {
    id: number;
    route: string[];
    last_warehouse_id: string | null;
    departure_warehouse_id: string;
    departure_address: string | null;
    destination_warehouse_id: string;
    destination_address: string;
    created_by: number;
    owner_id: number;
    status: ShipmentStatusRabbitMQType;
    total_weight: number;
    total_volume: number;
    special_instructions: string | null;
    created_at: string;
    updated_at: string;
    from_warehouse_id?: string;
    to_warehouse_id?: string;
    type: ShipmentQueueType;
}