-- name: CreateShipmentPerformanceReport :one
INSERT INTO performance_reports (
    period_start,
    period_end,
    warehouse_id,
    total_shipments,
    on_time_delivery_rate,
    average_delivery_time,
    inventory_turnover_rate,
    storage_utilization_rate
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetShipmentReportById :one
SELECT 
    pr.*
FROM performance_reports pr
WHERE pr.report_id = $1;

-- name: ListShipmentWarehouseReports :many
SELECT 
    pr.*
FROM performance_reports pr
WHERE pr.warehouse_id = $1
ORDER BY pr.period_start DESC
LIMIT $2 OFFSET $3;

-- name: GetShipmentWarehousePerformanceByPeriod :many
SELECT 
    pr.*
FROM performance_reports pr
WHERE pr.period_start >= $1
AND pr.period_end <= $2
ORDER BY pr.warehouse_id, pr.period_start;

-- name: GetLatestShipmentWarehousePerformance :one
SELECT 
    pr.*
FROM performance_reports pr
WHERE pr.warehouse_id = $1
ORDER BY pr.period_end DESC
LIMIT 1;

-- name: GetShipmentAveragePerformanceMetrics :one
SELECT 
    warehouse_id,
    AVG(on_time_delivery_rate) as avg_delivery_rate,
    AVG(average_delivery_time) as avg_delivery_time,
    AVG(inventory_turnover_rate) as avg_turnover_rate,
    AVG(storage_utilization_rate) as avg_utilization_rate
FROM performance_reports
WHERE warehouse_id = $1
AND period_start >= $2
AND period_end <= $3
GROUP BY warehouse_id; 