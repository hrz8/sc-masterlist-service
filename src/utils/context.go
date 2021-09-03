package utils

import (
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	CustomContext struct {
		echo.Context
		MysqlSess *gorm.DB
		Payload   interface{}
	}
)

func (c *CustomContext) SuccessResponse(data interface{}, message string, status uint16, meta interface{}) error {
	return c.JSON(int(status), &SuccessResponse{
		Data:    data,
		Message: message,
		Status:  int(status),
		Meta:    meta,
	})
}

func (c *CustomContext) ErrorResponse(data interface{}, message string, status uint16, errorCode string, meta interface{}) error {
	return c.JSON(int(status), &ErrorResponse{
		Data:      helpers.NilToEmptyMap(&data),
		Message:   message,
		Status:    int(status),
		ErrorCode: errorCode,
		Meta:      helpers.NilToEmptyMap(&meta),
	})
}
