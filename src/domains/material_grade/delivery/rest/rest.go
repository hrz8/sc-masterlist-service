package rest

import (
	MaterialGradeError "github.com/hrz8/sc-masterlist-service/src/domains/material_grade/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/material_grade/usecase"
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
	payload := ctx.Payload.(*models.MaterialGradePayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialGradeError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create material grade",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialGradePayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialGradeError.GetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all material grade",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialGradePayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialGradeError.GetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get material grade",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialGradePayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialGradeError.DeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete material grade",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MaterialGradePayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MaterialGradeError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update material grade",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewMaterialGradeError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
