package rest

import (
	PartnerError "github.com/hrz8/sc-masterlist-service/src/domains/partner/error"
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
		errorLib RestErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.PartnerPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartnerError.Create.Err, err)
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
