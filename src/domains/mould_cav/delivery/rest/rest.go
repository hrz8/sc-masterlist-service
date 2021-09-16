package rest

import (
	MouldCavError "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/usecase"
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
	payload := ctx.Payload.(*models.MouldCavPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MouldCavError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create mould cav",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldCavPayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MouldCavError.GetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all mould cav",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldCavPayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, MouldCavError.GetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get mould cav",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldCavPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, MouldCavError.DeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete mould cav",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldCavPayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MouldCavError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update mould cav",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewMouldCavError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
