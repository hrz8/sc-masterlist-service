package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddPartnerTypeEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/partner-type", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerTypePayloadCreate{}), false))
	e.GET("/api/v1/partner-type", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerTypePayloadGetAll{}), true))
	e.GET("/api/v1/partner-type/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerTypePayloadGet{}), false))
	e.DELETE("/api/v1/partner-type/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerTypePayloadDeleteById{}), false))
	e.PUT("/api/v1/partner-type/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerTypePayloadUpdateById{}), false))
}
