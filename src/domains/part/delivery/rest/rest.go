package rest

import (
	PartError "github.com/hrz8/sc-masterlist-service/src/domains/part/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/part/usecase"
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
	payload := ctx.Payload.(*models.PartPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, PartError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create part",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	errLib := NewPartError()
	return &impl{
		usecase:  u,
		errorLib: errLib,
	}
}
