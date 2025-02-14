-- name: CreateShipment :one
INSERT INTO shipments (
    origin_warehouse_id,
    destination_address,
    carrier_id,
    scheduled_departure,
    scheduled_arrival,
    status,
    total_weight,
    total_volume,
    special_instructions
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetShipmentById :one
SELECT * FROM shipments
WHERE shipment_id = $1;

-- name: ListShipments :many
SELECT * FROM shipments
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListShipmentsByWarehouse :many
SELECT * FROM shipments
WHERE origin_warehouse_id = $1
ORDER BY created_at DESC;

-- name: ListShipmentsByCarrier :many
SELECT * FROM shipments
WHERE carrier_id = $1
ORDER BY created_at DESC;

-- name: UpdateShipment :one
UPDATE shipments
SET 
    destination_address = COALESCE($1, destination_address),
    carrier_id = COALESCE($2, carrier_id),
    scheduled_departure = COALESCE($3, scheduled_departure),
    scheduled_arrival = COALESCE($4, scheduled_arrival),
    actual_departure = COALESCE($5, actual_departure),
    actual_arrival = COALESCE($6, actual_arrival),
    status = COALESCE($7, status),
    total_weight = COALESCE($8, total_weight),
    total_volume = COALESCE($9, total_volume),
    special_instructions = COALESCE($10, special_instructions),
    updated_at = CURRENT_TIMESTAMP
WHERE shipment_id = $11
RETURNING *;

-- name: UpdateShipmentStatus :one
UPDATE shipments
SET 
    status = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE shipment_id = $2
RETURNING *; 