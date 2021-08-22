package main

import (
	"fmt"
	"net/http"

	Config "github.com/hrz8/sc-masterlist-service/src/config"
	Database "github.com/hrz8/sc-masterlist-service/src/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	appConfig, err := Config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	mysqlSess := Database.NewMysql(appConfig).Connect()
	fmt.Println(mysqlSess)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE_PORT)))
}
