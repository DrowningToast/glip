-- name: CreateShipmentOwner :one
INSERT INTO owners (
    name,
    email,
    phone,
    address
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetShipmentOwnerById :one
SELECT * FROM owners
WHERE owner_id = $1 AND deleted_at IS NULL;

-- name: GetShipmentOwnerByEmail :one
SELECT * FROM owners
WHERE email = $1 AND deleted_at IS NULL;

-- name: ListShipmentOwners :many
SELECT * FROM owners
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateShipmentOwner :one
UPDATE owners
SET 
    name = $1,
    email = $2,
    phone = $3,
    address = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE owner_id = $5 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteShipmentOwner :exec
UPDATE owners
SET deleted_at = CURRENT_TIMESTAMP
WHERE owner_id = $1; 