package main

import (
	"fmt"

	ProcessRest "github.com/hrz8/sc-masterlist-service/src/domains/process/delivery/rest"
	ProcessRepository "github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	ProcessUsecase "github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
	Config "github.com/hrz8/sc-masterlist-service/src/shared/config"
	Container "github.com/hrz8/sc-masterlist-service/src/shared/container"
	Database "github.com/hrz8/sc-masterlist-service/src/shared/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	appContainer := Container.NewAppContainer()
	appConfig := appContainer.MustGet("shared.config").(*Config.AppConfig)
	mysql := appContainer.MustGet("shared.mysql").(Database.MysqlInterface)

	mysqlSess := mysql.Connect()

	// services loader
	processRepo := ProcessRepository.NewRepository(mysqlSess)
	processUsecase := ProcessUsecase.NewUsecase(processRepo)

	// rest loader
	processRest := ProcessRest.NewRest(processUsecase)

	// endpoints
	e.POST("/api/v1/process", processRest.Create)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE_PORT)))
}
