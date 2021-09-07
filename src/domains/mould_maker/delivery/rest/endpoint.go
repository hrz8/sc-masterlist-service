package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddMouldMakerEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/mould-maker", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldMakerPayloadCreate{}), false))
	e.GET("/api/v1/mould-maker", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldMakerPayloadGetAll{}), true))
	e.GET("/api/v1/mould-maker/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldMakerPayloadGet{}), false))
	e.DELETE("/api/v1/mould-maker/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldMakerPayloadDeleteById{}), false))
	e.PUT("/api/v1/mould-maker/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MouldMakerPayloadUpdateById{}), false))
}
