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
	payload := &models.ProcessCreatePayload{
		Name:        "ASSY",
		Description: "assembly process",
	}
	result, _ := i.usecase.Create(payload)
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   result,
	})
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	handler := &impl{
		usecase: u,
	}
	return handler
}
