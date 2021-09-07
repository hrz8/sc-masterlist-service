package main

import (
	"fmt"

	// domain project
	ProjectRest "github.com/hrz8/sc-masterlist-service/src/domains/project/delivery/rest"
	ProjectRepository "github.com/hrz8/sc-masterlist-service/src/domains/project/repository"
	ProjectUsecase "github.com/hrz8/sc-masterlist-service/src/domains/project/usecase"

	// domain process
	ProcessRest "github.com/hrz8/sc-masterlist-service/src/domains/process/delivery/rest"
	ProcessRepository "github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	ProcessUsecase "github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"

	// domain sourcing
	SourcingRepository "github.com/hrz8/sc-masterlist-service/src/domains/sourcing/repository"
	SourcingUsecase "github.com/hrz8/sc-masterlist-service/src/domains/sourcing/usecase"

	Config "github.com/hrz8/sc-masterlist-service/src/shared/config"
	Container "github.com/hrz8/sc-masterlist-service/src/shared/container"
	Database "github.com/hrz8/sc-masterlist-service/src/shared/database"
	"github.com/hrz8/sc-masterlist-service/src/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appContainer := Container.NewAppContainer()
	appConfig := appContainer.MustGet("shared.config").(*Config.AppConfig)
	mysql := appContainer.MustGet("shared.mysql").(Database.MysqlInterface)

	mysqlSess := mysql.Connect()

	// #region services loader
	// - domain project
	projectRepo := ProjectRepository.NewRepository(mysqlSess)
	projectUsecase := ProjectUsecase.NewUsecase(projectRepo)
	// - domain process
	processRepo := ProcessRepository.NewRepository(mysqlSess)
	processUsecase := ProcessUsecase.NewUsecase(processRepo)
	// - domain sourcing
	sourcingRepo := SourcingRepository.NewRepository(mysqlSess)
	SourcingUsecase.NewUsecase(sourcingRepo)
	// #endregion

	// #region rest loader
	// - domain project
	projectRest := ProjectRest.NewRest(projectUsecase)
	// - domain process
	processRest := ProcessRest.NewRest(processUsecase)
	// #endregion

	// rest server
	e := echo.New()
	e.Validator = utils.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &utils.CustomContext{
				Context:   ctx,
				MysqlSess: mysqlSess,
			}
			return next(cc)
		}
	})

	// #region delivery endpoint implementation
	ProcessRest.AddProcessEndpoints(e, processRest)
	ProjectRest.AddProjectEndpoints(e, projectRest)
	// #endregion

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
