export type ShipmentQueueType = "INBOUND" | "OUTBOUND";

export type ShipmentStatusRabbitMQType = "WAITING_FOR_PICKUP_TO_WAREHOUSE" | "IN_TRANSIT_ON_THE_WAY" | "ARRIVED_AT_WAREHOUSE" | "DELIVERED" | "CANCELLED";

export type InventoryRabbitMQType = {
    shipmentId: string;
    Route: string[];
    LastWarehouseId: string | null;
    DepartureWarehouseId: string;
    DepartureAddress: string | null;
    DestinationWarehouseId: string;
    DestinationAddress: string;
    CreatedBy: number;
    OwnerId: number;
    Status: ShipmentStatusRabbitMQType;
    TotalWeight: number;
    TotalVolume: number;
    SpecialInstructions: string | null;
    CreatedAt: string;
    UpdatedAt: string;
    FromWarehouseId: string;
    ToWarehouseId: string;
    QueueType: ShipmentQueueType;
}