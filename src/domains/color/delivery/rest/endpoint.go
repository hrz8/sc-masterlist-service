package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddColorEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/color", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.ColorPayloadCreate{}), false))
	e.GET("/api/v1/color", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.ColorPayloadGetAll{}), true))
	e.GET("/api/v1/color/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.ColorPayloadGet{}), false))
	e.DELETE("/api/v1/color/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.ColorPayloadDeleteById{}), false))
	e.PUT("/api/v1/color/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.ColorPayloadUpdateById{}), false))
}
