package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/domains/partner/usecase"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

type (
	RestInterface interface {
		Create(c echo.Context) error
	}

	impl struct {
		usecase  usecase.UsecaseInterface
		errorLib PartnerErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		errMessage := err.Error()
		errStatus := uint16(http.StatusBadRequest)
		return i.errorLib.PartnerErrorCreate(ctx, &errMessage, &errStatus)
	}
	return ctx.SuccessResponse(
		result,
		"success create partner",
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
