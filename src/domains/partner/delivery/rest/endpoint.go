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
}
