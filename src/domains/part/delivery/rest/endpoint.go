package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddPartEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/part", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.PartPayloadCreate{}), false))
}
