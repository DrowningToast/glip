-- name: CreateAccount :one
INSERT INTO accounts (
    username,
    password,
    role
) VALUES (
    @username, @password, @role
) RETURNING *;

-- name: GetAccountByUsername :one
SELECT * FROM accounts
WHERE username = @username AND deleted_at IS NULL;

-- name: GetAccountByUserId :one
SELECT * FROM accounts
WHERE user_id = @user_id AND deleted_at IS NULL;

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
    username = COALESCE(@username, username),
    password = COALESCE(@password, password),
    role = COALESCE(@role, role),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteAccount :exec
UPDATE accounts
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @id;