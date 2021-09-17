package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddMaterialGradeEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/material-grade", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialGradePayloadCreate{}), false))
	e.GET("/api/v1/material-grade", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialGradePayloadGetAll{}), true))
	e.GET("/api/v1/material-grade/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialGradePayloadGet{}), false))
	e.DELETE("/api/v1/material-grade/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialGradePayloadDeleteById{}), false))
	e.PUT("/api/v1/material-grade/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.MaterialGradePayloadUpdateById{}), false))
}
