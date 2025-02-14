-- name: CreateShipmentAccount :one
INSERT INTO accounts (
    username,
    password,
    role
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetShipmentAccountByUsername :one
SELECT * FROM accounts
WHERE username = $1 AND deleted_at IS NULL;

-- name: GetShipmentAccountById :one
SELECT * FROM accounts
WHERE account_id = $1 AND deleted_at IS NULL;

-- name: ListShipmentAccounts :many
SELECT * FROM accounts
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateShipmentAccount :one
UPDATE accounts
SET 
    username = COALESCE($1, username),
    password = COALESCE($2, password),
    role = COALESCE($3, role),
    updated_at = CURRENT_TIMESTAMP
WHERE account_id = $4 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteShipmentAccount :exec
UPDATE accounts
SET deleted_at = CURRENT_TIMESTAMP
WHERE account_id = $1; 