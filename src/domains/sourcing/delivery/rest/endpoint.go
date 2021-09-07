package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddSourcingEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/sourcing", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.SourcingPayloadCreate{}), false))
	e.GET("/api/v1/sourcing", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.SourcingPayloadGetAll{}), true))
	e.GET("/api/v1/sourcing/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.SourcingPayloadGet{}), false))
	e.DELETE("/api/v1/sourcing/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.SourcingPayloadDeleteById{}), false))
	e.PUT("/api/v1/sourcing/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.SourcingPayloadUpdateById{}), false))
}
