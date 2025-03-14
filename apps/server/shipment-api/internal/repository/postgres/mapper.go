package postgres

import (
	"github.com/drowningtoast/glip/apps/server/internal/utils/pgmapper"
	database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	shipment_database "github.com/drowningtoast/glip/apps/server/shipment-api/database/gen"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func mapShipmentModelToEntity(shipment *shipment_database.Shipment) *entity.Shipment {
	var lastWarehouseId *string
	if shipment.LastWarehouseID.Valid {
		lastWarehouseId = &shipment.LastWarehouseID.String
	}

	var specialInstructions *string
	if shipment.SpecialInstructions.Valid {
		specialInstructions = &shipment.SpecialInstructions.String
	}

	return &entity.Shipment{
		Id:                     int(shipment.ID),
		Route:                  lo.Map(shipment.Route, func(id string, _ int) string { return id }),
		LastWarehouseId:        lastWarehouseId,
		DepartureWarehouseId:   shipment.DepartureWarehouseID,
		DepartureAddress:       pgmapper.MapPgTextToStringPtr(shipment.DepartureAddress),
		DestinationWarehouseId: shipment.DestinationWarehouseID,
		DestinationAddress:     shipment.DestinationAddress,
		// CarrierId:              carrierId,
		Status:              entity.ShipmentStatus(shipment.Status),
		TotalWeight:         decimal.New(shipment.TotalWeight.Int.Int64(), shipment.TotalWeight.Exp),
		TotalVolume:         decimal.New(shipment.TotalVolume.Int.Int64(), shipment.TotalVolume.Exp),
		SpecialInstructions: specialInstructions,
	}
}

func mapShipmentJoinedAccountModelToEntity(s database.ListShipmentsByAccountUsernameRow) *entity.Shipment {
	var lastWarehouseId *string
	if s.LastWarehouseID.Valid {
		lastWarehouseId = &s.LastWarehouseID.String
	}

	var specialInstructions *string
	if s.SpecialInstructions.Valid {
		specialInstructions = &s.SpecialInstructions.String
	}

	return &entity.Shipment{
		Id:                     int(s.ID),
		Route:                  lo.Map(s.Route, func(id string, _ int) string { return id }),
		LastWarehouseId:        lastWarehouseId,
		DepartureWarehouseId:   s.DepartureWarehouseID,
		DepartureAddress:       pgmapper.MapPgTextToStringPtr(s.DepartureAddress),
		DestinationWarehouseId: s.DestinationWarehouseID,
		DestinationAddress:     s.DestinationAddress,
		Status:                 entity.ShipmentStatus(s.Status),
		TotalWeight:            decimal.New(s.TotalWeight.Int.Int64(), s.TotalWeight.Exp),
		TotalVolume:            decimal.New(s.TotalVolume.Int.Int64(), s.TotalVolume.Exp),
		SpecialInstructions:    specialInstructions,
	}
}

func mapAccountModelToEntity(account *shipment_database.Account) *entity.Account {
	return &entity.Account{
		Id:       int(account.ID),
		Username: account.Username,
		Password: account.Password,
		Email:    account.Email,
	}
}

func mapShipmentOwnerModelToEntity(owner *shipment_database.Owner) *entity.Customer {
	var accountId *int
	if owner.AccountID.Valid {
		i := int(owner.AccountID.Int32)
		accountId = &i
	}

	return &entity.Customer{
		Id:        int(owner.ID),
		Name:      owner.Name,
		Email:     owner.Email,
		Phone:     pgmapper.MapPgTextToStringPtr(owner.Phone),
		Address:   pgmapper.MapPgTextToStringPtr(owner.Address),
		AccountId: accountId,
	}
}

func mapAlertModelToEntity(alert *shipment_database.Alert) *entity.Alert {
	return &entity.Alert{
		Id:                int(alert.ID),
		RelatedEntityType: entity.AlertRelatedEntityType(alert.RelatedEntityType),
		RelatedEntityId:   int(alert.RelatedEntityID),
		Type:              entity.AlertType(alert.AlertType),
		Severity:          entity.AlertSeverity(alert.Severity),
		Description:       pgmapper.MapPgTextToStringPtr(alert.Description),
		Status:            entity.AlertStatus(alert.Status),
		CreatedAt:         alert.CreatedAt.Time,
		UpdatedAt:         alert.UpdatedAt.Time,
	}
}

func mapCarrierModelToEntity(carrier *shipment_database.Carrier) *entity.Carrier {
	return &entity.Carrier{
		Id:            int(carrier.ID),
		Name:          carrier.Name,
		Status:        entity.CarrierStatus(carrier.Status),
		ContactPerson: pgmapper.MapPgTextToStringPtr(carrier.ContactPerson),
		ContactPhone:  pgmapper.MapPgTextToStringPtr(carrier.ContactPhone),
		Email:         pgmapper.MapPgTextToStringPtr(carrier.Email),
		Description:   pgmapper.MapPgTextToStringPtr(carrier.Description),
	}
}
