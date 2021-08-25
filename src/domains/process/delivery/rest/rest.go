package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/labstack/echo/v4"
)

func NewService(e *echo.Echo, usecase usecase.UsecaseInterface) {
	payload := &models.ProcessCreatePayload{
		Name:        "ASSY",
		Description: "",
	}
	usecase.Create(payload)
	e.POST("/api/v1/process", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
