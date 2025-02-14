-- name: CreateInventory :one
INSERT INTO inventory (
    sku,
    name,
    description,
    category,
    subcategory,
    owner_id,
    storage_location_id,
    quantity,
    weight,
    dimensions,
    status
) VALUES (
    @sku, @name, @description, @category, @subcategory,
    @ownerId, @storageLocationId, @quantity, @weight,
    @dimensions, @status
) RETURNING *;

-- name: GetInventoryById :one
SELECT * FROM inventory
WHERE id = @id AND removed_at IS NULL;

-- name: ListInventory :many
SELECT * FROM inventory
WHERE removed_at IS NULL
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: ListInventoryByOwner :many
SELECT * FROM inventory
WHERE owner_id = @ownerId AND removed_at IS NULL
ORDER BY created_at DESC;

-- name: ListInventoryByStorageLocation :many
SELECT * FROM inventory
WHERE storage_location_id = @storageLocationId AND removed_at IS NULL
ORDER BY created_at DESC;

-- name: ListInventoryByStatus :many
SELECT * FROM inventory
WHERE status = @status AND removed_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateInventory :one
UPDATE inventory
SET 
    name = COALESCE(@name, name),
    description = COALESCE(@description, description),
    category = COALESCE(@category, category),
    subcategory = COALESCE(@subcategory, subcategory),
    storage_location_id = COALESCE(@storageLocationId, storage_location_id),
    quantity = COALESCE(@quantity, quantity),
    weight = COALESCE(@weight, weight),
    dimensions = COALESCE(@dimensions, dimensions),
    status = COALESCE(@status, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND removed_at IS NULL
RETURNING *;

-- name: UpdateInventoryQuantity :one
UPDATE inventory
SET 
    quantity = @quantity,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND removed_at IS NULL
RETURNING *;

-- name: UpdateInventoryStatus :one
UPDATE inventory
SET 
    status = @status,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND removed_at IS NULL
RETURNING *;

-- name: SoftDeleteInventory :exec
UPDATE inventory
SET 
    removed_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id;

-- name: GetInventoryStats :one
SELECT 
    COUNT(*) as total_items,
    COUNT(DISTINCT owner_id) as total_owners,
    COUNT(DISTINCT storage_location_id) as total_locations,
    SUM(quantity) as total_quantity,
    SUM(CASE WHEN status = 'available' THEN quantity ELSE 0 END) as available_quantity,
    SUM(CASE WHEN status = 'reserved' THEN quantity ELSE 0 END) as reserved_quantity,
    SUM(CASE WHEN status = 'damaged' THEN quantity ELSE 0 END) as damaged_quantity
FROM inventory
WHERE removed_at IS NULL;

-- name: GetInventoryBySKU :many
SELECT 
    i.*,
    o.name as owner_name,
    sl.area_name as storage_area,
    w.name as warehouse_name
FROM inventory i
JOIN owners o ON i.owner_id = o.owner_id
JOIN storage_locations sl ON i.storage_location_id = sl.storage_location_id
JOIN warehouses w ON sl.warehouse_id = w.warehouse_id
WHERE i.sku = $1 AND i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: GetInventoryBySKUAndOwner :many
SELECT 
    i.*,
    sl.area_name as storage_area,
    w.name as warehouse_name
FROM inventory i
JOIN storage_locations sl ON i.storage_location_id = sl.id
JOIN warehouses w ON sl.warehouse_id = w.id
WHERE i.sku = @sku AND i.owner_id = @ownerId AND i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: ListInventoryByWarehouse :many
SELECT 
    i.*,
    o.name as owner_name,
    sl.area_name as storage_area
FROM inventory i
JOIN owners o ON i.owner_id = o.id
JOIN storage_locations sl ON i.storage_location_id = sl.id
WHERE sl.warehouse_id = @warehouseId AND i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: GetTotalQuantityBySKUAndOwner :one
SELECT 
    SUM(quantity) as total_quantity
FROM inventory
WHERE sku = @sku AND owner_id = @ownerId AND removed_at IS NULL;

-- name: GetInventoryByLocation :many
SELECT 
    i.*,
    o.name as owner_name
FROM inventory i
JOIN owners o ON i.owner_id = o.id
WHERE i.storage_location_id = @storageLocationId AND i.removed_at IS NULL
ORDER BY i.sku, o.name; 