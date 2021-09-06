package rest

import (
	"reflect"

	"github.com/hrz8/sc-masterlist-service/src/models"
	Utils "github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func AddProjectEndpoints(e *echo.Echo, rest RestInterface) {
	e.POST("/api/v1/project", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProjectPayloadCreate{}), false))
	e.GET("/api/v1/project", rest.GetAll, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProjectPayloadGetAll{}), true))
	e.GET("/api/v1/project/:id", rest.GetById, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProjectPayloadGet{}), false))
	e.DELETE("/api/v1/project/:id", rest.DeleteById, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProjectPayloadDeleteById{}), false))
	e.PUT("/api/v1/project/:id", rest.UpdateById, Utils.ValidatorMiddleware(reflect.TypeOf(models.ProjectPayloadUpdateById{}), false))
}
