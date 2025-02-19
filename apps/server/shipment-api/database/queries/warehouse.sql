-- name: CreateWarehouse :one
INSERT INTO warehouses (
    name,
    location_address,
    country,
    city,
    total_capacity,
    current_capacity,
    description,
    status
) VALUES (
    @name, @location_address, @country, @city,
    @total_capacity, @current_capacity, @description, @status
) RETURNING *;

-- name: GetWarehouse :one
SELECT * FROM warehouses
WHERE id = @id;

-- name: ListWarehouses :many
SELECT * FROM warehouses
ORDER BY created_at DESC;

-- name: UpdateWarehouse :one
UPDATE warehouses
SET
    name = COALESCE(@name, name),
    location_address = COALESCE(@location_address, location_address),
    country = COALESCE(@country, country),
    city = COALESCE(@city, city),
    total_capacity = COALESCE(@total_capacity, total_capacity),
    current_capacity = COALESCE(@current_capacity, current_capacity),
    description = COALESCE(@description, description),
    status = COALESCE(@status, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteWarehouse :exec
UPDATE warehouses
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @id;

-- name: GetWarehouseByStatus :many
SELECT * FROM warehouses
WHERE status = @status;

-- name: UpdateWarehouseCapacity :one
UPDATE warehouses
SET current_capacity = @current_capacity,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: GetWarehousesByCountry :many
SELECT * FROM warehouses
WHERE country = @country
ORDER BY city, name;

-- name: GetWarehousesByCity :many
SELECT * FROM warehouses
WHERE city = @city
ORDER BY name;

