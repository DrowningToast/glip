-- name: CreateAlert :one
INSERT INTO alerts (
    related_entity_type,
    related_entity_id,
    alert_type,
    severity,
    description,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAlertById :one
SELECT * FROM alerts
WHERE alert_id = $1;

-- name: ListActiveAlerts :many
SELECT * FROM alerts
WHERE status != 'resolved'
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: ListAlertsByType :many
SELECT * FROM alerts
WHERE alert_type = $1 AND status != 'resolved'
ORDER BY created_at DESC;

-- name: ListAlertsBySeverity :many
SELECT * FROM alerts
WHERE severity = $1 AND status != 'resolved'
ORDER BY created_at DESC;

-- name: ListAlertsByEntity :many
SELECT * FROM alerts
WHERE related_entity_type = $1 
AND related_entity_id = $2 
AND status != 'resolved'
ORDER BY created_at DESC;

-- name: UpdateAlertStatus :one
UPDATE alerts
SET 
    status = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE alert_id = $2
RETURNING *;

-- name: GetActiveAlertCount :one
SELECT 
    COUNT(*) as total_alerts,
    COUNT(CASE WHEN severity = 'high' THEN 1 END) as high_severity,
    COUNT(CASE WHEN severity = 'medium' THEN 1 END) as medium_severity,
    COUNT(CASE WHEN severity = 'low' THEN 1 END) as low_severity
FROM alerts
WHERE status != 'resolved'; 