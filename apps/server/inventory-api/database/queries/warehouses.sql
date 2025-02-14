-- name: CreateWarehouse :one
INSERT INTO warehouses (
    name,
    location_address,
    total_capacity,
    current_capacity,
    description,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetWarehouseByID :one
SELECT * FROM warehouses
WHERE warehouse_id = $1;

-- name: ListWarehouses :many
SELECT * FROM warehouses
ORDER BY created_at DESC;

-- name: UpdateWarehouse :one
UPDATE warehouses
SET 
    name = COALESCE(sqlc.narg('name'), name),
    location_address = COALESCE(sqlc.narg('location_address'), location_address),
    total_capacity = COALESCE(sqlc.narg('total_capacity'), total_capacity),
    current_capacity = COALESCE(sqlc.narg('current_capacity'), current_capacity),
    description = COALESCE(sqlc.narg('description'), description),
    status = COALESCE(sqlc.narg('status'), status),
    updated_at = CURRENT_TIMESTAMP
WHERE warehouse_id = sqlc.arg('warehouse_id')
RETURNING *;

-- name: UpdateWarehouseCapacity :one
UPDATE warehouses
SET 
    current_capacity = current_capacity + sqlc.arg('capacity_change'),
    updated_at = CURRENT_TIMESTAMP
WHERE warehouse_id = sqlc.arg('warehouse_id')
RETURNING *;

-- name: CreateStorageLocation :one
INSERT INTO storage_locations (
    warehouse_id,
    area_name,
    capacity,
    current_occupancy
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetStorageLocationByID :one
SELECT * FROM storage_locations
WHERE storage_location_id = $1;

-- name: ListStorageLocations :many
SELECT * FROM storage_locations
WHERE warehouse_id = $1
ORDER BY area_name;

-- name: UpdateStorageLocation :one
UPDATE storage_locations
SET 
    area_name = COALESCE(sqlc.narg('area_name'), area_name),
    capacity = COALESCE(sqlc.narg('capacity'), capacity),
    current_occupancy = COALESCE(sqlc.narg('current_occupancy'), current_occupancy),
    updated_at = CURRENT_TIMESTAMP
WHERE storage_location_id = sqlc.arg('storage_location_id')
RETURNING *;

-- name: UpdateStorageOccupancy :one
UPDATE storage_locations
SET 
    current_occupancy = current_occupancy + sqlc.arg('occupancy_change'),
    updated_at = CURRENT_TIMESTAMP
WHERE storage_location_id = sqlc.arg('storage_location_id')
RETURNING *;

-- name: GetAvailableStorageLocations :many
SELECT * FROM storage_locations
WHERE warehouse_id = $1 
AND (capacity - current_occupancy) >= sqlc.arg('required_space')
ORDER BY (capacity - current_occupancy) DESC; 