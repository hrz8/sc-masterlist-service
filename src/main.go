package main

import (
	"fmt"
	"reflect"

	ProcessRest "github.com/hrz8/sc-masterlist-service/src/domains/process/delivery/rest"
	ProcessRepository "github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	ProcessUsecase "github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"
	"github.com/hrz8/sc-masterlist-service/src/models"
	Config "github.com/hrz8/sc-masterlist-service/src/shared/config"
	Container "github.com/hrz8/sc-masterlist-service/src/shared/container"
	Database "github.com/hrz8/sc-masterlist-service/src/shared/database"
	"github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	appContainer := Container.NewAppContainer()
	appConfig := appContainer.MustGet("shared.config").(*Config.AppConfig)
	mysql := appContainer.MustGet("shared.mysql").(Database.MysqlInterface)

	mysqlSess := mysql.Connect()

	// services loader
	processRepo := ProcessRepository.NewRepository(mysqlSess)
	processUsecase := ProcessUsecase.NewUsecase(processRepo)

	// rest loader
	processRest := ProcessRest.NewRest(processUsecase)

	// rest server
	e := echo.New()
	e.Validator = utils.NewValidator()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &utils.CustomContext{
				Context:   ctx,
				MysqlSess: mysqlSess,
			}
			return next(cc)
		}
	})

	// endpoints
	e.POST("/api/v1/process", processRest.Create, utils.ValidatorMiddleware(reflect.TypeOf(models.ProcessCreatePayload{})))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
