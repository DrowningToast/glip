-- name: CreateWarehouseConnection :one
INSERT INTO warehouse_connections (
    warehouse_id,
    api_key,
    name,
    status,
    created_by
) VALUES (
    @warehouse_id, @api_key, @name, @status, @created_by
) RETURNING *;

-- name: GetWarehouseConnectionById :one
SELECT * FROM warehouse_connections
WHERE id = @id;

-- name: GetWarehouseConnectionByApiKey :one
SELECT * FROM warehouse_connections
WHERE api_key = @api_key AND status = 'ACTIVE';

-- name: ListWarehouseConnections :many
SELECT * FROM warehouse_connections
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListWarehouseConnectionsByStatus :many
SELECT * FROM warehouse_connections
WHERE status = @status
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: UpdateWarehouseConnection :one
UPDATE warehouse_connections
SET 
    status = COALESCE(@status, status),
    updated_at = CURRENT_TIMESTAMP,
    warehouse_id = COALESCE(@warehouse_id, warehouse_id),
    api_key = COALESCE(@api_key, api_key), 
    name = COALESCE(@name, name),
    status = COALESCE(@status, status),
    last_used_at = COALESCE(@last_used_at, last_used_at),
    created_by = COALESCE(@created_by, created_by)
WHERE id = @id
RETURNING *;

-- name: RevokeWarehouseConnection :one
UPDATE warehouse_connections
SET 
    status = 'REVOKED',
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;
