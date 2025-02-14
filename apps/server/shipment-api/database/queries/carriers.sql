-- name: CreateCarrier :one
INSERT INTO carriers (
    name,
    carrier_type,
    contact_person,
    contact_phone,
    email,
    description,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetCarrierById :one
SELECT * FROM carriers
WHERE carrier_id = $1;

-- name: ListCarriers :many
SELECT * FROM carriers
ORDER BY created_at DESC;

-- name: ListActiveCarriers :many
SELECT * FROM carriers
WHERE status = 'active'
ORDER BY name;

-- name: ListCarriersByType :many
SELECT * FROM carriers
WHERE carrier_type = $1
ORDER BY name;

-- name: UpdateCarrier :one
UPDATE carriers
SET 
    name = $1,
    carrier_type = $2,
    contact_person = $3,
    contact_phone = $4,
    email = $5,
    description = $6,
    status = $7
WHERE carrier_id = $8
RETURNING *;

-- name: GetCarrierShipmentStats :one
SELECT 
    c.*,
    COUNT(s.shipment_id) as total_shipments,
    COUNT(CASE WHEN s.status = 'DELIVERED' THEN 1 END) as delivered_shipments,
    COUNT(CASE WHEN s.status = 'CANCELLED' THEN 1 END) as cancelled_shipments,
    AVG(CASE 
        WHEN s.actual_arrival IS NOT NULL AND s.scheduled_arrival IS NOT NULL 
        THEN EXTRACT(EPOCH FROM (s.actual_arrival - s.scheduled_arrival))/3600 
    END) as avg_delay_hours
FROM carriers c
LEFT JOIN shipments s ON c.carrier_id = s.carrier_id
WHERE c.carrier_id = $1
GROUP BY c.carrier_id; 