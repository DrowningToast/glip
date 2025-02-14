-- name: CreateShipmentAccount :one
INSERT INTO accounts (
    username,
    password,
    role
) VALUES (
    @username, @password, @role
) RETURNING *;

-- name: GetShipmentAccountByUsername :one
SELECT * FROM accounts
WHERE username = @username AND deleted_at IS NULL;

-- name: GetShipmentAccountById :one
SELECT * FROM accounts
WHERE id = @id AND deleted_at IS NULL;

-- name: ListShipmentAccounts :many
SELECT * FROM accounts
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateShipmentAccount :one
UPDATE accounts
SET 
    username = COALESCE(@username, username),
    password = COALESCE(@password, password),
    role = COALESCE(@role, role),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteShipmentAccount :exec
UPDATE accounts
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @id; 