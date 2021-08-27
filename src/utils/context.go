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
