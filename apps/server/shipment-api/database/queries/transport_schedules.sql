-- name: CreateShipmentTransportSchedule :one
INSERT INTO transport_schedules (
    shipment_id,
    planned_route,
    planned_departure,
    planned_arrival,
    estimated_duration,
    route_status
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetShipmentTransportScheduleById :one
SELECT 
    ts.*,
    s.origin_warehouse_id,
    s.destination_address,
    s.carrier_id,
    c.name as carrier_name
FROM transport_schedules ts
JOIN shipments s ON ts.shipment_id = s.shipment_id
LEFT JOIN carriers c ON s.carrier_id = c.carrier_id
WHERE ts.schedule_id = $1;

-- name: ListShipmentTransportSchedules :many
SELECT 
    ts.*,
    s.origin_warehouse_id,
    s.destination_address,
    c.name as carrier_name
FROM transport_schedules ts
JOIN shipments s ON ts.shipment_id = s.shipment_id
LEFT JOIN carriers c ON s.carrier_id = c.carrier_id
ORDER BY ts.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListShipmentSchedulesByWarehouse :many
SELECT 
    ts.*,
    s.destination_address,
    c.name as carrier_name
FROM transport_schedules ts
JOIN shipments s ON ts.shipment_id = s.shipment_id
LEFT JOIN carriers c ON s.carrier_id = c.carrier_id
WHERE s.origin_warehouse_id = $1
ORDER BY ts.planned_departure;

-- name: ListShipmentSchedulesByStatus :many
SELECT 
    ts.*,
    s.origin_warehouse_id,
    s.destination_address,
    c.name as carrier_name
FROM transport_schedules ts
JOIN shipments s ON ts.shipment_id = s.shipment_id
LEFT JOIN carriers c ON s.carrier_id = c.carrier_id
WHERE ts.route_status = $1
ORDER BY ts.planned_departure;

-- name: UpdateShipmentTransportSchedule :one
UPDATE transport_schedules
SET 
    planned_route = $1,
    planned_departure = $2,
    planned_arrival = $3,
    estimated_duration = $4,
    route_status = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE schedule_id = $6
RETURNING *;

-- name: GetShipmentSchedulesByDateRange :many
SELECT 
    ts.*,
    s.origin_warehouse_id,
    s.destination_address,
    c.name as carrier_name
FROM transport_schedules ts
JOIN shipments s ON ts.shipment_id = s.shipment_id
LEFT JOIN carriers c ON s.carrier_id = c.carrier_id
WHERE ts.planned_departure BETWEEN $1 AND $2
ORDER BY ts.planned_departure; 