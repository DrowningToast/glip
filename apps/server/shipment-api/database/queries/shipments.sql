-- name: CreateShipment :one
INSERT INTO shipments (
    route,
    last_warehouse_id,
    destination_address,
    carrier_id,
        status,
    total_weight,
    total_volume,
    special_instructions
) VALUES (
    @route, @last_warehouse_id, @destination_address, @carrier_id,  @status, @total_weight, @total_volume, @special_instructions
) RETURNING *;

-- name: GetShipmentById :one
SELECT * FROM shipments
WHERE id = @id;

-- name: ListShipments :many
SELECT * FROM shipments
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentsByLastWarehouse :many
SELECT * FROM shipments
WHERE last_warehouse_id = @warehouse_id
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentsByCarrier :many
SELECT * FROM shipments
WHERE carrier_id = @carrier_id
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: UpdateShipment :one
UPDATE shipments
SET 
    route = COALESCE(@route, route),
    last_warehouse_id = COALESCE(@last_warehouse_id, last_warehouse_id),
    destination_address = COALESCE(@destination_address, destination_address),
    carrier_id = COALESCE(@carrier_id, carrier_id),
    status = COALESCE(@status, status),
    total_weight = COALESCE(@total_weight, total_weight),
    total_volume = COALESCE(@total_volume, total_volume),
    special_instructions = COALESCE(@special_instructions, special_instructions),
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