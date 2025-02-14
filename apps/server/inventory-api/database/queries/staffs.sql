-- name: CreateStaff :one
INSERT INTO staffs (
    name,
    email,
    phone,
    address,
    account_id
) VALUES (
    @name, @email, @phone, @address, @accountId
) RETURNING *;

-- name: GetStaffById :one
SELECT * FROM staffs
WHERE id = @id AND deleted_at IS NULL;

-- name: GetStaffByEmail :one
SELECT * FROM staffs
WHERE email = @email AND deleted_at IS NULL;

-- name: GetStaffWithAccount :one
SELECT 
    s.*,
    a.username,
    a.role
FROM staffs s
JOIN accounts a ON s.account_id = a.id
WHERE s.id = @id AND s.deleted_at IS NULL;

-- name: ListStaffs :many
SELECT * FROM staffs
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: ListStaffsByRole :many
SELECT s.* FROM staffs s
JOIN accounts a ON s.account_id = a.id
WHERE a.role = @role AND s.deleted_at IS NULL
ORDER BY s.created_at DESC;

-- name: UpdateStaff :one
UPDATE staffs
SET 
    name = @name,
    email = @email,
    phone = @phone,
    address = @address,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteStaff :exec
UPDATE staffs
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @id; 