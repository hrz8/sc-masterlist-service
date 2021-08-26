package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/labstack/echo/v4"
)

type (
	RestInterface interface {
		Create(c echo.Context) error
	}

	impl struct {
		usecase usecase.UsecaseInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	payload := new(models.ProcessCreatePayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err := c.Validate(payload); err != nil {
		return err
	}
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
