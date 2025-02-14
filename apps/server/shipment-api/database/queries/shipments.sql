-- name: CreateShipment :one
INSERT INTO shipments (
    route,
    last_warehouse_id,
    destination_address,
    carrier_id,
    scheduled_departure,
    scheduled_arrival,
    actual_departure,
    actual_arrival,
    status,
    total_weight,
    total_volume,
    special_instructions
) VALUES (
    @route, @lastWarehouseId, @destinationAddress, @carrierId,
    @scheduledDeparture, @scheduledArrival, @actualDeparture, @actualArrival,
    @status, @totalWeight, @totalVolume, @specialInstructions
) RETURNING *;

-- name: GetShipmentById :one
SELECT * FROM shipments
WHERE id = @id;

-- name: ListShipments :many
SELECT * FROM shipments
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: ListShipmentsByLastWarehouse :many
SELECT * FROM shipments
WHERE last_warehouse_id = @warehouseId
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: ListShipmentsByCarrier :many
SELECT * FROM shipments
WHERE carrier_id = @carrierId
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: UpdateShipment :one
UPDATE shipments
SET 
    route = COALESCE(@route, route),
    last_warehouse_id = COALESCE(@lastWarehouseId, last_warehouse_id),
    destination_address = COALESCE(@destinationAddress, destination_address),
    carrier_id = COALESCE(@carrierId, carrier_id),
    scheduled_departure = COALESCE(@scheduledDeparture, scheduled_departure),
    scheduled_arrival = COALESCE(@scheduledArrival, scheduled_arrival),
    actual_departure = COALESCE(@actualDeparture, actual_departure),
    actual_arrival = COALESCE(@actualArrival, actual_arrival),
    status = COALESCE(@status, status),
    total_weight = COALESCE(@totalWeight, total_weight),
    total_volume = COALESCE(@totalVolume, total_volume),
    special_instructions = COALESCE(@specialInstructions, special_instructions),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: UpdateShipmentStatus :one
UPDATE shipments
SET 
    status = @status,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *; 