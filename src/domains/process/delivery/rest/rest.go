package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
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
		errorLib ProcessErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadCreate)
	result, err := i.usecase.Create(payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := uint16(http.StatusBadRequest)
		return i.errorLib.ProcessErrorCreate(ctx, &errMessage, &errStatus)
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
	result, err := i.usecase.GetAll(payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := uint16(http.StatusBadRequest)
		return i.errorLib.ProcessErrorGetAll(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all process",
		utils.ListMetaResponse{
			Total: len(*result),
		},
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadGet)
	result, err := i.usecase.GetById(&payload.ID)
	if err != nil {
		errMessage := err.Error()
		errStatus := helpers.ParseStatusResponse(err, uint16(http.StatusBadRequest))
		return i.errorLib.ProcessErrorGet(ctx, &errMessage, &errStatus)
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
	result, err := i.usecase.DeleteById(&payload.ID)
	if err != nil {
		errMessage := err.Error()
		errStatus := helpers.ParseStatusResponse(err, uint16(http.StatusBadRequest))
		return i.errorLib.ProcessErrorGet(ctx, &errMessage, &errStatus)
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
	result, err := i.usecase.UpdateById(&payload.ID, payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := helpers.ParseStatusResponse(err, uint16(http.StatusBadRequest))
		return i.errorLib.ProcessErrorGet(ctx, &errMessage, &errStatus)
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
