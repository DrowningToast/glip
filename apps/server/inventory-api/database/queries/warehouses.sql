-- name: CreateWarehouse :one
INSERT INTO warehouses (
    name,
    location_address,
    total_capacity,
    current_capacity,
    description,
    status
) VALUES (
    @name, @locationAddress, @totalCapacity, 
    @currentCapacity, @description, @status
) RETURNING *;

-- name: GetWarehouseById :one
SELECT * FROM warehouses
WHERE id = @id;

-- name: ListWarehouses :many
SELECT * FROM warehouses
ORDER BY created_at DESC;

-- name: ListWarehousesByStatus :many
SELECT * FROM warehouses
WHERE status = @status
ORDER BY name;

-- name: UpdateWarehouse :one
UPDATE warehouses
SET 
    name = @name,
    location_address = @locationAddress,
    total_capacity = @totalCapacity,
    current_capacity = @currentCapacity,
    description = @description,
    status = @status,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: GetWarehouseStorageStats :one
SELECT 
    w.*,
    COUNT(sl.id) as total_storage_locations,
    SUM(sl.capacity) as total_storage_capacity,
    SUM(sl.current_occupancy) as total_current_occupancy,
    COUNT(DISTINCT i.owner_id) as total_owners,
    COUNT(i.id) as total_inventory_items
FROM warehouses w
LEFT JOIN storage_locations sl ON w.id = sl.warehouse_id
LEFT JOIN inventory i ON sl.id = i.storage_location_id
WHERE w.id = @id
GROUP BY w.id;

-- name: GetWarehouseStorageLocations :many
SELECT 
    sl.*,
    COUNT(i.id) as total_items,
    SUM(i.quantity) as total_quantity
FROM storage_locations sl
LEFT JOIN inventory i ON sl.id = i.storage_location_id
WHERE sl.warehouse_id = @warehouseId
GROUP BY sl.id
ORDER BY sl.area_name;

-- name: GetWarehouseInventoryItems :many
SELECT 
    i.*,
    sl.area_name as storage_location_name,
    o.name as owner_name
FROM inventory i
JOIN storage_locations sl ON i.storage_location_id = sl.id
LEFT JOIN owners o ON i.owner_id = o.id
WHERE sl.warehouse_id = @warehouseId
ORDER BY i.created_at DESC;

-- name: UpdateWarehouseCapacity :one
UPDATE warehouses
SET 
    current_capacity = current_capacity + @capacityChange,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: CreateStorageLocation :one
INSERT INTO storage_locations (
    warehouse_id,
    area_name,
    capacity,
    current_occupancy
) VALUES (
    @warehouseId, @areaName, @capacity, @currentOccupancy
) RETURNING *;

-- name: GetStorageLocationById :one
SELECT * FROM storage_locations
WHERE id = @id;

-- name: ListStorageLocations :many
SELECT * FROM storage_locations
WHERE warehouse_id = @warehouseId
ORDER BY area_name;

-- name: UpdateStorageLocation :one
UPDATE storage_locations
SET 
    area_name = COALESCE(@areaName, area_name),
    capacity = COALESCE(@capacity, capacity),
    current_occupancy = COALESCE(@currentOccupancy, current_occupancy),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: UpdateStorageOccupancy :one
UPDATE storage_locations
SET 
    current_occupancy = current_occupancy + @occupancyChange,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: GetAvailableStorageLocations :many
SELECT * FROM storage_locations
WHERE warehouse_id = @warehouseId 
AND (capacity - current_occupancy) >= @requiredSpace
ORDER BY (capacity - current_occupancy) DESC; 