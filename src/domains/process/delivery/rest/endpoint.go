package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddProcessEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/process", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProcessPayloadCreate{}), false))
	e.GET("/api/v1/process", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProcessPayloadGetAll{}), true))
	e.GET("/api/v1/process/:id", rest.Get, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProcessPayloadGet{}), false))
	e.DELETE("/api/v1/process/:id", rest.Delete, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProcessPayloadDelete{}), false))
}
