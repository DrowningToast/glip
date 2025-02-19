-- name: CreateShipmentAlert :one
INSERT INTO alerts (
    related_entity_type,
    related_entity_id,
    alert_type,
    severity,
    description,
    status
) VALUES (
    @related_entity_type, @related_entity_id,
    @alert_type, @severity, @description, @status
) RETURNING *;

-- name: GetShipmentAlertById :one
SELECT * FROM alerts
WHERE id = @id;

-- name: ListShipmentActiveAlerts :many
SELECT * FROM alerts
WHERE status != 'resolved'
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentAlertsByType :many
SELECT * FROM alerts
WHERE alert_type = @alert_type AND status != 'resolved'
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentAlertsBySeverity :many
SELECT * FROM alerts
WHERE severity = @severity AND status != 'resolved'
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentAlertsByEntityType :many
SELECT * FROM alerts
WHERE related_entity_type = @entity_type
AND status != 'resolved'
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: ListShipmentAlertsByEntityId :many
SELECT * FROM alerts
WHERE related_entity_type = @entity_type
AND related_entity_id = @entity_id
AND status != 'resolved'
ORDER BY created_at DESC
LIMIT sqlc.narg(return_limit) OFFSET sqlc.narg(return_offset);

-- name: UpdateShipmentAlertStatus :one
UPDATE alerts
SET 
    status = @status,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;
