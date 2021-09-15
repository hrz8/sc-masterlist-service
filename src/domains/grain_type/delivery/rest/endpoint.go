package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddGrainTypeEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/grain-type", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.GrainTypePayloadCreate{}), false))
	e.GET("/api/v1/grain-type", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.GrainTypePayloadGetAll{}), true))
	e.GET("/api/v1/grain-type/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.GrainTypePayloadGet{}), false))
	e.DELETE("/api/v1/grain-type/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.GrainTypePayloadDeleteById{}), false))
	e.PUT("/api/v1/grain-type/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.GrainTypePayloadUpdateById{}), false))
}
