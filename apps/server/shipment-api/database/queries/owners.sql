-- name: CreateShipmentOwner :one
INSERT INTO owners (
    name,
    email,
    phone,
    address,
    account_id
) VALUES (
    @name, @email, @phone, @address, @account_id
) RETURNING *;

-- name: GetShipmentOwnerByAccountId :one
SeLECT * FROM owners
WHERE account_id = @account_id AND deleted_at IS NULL;

-- name: GetShipmentOwnerById :one
SELECT * FROM owners
WHERE id = @id AND deleted_at IS NULL;

-- name: GetShipmentOwnerByEmail :one
SELECT * FROM owners
WHERE email = @email AND deleted_at IS NULL;

-- name: ListShipmentOwners :many
SELECT * FROM owners
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: UpdateShipmentOwner :one
UPDATE owners
SET 
    name = @name,
    email = @email,
    phone = @phone,
    address = @address,
    account_id = COALESCE(@account_id, account_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteShipmentOwner :exec
UPDATE owners
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @id; 