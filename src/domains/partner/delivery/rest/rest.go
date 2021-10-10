package rest

import (
	PartnerError "github.com/hrz8/sc-masterlist-service/src/domains/partner/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner/usecase"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

type (
	RestInterface interface {
		Create(c echo.Context) error
		GetAll(c echo.Context) error
		GetById(c echo.Context) error
		DeleteById(c echo.Context) error
		UpdateById(c echo.Context) error
		AddPartnerType(c echo.Context) error
		DeletePartnerType(c echo.Context) error
	}

	impl struct {
		usecase  usecase.UsecaseInterface
		errorLib RestErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create partner",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.GetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all partner",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.GetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get partner",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.DeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete partner",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update partner",
		nil,
	)
}

func (i *impl) AddPartnerType(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerAddPartnerTypePayload)
	result, err := i.usecase.AddPartnerType(ctx, &payload.ID, &payload.PartnerTypeID)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.AddPartnerType.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success add partner type into partner",
		nil,
	)
}

func (i *impl) DeletePartnerType(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerDeletePartnerTypePayload)
	result, err := i.usecase.DeletePartnerType(ctx, &payload.ID, &payload.PartnerTypeID)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.DeletePartnerType.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete partner type from partner",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewPartnerError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
