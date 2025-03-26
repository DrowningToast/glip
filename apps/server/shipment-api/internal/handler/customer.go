package handler

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/drowningtoast/glip/apps/server/internal/common"
	"github.com/drowningtoast/glip/apps/server/internal/errs"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/entity"
	"github.com/drowningtoast/glip/apps/server/shipment-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateCustomer(ctx *fiber.Ctx) error {
	var body struct {
		Customer usecase.CreateCustomerParams `json:"customer" validate:"required"`
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return errors.Wrap(errs.ErrInvalidBody, err.Error())
	}

	account, customer, err := h.Uc.CreateCustomer(ctx.Context(), body.Customer)
	if err != nil {
		return err
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Account  entity.Account  `json:"account"`
			Customer entity.Customer `json:"customer"`
		}{
			Account:  *account,
			Customer: *customer,
		},
	})
}

// query: limit, offset
func (h *Handler) ListCustomers(ctx *fiber.Ctx) error {
	var queries struct {
		Limit  *int `query:"limit,omitempty"`
		Offset *int `query:"offset,omitempty"`
	}

	if err := ctx.QueryParser(&queries); err != nil {
		return errors.Wrap(errs.ErrInvalidQueryString, err.Error())
	}

	var limit, offset int = 100, 0
	if queries.Limit != nil {
		limit = *queries.Limit
	}
	if queries.Offset != nil {
		offset = *queries.Offset
	}

	customerPtrs, err := h.Uc.ListCustomers(ctx.Context(), &usecase.ListCustomersQuery{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get customers")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: common.PaginatedResult[*entity.Customer]{
			Count:  len(customerPtrs),
			Items:  customerPtrs,
			Offset: offset,
			Limit:  limit,
		},
	})
}

// query: id, email
func (h *Handler) GetCustomer(ctx *fiber.Ctx) error {
	var queries struct {
		Id    *int    `query:"id,omitempty"`
		Email *string `query:"email,omitempty"`
	}

	if err := ctx.QueryParser(&queries); err != nil {
		return errors.Wrap(errs.ErrInvalidQueryString, err.Error())
	}

	if queries.Id == nil && queries.Email == nil {
		return errors.Wrap(errs.ErrInvalidQueryString, "id or email is required")
	}

	customer, err := h.Uc.GetCustomer(ctx.Context(), usecase.GetCustomerQuery{
		Id:    queries.Id,
		Email: queries.Email,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get customer")
	}
	if customer == nil {
		return errors.Wrap(errs.ErrNotFound, "customer not found")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Customer *entity.Customer `json:"customer"`
		}{
			Customer: customer,
		},
	})
}

func (h *Handler) UpdateCustomer(ctx *fiber.Ctx) error {
	var body struct {
		Customer struct {
			Id      int     `json:"id" validate:"required"`
			Name    string  `json:"name" validate:"required"`
			Email   string  `json:"email,omitempty"`
			Phone   *string `json:"phone,omitempty"`
			Address *string `json:"address,omitempty"`
		}
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return errors.Wrap(errs.ErrInvalidBody, err.Error())
	}

	customer, err := h.Uc.UpdateCustomer(ctx.Context(), &entity.Customer{
		Id:      body.Customer.Id,
		Name:    body.Customer.Name,
		Email:   body.Customer.Email,
		Phone:   body.Customer.Phone,
		Address: body.Customer.Address,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update customer")
	}

	return ctx.JSON(common.HTTPResponse{
		Result: struct {
			Customer *entity.Customer `json:"customer"`
		}{
			Customer: customer,
		},
	})
}

func (h *Handler) DeleteCustomer(ctx *fiber.Ctx) error {
	customerId := ctx.Params("id")

	customerIdInt, err := strconv.ParseInt(customerId, 10, 64)
	if err != nil {
		return errors.Wrap(errs.ErrInvalidQueryString, err.Error())
	}

	err = h.Uc.DeleteCustomer(ctx.Context(), int(customerIdInt))
	if err != nil {
		return errors.Wrap(err, "failed to delete customer")
	}

	return ctx.JSON(common.EmptyHTTPResponse)
}
