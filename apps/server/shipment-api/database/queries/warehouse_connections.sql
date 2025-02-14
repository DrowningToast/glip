-- name: CreateWarehouseConnection :one
INSERT INTO warehouse_connections (
    warehouse_id,
    api_key,
    name,
    status,
    created_by
) VALUES (
    @warehouseId, @apiKey, @name, @status, @createdBy
) RETURNING *;

-- name: GetWarehouseConnectionById :one
SELECT * FROM warehouse_connections
WHERE id = @id;

-- name: GetWarehouseConnectionByApiKey :one
SELECT * FROM warehouse_connections
WHERE api_key = @apiKey AND status = 'ACTIVE';

-- name: ListWarehouseConnections :many
SELECT * FROM warehouse_connections
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: ListWarehouseConnectionsByStatus :many
SELECT * FROM warehouse_connections
WHERE status = @status
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: UpdateWarehouseConnection :one
UPDATE warehouse_connections
SET 
    status = COALESCE(@status, status),
    updated_at = CURRENT_TIMESTAMP,
    warehouse_id = COALESCE(@warehouseId, warehouse_id),
    api_key = COALESCE(@apiKey, api_key), 
    name = COALESCE(@name, name),
    status = COALESCE(@status, status),
    last_used_at = COALESCE(@lastUsedAt, last_used_at),
    created_by = COALESCE(@createdBy, created_by)
WHERE id = @id
RETURNING *;

-- name: RevokeWarehouseConnection :one
UPDATE warehouse_connections
SET 
    status = 'REVOKED',
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;
