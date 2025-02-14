-- name: CreateInventoryAlert :one
INSERT INTO alerts (
    related_entity_type,
    related_entity_id,
    alert_type,
    severity,
    description,
    status
) VALUES (
    @relatedEntityType, @relatedEntityId,
    @alertType, @severity, @description, @status
) RETURNING *;

-- name: GetInventoryAlertById :one
SELECT * FROM alerts
WHERE id = @id;

-- name: ListInventoryActiveAlerts :many
SELECT * FROM alerts
WHERE status != 'resolved'
ORDER BY created_at DESC
LIMIT sqlc.narg(returnLimit) OFFSET sqlc.narg(returnOffset);

-- name: ListInventoryAlertsByType :many
SELECT * FROM alerts
WHERE alert_type = @alertType AND status != 'resolved'
ORDER BY created_at DESC;

-- name: ListInventoryAlertsBySeverity :many
SELECT * FROM alerts
WHERE severity = @severity AND status != 'resolved'
ORDER BY created_at DESC;

-- name: ListInventoryAlertsByEntity :many
SELECT * FROM alerts
WHERE related_entity_type = @entityType 
AND related_entity_id = @entityId 
AND status != 'resolved'
ORDER BY created_at DESC;

-- name: UpdateInventoryAlertStatus :one
UPDATE alerts
SET 
    status = @status,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: GetInventoryActiveAlertCount :one
SELECT 
    COUNT(*) as total_alerts,
    COUNT(CASE WHEN severity = 'high' THEN 1 END) as high_severity,
    COUNT(CASE WHEN severity = 'medium' THEN 1 END) as medium_severity,
    COUNT(CASE WHEN severity = 'low' THEN 1 END) as low_severity
FROM alerts
WHERE status != 'resolved'; 