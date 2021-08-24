package main

import (
	"fmt"
	"net/http"

	Config "github.com/hrz8/sc-masterlist-service/src/config"
	Container "github.com/hrz8/sc-masterlist-service/src/container"
	Database "github.com/hrz8/sc-masterlist-service/src/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	appContainer := Container.NewContainer()
	appConfig := appContainer.MustGet("shared.config").(*Config.AppConfig)
	mysql := appContainer.MustGet("shared.mysql").(Database.MysqlInterface)

	mysqlSess, _ := mysql.Connect()
	fmt.Println(mysqlSess)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE_PORT)))
}
