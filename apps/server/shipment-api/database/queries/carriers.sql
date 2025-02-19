-- name: CreateCarrier :one
INSERT INTO carriers (
    name,
    
    contact_person,
    contact_phone,
    email,
    description,
    status
) VALUES (
    @name, @contact_person, @contact_phone, @email, @description, @status
) RETURNING *;

-- name: GetCarrierById :one
SELECT * FROM carriers
WHERE id = @id;

-- name: ListCarriers :many
SELECT * FROM carriers
ORDER BY created_at DESC;

-- name: ListActiveCarriers :many
SELECT * FROM carriers
WHERE status = 'active'
ORDER BY name;

-- name: UpdateCarrier :one
UPDATE carriers
SET 
    name = @name,
    contact_person = @contact_person,
    contact_phone = @contact_phone,
    email = @email,
    description = @description,
    status = @status
WHERE id = @id
RETURNING *;

-- name: GetCarrierShipmentStats :one
SELECT 
    c.*,
    COUNT(s.id) as total_shipments,
    COUNT(CASE WHEN s.status = 'DELIVERED' THEN 1 END) as delivered_shipments,
    COUNT(CASE WHEN s.status = 'CANCELLED' THEN 1 END) as cancelled_shipments,
    AVG(CASE 
        WHEN s.actual_arrival IS NOT NULL AND s.scheduled_arrival IS NOT NULL 
        THEN EXTRACT(EPOCH FROM (s.actual_arrival - s.scheduled_arrival))/3600 
    END) as avg_delay_hours
FROM carriers c
LEFT JOIN shipments s ON c.id = s.carrier_id
WHERE c.id = @id
GROUP BY c.id; 