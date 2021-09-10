package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
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
		errorLib ProcessErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, ProcessErrorCreate.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create process",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, ProcessErrorGetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all process",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, ProcessErrorGetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get process",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, ProcessErrorDeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete process",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, ProcessErrorUpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update process",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewProcessError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
