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
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetInventoryByID :one
SELECT * FROM inventory
WHERE inventory_id = $1 AND removed_at IS NULL;

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
JOIN storage_locations sl ON i.storage_location_id = sl.storage_location_id
JOIN warehouses w ON sl.warehouse_id = w.warehouse_id
WHERE i.sku = $1 AND i.owner_id = $2 AND i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: ListInventory :many
SELECT 
    i.*,
    o.name as owner_name,
    sl.area_name as storage_area,
    w.name as warehouse_name
FROM inventory i
JOIN owners o ON i.owner_id = o.owner_id
JOIN storage_locations sl ON i.storage_location_id = sl.storage_location_id
JOIN warehouses w ON sl.warehouse_id = w.warehouse_id
WHERE i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: ListInventoryByOwner :many
SELECT 
    i.*,
    sl.area_name as storage_area,
    w.name as warehouse_name
FROM inventory i
JOIN storage_locations sl ON i.storage_location_id = sl.storage_location_id
JOIN warehouses w ON sl.warehouse_id = w.warehouse_id
WHERE i.owner_id = $1 AND i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: ListInventoryByWarehouse :many
SELECT 
    i.*,
    o.name as owner_name,
    sl.area_name as storage_area
FROM inventory i
JOIN owners o ON i.owner_id = o.owner_id
JOIN storage_locations sl ON i.storage_location_id = sl.storage_location_id
WHERE sl.warehouse_id = $1 AND i.removed_at IS NULL
ORDER BY i.created_at DESC;

-- name: UpdateInventory :one
UPDATE inventory
SET 
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    category = COALESCE(sqlc.narg('category'), category),
    subcategory = COALESCE(sqlc.narg('subcategory'), subcategory),
    storage_location_id = COALESCE(sqlc.narg('storage_location_id'), storage_location_id),
    weight = COALESCE(sqlc.narg('weight'), weight),
    dimensions = COALESCE(sqlc.narg('dimensions'), dimensions),
    status = COALESCE(sqlc.narg('status'), status),
    updated_at = CURRENT_TIMESTAMP
WHERE inventory_id = sqlc.arg('inventory_id') AND removed_at IS NULL
RETURNING *;

-- name: UpdateInventoryQuantity :one
UPDATE inventory
SET 
    quantity = quantity + sqlc.arg('quantity_change'),
    updated_at = CURRENT_TIMESTAMP
WHERE inventory_id = sqlc.arg('inventory_id') AND removed_at IS NULL
RETURNING *;

-- name: SoftDeleteInventory :exec
UPDATE inventory
SET removed_at = CURRENT_TIMESTAMP
WHERE inventory_id = $1;

-- name: GetTotalQuantityBySKUAndOwner :one
SELECT 
    SUM(quantity) as total_quantity
FROM inventory
WHERE sku = $1 AND owner_id = $2 AND removed_at IS NULL;

-- name: GetInventoryByLocation :many
SELECT 
    i.*,
    o.name as owner_name
FROM inventory i
JOIN owners o ON i.owner_id = o.owner_id
WHERE i.storage_location_id = $1 AND i.removed_at IS NULL
ORDER BY i.sku, o.name; 