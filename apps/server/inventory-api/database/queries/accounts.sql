-- name: CreateAccount :one
INSERT INTO accounts (
    username,
    password,
    role
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccountByUsername :one
SELECT * FROM accounts
WHERE username = $1 AND deleted_at IS NULL;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE account_id = $1 AND deleted_at IS NULL;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateAccount :one
UPDATE accounts
SET 
    username = COALESCE(sqlc.narg('username'), username),
    password = COALESCE(sqlc.narg('password'), password),
    role = COALESCE(sqlc.narg('role'), role),
    updated_at = CURRENT_TIMESTAMP
WHERE account_id = sqlc.arg('account_id') AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteAccount :exec
UPDATE accounts
SET deleted_at = CURRENT_TIMESTAMP
WHERE account_id = $1;
