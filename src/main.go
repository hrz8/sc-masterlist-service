package main

import (
	"fmt"

	// #region domain project
	ProjectRest "github.com/hrz8/sc-masterlist-service/src/domains/project/delivery/rest"
	ProjectRepository "github.com/hrz8/sc-masterlist-service/src/domains/project/repository"
	ProjectUsecase "github.com/hrz8/sc-masterlist-service/src/domains/project/usecase"

	// #endregion

	// #region domain process
	ProcessRest "github.com/hrz8/sc-masterlist-service/src/domains/process/delivery/rest"
	ProcessRepository "github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	ProcessUsecase "github.com/hrz8/sc-masterlist-service/src/domains/process/usecase"

	// #endregion

	// #region domain sourcing
	SourcingRest "github.com/hrz8/sc-masterlist-service/src/domains/sourcing/delivery/rest"
	SourcingRepository "github.com/hrz8/sc-masterlist-service/src/domains/sourcing/repository"
	SourcingUsecase "github.com/hrz8/sc-masterlist-service/src/domains/sourcing/usecase"

	// #endregion

	// #region domain mould_maker
	MouldMakerRepository "github.com/hrz8/sc-masterlist-service/src/domains/mould_maker/repository"
	// #endregion

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
	sourcingUsecase := SourcingUsecase.NewUsecase(sourcingRepo)
	// - domain mould_maker
	MouldMakerRepository.NewRepository(mysqlSess)
	// #endregion

	// #region rest loader
	// - domain project
	projectRest := ProjectRest.NewRest(projectUsecase)
	// - domain process
	processRest := ProcessRest.NewRest(processUsecase)
	// - domain sourcing
	sourcingRest := SourcingRest.NewRest(sourcingUsecase)
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
	// - domain project
	ProcessRest.AddProcessEndpoints(e, processRest)
	// - domain process
	ProjectRest.AddProjectEndpoints(e, projectRest)
	// - domain sourcing
	SourcingRest.AddSourcingEndpoints(e, sourcingRest)
	// #endregion

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
