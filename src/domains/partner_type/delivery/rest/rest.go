package rest

import (
	PartnerTypeError "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner_type/usecase"
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
	}

	impl struct {
		usecase  usecase.UsecaseInterface
		errorLib RestErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerTypePayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerTypeError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create partner type",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerTypePayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerTypeError.GetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all partner type",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerTypePayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerTypeError.GetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get partner type",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerTypePayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerTypeError.DeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete partner type",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerTypePayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerTypeError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update partner type",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewRestError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
