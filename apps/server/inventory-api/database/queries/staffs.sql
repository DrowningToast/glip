-- name: CreateStaff :one
INSERT INTO staffs (
    name,
    email,
    phone,
    address,
    account_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStaffByID :one
SELECT * FROM staffs
WHERE staff_id = $1 AND deleted_at IS NULL;

-- name: GetStaffByEmail :one
SELECT * FROM staffs
WHERE email = $1 AND deleted_at IS NULL;

-- name: GetStaffWithAccount :one
SELECT 
    s.*,
    a.username,
    a.role
FROM staffs s
JOIN accounts a ON s.account_id = a.account_id
WHERE s.staff_id = $1 AND s.deleted_at IS NULL;

-- name: ListStaffs :many
SELECT * FROM staffs
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: ListStaffsByRole :many
SELECT s.* FROM staffs s
JOIN accounts a ON s.account_id = a.account_id
WHERE a.role = $1 AND s.deleted_at IS NULL
ORDER BY s.created_at DESC;

-- name: UpdateStaff :one
UPDATE staffs
SET 
    name = COALESCE(sqlc.narg('name'), name),
    email = COALESCE(sqlc.narg('email'), email),
    phone = COALESCE(sqlc.narg('phone'), phone),
    address = COALESCE(sqlc.narg('address'), address),
    updated_at = CURRENT_TIMESTAMP
WHERE staff_id = sqlc.arg('staff_id') AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteStaff :exec
UPDATE staffs
SET deleted_at = CURRENT_TIMESTAMP
WHERE staff_id = $1; 