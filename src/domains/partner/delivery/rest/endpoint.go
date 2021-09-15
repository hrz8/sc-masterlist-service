package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddPartnerEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/partner", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerPayloadCreate{}), false))
	e.GET("/api/v1/partner", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerPayloadGetAll{}), true))
	e.GET("/api/v1/partner/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerPayloadGet{}), false))
	e.DELETE("/api/v1/partner/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerPayloadDeleteById{}), false))
	e.PUT("/api/v1/partner/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerPayloadUpdateById{}), false))
	e.PUT("/api/v1/partner/:id/partner-type/:partnerTypeId", rest.AddPartnerType, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerAddPartnerTypePayload{}), false))
	e.DELETE("/api/v1/partner/:id/partner-type/:partnerTypeId", rest.DeletePartnerType, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartnerDeletePartnerTypePayload{}), false))
}
