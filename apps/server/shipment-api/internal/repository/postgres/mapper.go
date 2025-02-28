package postgres

import (
	"math/big"
	"time"

	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func mapShipmentModelToEntity(shipment *shipment_database.Shipment) *entity.Shipment {
	var lastWarehouseId *int
	if shipment.LastWarehouseID.Valid {
		i := int(shipment.LastWarehouseID.Int32)
		lastWarehouseId = &i
	}

	var carrierId *int
	if shipment.CarrierID.Valid {
		i := int(shipment.CarrierID.Int32)
		carrierId = &i
	}

	var actualDeparture *time.Time
	if shipment.ActualDeparture.Valid {
		actualDeparture = &shipment.ActualDeparture.Time
	}

	var actualArrival *time.Time
	if shipment.ActualArrival.Valid {
		actualArrival = &shipment.ActualArrival.Time
	}

	var specialInstructions *string
	if shipment.SpecialInstructions.Valid {
		specialInstructions = &shipment.SpecialInstructions.String
	}

	return &entity.Shipment{
		Id:                  int(shipment.ID),
		Route:               lo.Map(shipment.Route, func(id int32, _ int) int { return int(id) }),
		LastWarehouseId:     lastWarehouseId,
		DestinationAddress:  shipment.DestinationAddress,
		CarrierId:           carrierId,
		ScheduledDeparture:  shipment.ScheduledDeparture.Time,
		ScheduledArrival:    shipment.ScheduledArrival.Time,
		ActualDeparture:     actualDeparture,
		ActualArrival:       actualArrival,
		Status:              entity.ShipmentStatus(shipment.Status),
		TotalWeight:         decimal.New(shipment.TotalWeight.Int.Int64(), shipment.TotalWeight.Exp),
		TotalVolume:         decimal.New(shipment.TotalVolume.Int.Int64(), shipment.TotalVolume.Exp),
		SpecialInstructions: specialInstructions,
	}
}

func mapAccountModelToEntity(account *shipment_database.Account) *entity.Account {
	return &entity.Account{
		Id:       int(account.ID),
		Username: account.Username,
		Password: account.Password,
		Role:     entity.AccountRole(account.Role),
	}
}

func mapShipmentOwnerModelToEntity(owner *shipment_database.Owner) *entity.Customer {
	return &entity.Customer{
		Id:      int(owner.ID),
		Name:    owner.Name,
		Email:   owner.Email,
		Phone:   mapPgTextToStringPtr(owner.Phone),
		Address: mapPgTextToStringPtr(owner.Address),
	}
}

func mapCarrierModelToEntity(carrier *shipment_database.Carrier) *entity.Carrier {
	return &entity.Carrier{
		Id:            int(carrier.ID),
		Name:          carrier.Name,
		Status:        entity.CarrierStatus(carrier.Status),
		ContactPerson: mapPgTextToStringPtr(carrier.ContactPerson),
		ContactPhone:  mapPgTextToStringPtr(carrier.ContactPhone),
		Email:         mapPgTextToStringPtr(carrier.Email),
		Description:   mapPgTextToStringPtr(carrier.Description),
	}
}

func mapWarehouseConnectionModelToEntity(connection *shipment_database.WarehouseConnection) *entity.WarehouseConnection {
	return &entity.WarehouseConnection{
		Id:          int(connection.ID),
		WarehouseId: int(connection.WarehouseID),
		ApiKey:      connection.ApiKey,
		Name:        connection.Name,
		Status:      entity.WarehouseConnectionStatus(connection.Status),
		CreatedAt:   &connection.CreatedAt.Time,
		UpdatedAt:   &connection.UpdatedAt.Time,
		LastUsedAt:  &connection.LastUsedAt.Time,
	}
}

func mapAlertModelToEntity(alert *shipment_database.Alert) *entity.Alert {
	return &entity.Alert{
		Id:                int(alert.ID),
		RelatedEntityType: entity.AlertRelatedEntityType(alert.RelatedEntityType),
		RelatedEntityId:   int(alert.RelatedEntityID),
		Type:              entity.AlertType(alert.AlertType),
		Severity:          entity.AlertSeverity(alert.Severity),
		Description:       mapPgTextToStringPtr(alert.Description),
		Status:            entity.AlertStatus(alert.Status),
		CreatedAt:         alert.CreatedAt.Time,
		UpdatedAt:         alert.UpdatedAt.Time,
	}
}

func mapIntPtrToPgInt4(i *int) pgtype.Int4 {
	if i == nil {
		return pgtype.Int4{Valid: false}
	}
	return pgtype.Int4{Int32: int32(*i), Valid: true}
}

func mapTimeToTimestamp(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: t, Valid: true}
}

func mapTimeToTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: true}
}

func mapTimePtrToTimestamp(t *time.Time) pgtype.Timestamp {
	if t == nil {
		return pgtype.Timestamp{Valid: false}
	}
	return pgtype.Timestamp{Time: *t, Valid: true}
}

func mapTimePtrToTimestamptz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{Time: *t, Valid: true}
}

func mapDecimalToPgNumeric(d decimal.Decimal) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   big.NewInt(d.IntPart()),
		Exp:   d.Exponent(),
		Valid: true,
	}
}

func mapDecimalPtrToPgNumeric(d *decimal.Decimal) pgtype.Numeric {
	if d == nil {
		return pgtype.Numeric{Valid: false}
	}
	return mapDecimalToPgNumeric(*d)
}

func mapStringPtrToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func mapPgTextToStringPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}
