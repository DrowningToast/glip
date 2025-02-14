-- name: CreateStockTransaction :one
INSERT INTO stock_transactions (
    inventory_id,
    transaction_type,
    carrier_id,
    quantity,
    previous_quantity,
    current_quantity,
    reference_id,
    reason,
    staff_performed_id,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetStockTransactionByID :one
SELECT * FROM stock_transactions
WHERE transaction_id = $1;

-- name: ListStockTransactions :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    o.name as owner_name,
    s.name as staff_name,
    c.name as carrier_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.inventory_id
JOIN owners o ON i.owner_id = o.owner_id
JOIN staffs s ON st.staff_performed_id = s.staff_id
LEFT JOIN carriers c ON st.carrier_id = c.carrier_id
ORDER BY st.created_at DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: ListStockTransactionsByInventory :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    s.name as staff_name,
    c.name as carrier_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.inventory_id
JOIN staffs s ON st.staff_performed_id = s.staff_id
LEFT JOIN carriers c ON st.carrier_id = c.carrier_id
WHERE st.inventory_id = $1
ORDER BY st.created_at DESC;

-- name: ListStockTransactionsByOwner :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    s.name as staff_name,
    c.name as carrier_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.inventory_id
JOIN staffs s ON st.staff_performed_id = s.staff_id
LEFT JOIN carriers c ON st.carrier_id = c.carrier_id
WHERE i.owner_id = $1
ORDER BY st.created_at DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: ListStockTransactionsByWarehouse :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    o.name as owner_name,
    s.name as staff_name,
    c.name as carrier_name,
    w.name as warehouse_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.inventory_id
JOIN owners o ON i.owner_id = o.owner_id
JOIN staffs s ON st.staff_performed_id = s.staff_id
JOIN storage_locations sl ON i.storage_location_id = sl.storage_location_id
JOIN warehouses w ON sl.warehouse_id = w.warehouse_id
LEFT JOIN carriers c ON st.carrier_id = c.carrier_id
WHERE sl.warehouse_id = $1
ORDER BY st.created_at DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: GetTransactionsByDateRange :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    o.name as owner_name,
    s.name as staff_name,
    c.name as carrier_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.inventory_id
JOIN owners o ON i.owner_id = o.owner_id
JOIN staffs s ON st.staff_performed_id = s.staff_id
LEFT JOIN carriers c ON st.carrier_id = c.carrier_id
WHERE st.created_at BETWEEN sqlc.arg('start_date') AND sqlc.arg('end_date')
ORDER BY st.created_at DESC; 