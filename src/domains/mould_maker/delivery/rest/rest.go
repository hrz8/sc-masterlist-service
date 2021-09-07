package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/domains/mould_maker/usecase"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

type (
	RestInterface interface {
		Create(echo.Context) error
		GetAll(echo.Context) error
		GetById(echo.Context) error
		DeleteById(echo.Context) error
		UpdateById(echo.Context) error
	}

	impl struct {
		usecase  usecase.UsecaseInterface
		errorLib MouldMakerErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldMakerPayloadCreate)
	result, err := i.usecase.Create(payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := uint16(http.StatusBadRequest)
		return i.errorLib.MouldMakerErrorCreate(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success create mould maker",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldMakerPayloadGetAll)
	result, total, err := i.usecase.GetAll(payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := uint16(http.StatusBadRequest)
		return i.errorLib.MouldMakerErrorGetAll(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all mould maker",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldMakerPayloadGet)
	result, err := i.usecase.GetById(&payload.ID)
	if err != nil {
		errMessage := err.Error()
		errStatus := helpers.ParseStatusResponse(err, uint16(http.StatusBadRequest))
		return i.errorLib.MouldMakerErrorGet(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success get mould maker",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldMakerPayloadDeleteById)
	result, err := i.usecase.DeleteById(&payload.ID)
	if err != nil {
		errMessage := err.Error()
		errStatus := helpers.ParseStatusResponse(err, uint16(http.StatusBadRequest))
		return i.errorLib.MouldMakerErrorGet(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success delete mould maker",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.MouldMakerPayloadUpdateById)
	result, err := i.usecase.UpdateById(&payload.ID, payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := helpers.ParseStatusResponse(err, uint16(http.StatusBadRequest))
		return i.errorLib.MouldMakerErrorGet(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success update mould maker",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewMouldMakerError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
