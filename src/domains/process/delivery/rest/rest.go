package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

type (
	RestInterface interface {
		Create(echo.Context) error
		GetAll(echo.Context) error
		Get(echo.Context) error
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
		return i.errorLib.ProcessErrorCreate(ctx, err.Error())
	}
	return ctx.SuccessResponse(result, "success create process", http.StatusOK, nil)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadGetAll)
	result, err := i.usecase.GetAll(payload)
	if err != nil {
		return i.errorLib.ProcessErrorGetAll(ctx, err.Error())
	}
	return ctx.SuccessResponse(result, "success fetch all process", http.StatusOK, nil)
}

func (i *impl) Get(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessPayloadGet)
	result, err := i.usecase.Get(&payload.ID)
	if err != nil {
		return i.errorLib.ProcessErrorGet(ctx, err.Error())
	}
	return ctx.SuccessResponse(result, "success get process", http.StatusOK, nil)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewProcessError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
