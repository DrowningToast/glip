export type InventoryStatus = "DELIVERED" | "CANCELLED" | "INCOMING_SHIPMENT" | "WAREHOUSE_RECEIVED" | "WAREHOUSE_DEPARTED"

export type Inventory = {
    id: string;
    shipmentId: string;
    Route: string[];
    LastWarehouseId: string | null;
    DepartureWarehouseId: string;
    DepartureAddress: string | null;
    DestinationWarehouseId: string;
    DestinationAddress: string;
    CreatedBy: number;
    OwnerId: number;
    Status: InventoryStatus;
    TotalWeight: number;
    TotalVolume: number;
    SpecialInstructions: string | null;
    CreatedAt: string;
    UpdatedAt: string;
    FromWarehouseId: string;
    ToWarehouseId: string;
    DeliveryTime: string;
}