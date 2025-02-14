-- name: CreateWarehouseConnection :one
INSERT INTO warehouse_connections (
    warehouse_id,
    api_key,
    name,
    status,
    created_by
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetWarehouseConnectionById :one
SELECT * FROM warehouse_connections
WHERE connection_id = $1;

-- name: GetWarehouseConnectionByApiKey :one
SELECT * FROM warehouse_connections
WHERE api_key = $1 AND status = 'active';

-- name: ListWarehouseConnections :many
SELECT * FROM warehouse_connections
ORDER BY created_at DESC;

-- name: ListWarehouseConnectionsByWarehouse :many
SELECT * FROM warehouse_connections
WHERE warehouse_id = $1
ORDER BY created_at DESC;

-- name: UpdateWarehouseConnectionStatus :one
UPDATE warehouse_connections
SET 
    status = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE connection_id = $2
RETURNING *;

-- name: UpdateWarehouseConnectionLastUsed :one
UPDATE warehouse_connections
SET 
    last_used_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE connection_id = $1
RETURNING *;

-- name: RevokeWarehouseConnection :one
UPDATE warehouse_connections
SET 
    status = 'revoked',
    updated_at = CURRENT_TIMESTAMP
WHERE connection_id = $1
RETURNING *;

-- name: GetActiveWarehouseConnectionCount :one
SELECT 
    COUNT(*) as total_connections,
    COUNT(CASE WHEN status = 'active' THEN 1 END) as active_connections,
    COUNT(CASE WHEN status = 'inactive' THEN 1 END) as inactive_connections,
    COUNT(CASE WHEN status = 'revoked' THEN 1 END) as revoked_connections
FROM warehouse_connections
WHERE warehouse_id = $1; 