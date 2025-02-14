package entity

import "time"

type AlertRelatedEntityType string

const (
	AlertRelatedEntityTypeShipment AlertRelatedEntityType = "SHIPMENT"
	AlertRelatedEntityTypeCarrier  AlertRelatedEntityType = "CARRIER"
)

type AlertType string

const (
	AlertTypeDelay        AlertType = "DELAY"
	AlertTypeRouteChange  AlertType = "ROUTE_CHANGE"
	AlertTypeCarrierIssue AlertType = "CARRIER_ISSUE"
)

type AlertSeverity string

const (
	AlertSeverityLow    AlertSeverity = "LOW"
	AlertSeverityMedium AlertSeverity = "MEDIUM"
	AlertSeverityHigh   AlertSeverity = "HIGH"
)

type AlertStatus string

const (
	AlertStatusNew          AlertStatus = "NEW"
	AlertStatusAcknowledged AlertStatus = "ACKNOWLEDGED"
	AlertStatusResolved     AlertStatus = "RESOLVED"
)

type Alert struct {
	Id                int                    `json:"id"`
	RelatedEntityType AlertRelatedEntityType `json:"related_entity_type"`
	RelatedEntityId   int                    `json:"related_entity_id"`
	AlertType         AlertType              `json:"alert_type"`
	AlertSeverity     AlertSeverity          `json:"alert_severity"`
	Description       string                 `json:"description"`
	Status            AlertStatus            `json:"status"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}
