package main

import (
	"fmt"
	"net/http"

	Config "github.com/hrz8/sc-masterlist-service/src/config"
	Container "github.com/hrz8/sc-masterlist-service/src/container"
	Database "github.com/hrz8/sc-masterlist-service/src/database"
	ProcessRest "github.com/hrz8/sc-masterlist-service/src/domains/process/delivery/rest"
	ProcessRepository "github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	ProcessUsecase "github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	appContainer := Container.NewAppContainer()
	appConfig := appContainer.MustGet("shared.config").(*Config.AppConfig)
	mysql := appContainer.MustGet("shared.mysql").(Database.MysqlInterface)

	mysqlSess := mysql.Connect()

	processRepo := ProcessRepository.NewRepository(mysqlSess)
	processUsecase := ProcessUsecase.NewUsecase(processRepo)

	ProcessRest.NewService(e, processUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE_PORT)))
}
