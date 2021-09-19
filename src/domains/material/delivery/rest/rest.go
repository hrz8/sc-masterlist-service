package rest

import (
	MaterialError "github.com/hrz8/sc-masterlist-service/src/domains/material/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/material/usecase"
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
	payload := ctx.Payload.(*models.MaterialPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create material",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialPayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialError.GetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all material",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialPayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialError.GetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get material",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialError.DeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete material",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialPayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update material",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewMaterialError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
