package utils

import (
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

func (c *CustomContext) SuccessResponse(data interface{}, message string, status int, meta interface{}) error {
	return c.JSON(status, &SuccessResponse{
		Data:    data,
		Message: message,
		Status:  status,
		Meta:    meta,
	})
}

func (c *CustomContext) ErrorResponse(data interface{}, message string, status int, errorCode string, meta interface{}) error {
	return c.JSON(status, &ErrorResponse{
		Data:      data,
		Message:   message,
		Status:    status,
		ErrorCode: errorCode,
		Meta:      meta,
	})
}
