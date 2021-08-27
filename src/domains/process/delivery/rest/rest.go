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
	}

	impl struct {
		usecase usecase.UsecaseInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProcessCreatePayload)
	result, err := i.usecase.Create(payload)
	if err != nil {
		return ctx.ErrorResponse(nil, err.Error(), http.StatusBadRequest, "SCM-PROCESS-001", nil)
	}
	return ctx.SuccessResponse(result, "success create process", http.StatusOK, nil)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	return &impl{
		usecase: u,
	}
}
