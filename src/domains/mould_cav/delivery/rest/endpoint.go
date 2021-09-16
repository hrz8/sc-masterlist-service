package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddMouldCavEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/mould-cav", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldCavPayloadCreate{}), false))
	e.GET("/api/v1/mould-cav", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldCavPayloadGetAll{}), true))
	e.GET("/api/v1/mould-cav/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldCavPayloadGet{}), false))
	e.DELETE("/api/v1/mould-cav/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldCavPayloadDeleteById{}), false))
	e.PUT("/api/v1/mould-cav/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldCavPayloadUpdateById{}), false))
}
