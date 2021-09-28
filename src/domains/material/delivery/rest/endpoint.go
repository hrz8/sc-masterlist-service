package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddMaterialEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/material", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialPayloadCreate{}), false))
	e.GET("/api/v1/material", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialPayloadGetAll{}), true))
	e.GET("/api/v1/material/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialPayloadGet{}), false))
	e.DELETE("/api/v1/material/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialPayloadDeleteById{}), false))
	e.PUT("/api/v1/material/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialPayloadUpdateById{}), false))
}
