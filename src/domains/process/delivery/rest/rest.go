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
	cc := c.(*utils.CustomContext)
	payload := cc.Payload.(*models.ProcessCreatePayload)
	result, err := i.usecase.Create(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"data":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   result,
	})
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	return &impl{
		usecase: u,
	}
}
