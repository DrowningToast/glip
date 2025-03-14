-- name: CreateAccount :one
INSERT INTO accounts (
    username,
    password,
    email
) VALUES (
    @username, @password, @email
) RETURNING *;

-- name: GetAccountByUsername :one
SELECT * FROM accounts
WHERE username = @username AND deleted_at IS NULL;

-- name: GetAccountByEmail :one
SELECT * FROM accounts
WHERE email = @email AND deleted_at IS NULL;

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = @id AND deleted_at IS NULL;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: UpdateAccount :one
UPDATE accounts
SET 
    username = COALESCE(sqlc.narg(username), username),
    password = COALESCE(sqlc.narg(password), password),
    email = COALESCE(sqlc.narg(email), email),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteAccount :exec
UPDATE accounts
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @id;