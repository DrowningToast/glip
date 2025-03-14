-- name: CreateShipment :one
INSERT INTO shipments (
    route,
    last_warehouse_id,
    departure_warehouse_id,
    departure_address,
    destination_warehouse_id,
    destination_address,
    owner_id,
    created_by,
    status,
    total_weight,
    total_volume,
    special_instructions
) VALUES (
    @route, @last_warehouse_id, @departure_warehouse_id, @departure_address, @destination_warehouse_id, @destination_address, @owner_id, @created_by, @status, @total_weight, @total_volume, @special_instructions
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

-- name: ListShipmentsByStatus :many
SELECT * FROM shipments
WHERE status = @status
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentsByStatusAndLastWarehouse :many
SELECT * FROM shipments
WHERE status = @status AND last_warehouse_id = @warehouse_id
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);


-- name: ListShipmentsByAccountUsername :many
SELECT * FROM shipments
JOIN accounts ON owners.account_id = accounts.id
WHERE accounts.username = @username AND status = COALESCE(sqlc.narg(status), status)
ORDER BY shipments.created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: UpdateShipment :one
UPDATE shipments
SET 
    route = COALESCE(@route, route),
    last_warehouse_id = COALESCE(@last_warehouse_id, last_warehouse_id),
    destination_address = COALESCE(@destination_address, destination_address),
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