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
    @inventoryId, @transactionType, @carrierId,
    @quantity, @previousQuantity, @currentQuantity,
    @referenceId, @reason, @staffPerformedId, @meta
) RETURNING *;

-- name: GetStockTransactionById :one
SELECT * FROM stock_transactions
WHERE id = @id;

-- name: ListStockTransactions :many
SELECT * FROM stock_transactions
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: ListStockTransactionsByInventory :many
SELECT * FROM stock_transactions
WHERE inventory_id = @inventoryId
ORDER BY created_at DESC;

-- name: ListStockTransactionsByType :many
SELECT * FROM stock_transactions
WHERE transaction_type = @transactionType
ORDER BY created_at DESC;

-- name: ListStockTransactionsByStaff :many
SELECT * FROM stock_transactions
WHERE staff_performed_id = @staffId
ORDER BY created_at DESC;

-- name: GetStockTransactionStats :one
SELECT 
    COUNT(*) as total_transactions,
    COUNT(CASE WHEN transaction_type = 'IN' THEN 1 END) as total_in,
    COUNT(CASE WHEN transaction_type = 'OUT' THEN 1 END) as total_out,
    COUNT(CASE WHEN transaction_type = 'TRANSFER' THEN 1 END) as total_transfers,
    SUM(CASE WHEN transaction_type = 'IN' THEN quantity ELSE 0 END) as total_in_quantity,
    SUM(CASE WHEN transaction_type = 'OUT' THEN quantity ELSE 0 END) as total_out_quantity
FROM stock_transactions
WHERE inventory_id = @inventoryId;

-- name: ListStockTransactionsByOwner :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    s.name as staff_name,
    c.name as carrier_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.id
JOIN staffs s ON st.staff_performed_id = s.id
LEFT JOIN carriers c ON st.carrier_id = c.id
WHERE i.owner_id = @ownerId
ORDER BY st.created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

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
JOIN inventory i ON st.inventory_id = i.id
JOIN owners o ON i.owner_id = o.id
JOIN staffs s ON st.staff_performed_id = s.id
JOIN storage_locations sl ON i.storage_location_id = sl.id
JOIN warehouses w ON sl.warehouse_id = w.id
LEFT JOIN carriers c ON st.carrier_id = c.id
WHERE sl.warehouse_id = @warehouseId
ORDER BY st.created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: GetTransactionsByDateRange :many
SELECT 
    st.*,
    i.sku,
    i.name as inventory_name,
    o.name as owner_name,
    s.name as staff_name,
    c.name as carrier_name
FROM stock_transactions st
JOIN inventory i ON st.inventory_id = i.id
JOIN owners o ON i.owner_id = o.id
JOIN staffs s ON st.staff_performed_id = s.id
LEFT JOIN carriers c ON st.carrier_id = c.id
WHERE st.created_at BETWEEN @startDate AND @endDate
ORDER BY st.created_at DESC; 