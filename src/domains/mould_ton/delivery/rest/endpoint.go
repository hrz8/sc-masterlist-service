package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddMouldTonEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/mould-ton", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldTonPayloadCreate{}), false))
	e.GET("/api/v1/mould-ton", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldTonPayloadGetAll{}), true))
	e.GET("/api/v1/mould-ton/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldTonPayloadGet{}), false))
	e.DELETE("/api/v1/mould-ton/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldTonPayloadDeleteById{}), false))
	e.PUT("/api/v1/mould-ton/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldTonPayloadUpdateById{}), false))
}
