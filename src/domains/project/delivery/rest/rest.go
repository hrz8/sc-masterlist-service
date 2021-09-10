package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/domains/project/usecase"
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
		errorLib ProjectErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProjectPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, ProjectErrorCreate.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create project",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProjectPayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, ProjectErrorGetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all project",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProjectPayloadGet)
	result, err := i.usecase.GetById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, ProjectErrorGetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get project",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProjectPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, &payload.ID)
	if err != nil {
		return i.errorLib.Throw(ctx, ProjectErrorDeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete project",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProjectPayloadUpdateById)
	result, err := i.usecase.UpdateById(ctx, &payload.ID, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, ProjectErrorUpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update project",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewProjectError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
