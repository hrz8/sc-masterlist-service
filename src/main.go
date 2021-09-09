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
	MouldMakerRest "github.com/hrz8/sc-masterlist-service/src/domains/mould_maker/delivery/rest"
	MouldMakerRepository "github.com/hrz8/sc-masterlist-service/src/domains/mould_maker/repository"
	MouldMakerUsecase "github.com/hrz8/sc-masterlist-service/src/domains/mould_maker/usecase"

	// #endregion

	// #region domain partner
	PartnerRest "github.com/hrz8/sc-masterlist-service/src/domains/partner/delivery/rest"
	PartnerRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	PartnerUsecase "github.com/hrz8/sc-masterlist-service/src/domains/partner/usecase"

	// #endregion

	// #region domain partner_type
	PartnerTypeRest "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/delivery/rest"
	PartnerTypeRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	PartnerTypeUsecase "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/usecase"

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

	// #region services loader
	// - domain project
	projectRepo := ProjectRepository.NewRepository(mysqlSess)
	projectUsecase := ProjectUsecase.NewUsecase(projectRepo)
	projectRest := ProjectRest.NewRest(projectUsecase)
	// - domain process
	processRepo := ProcessRepository.NewRepository(mysqlSess)
	processUsecase := ProcessUsecase.NewUsecase(processRepo)
	processRest := ProcessRest.NewRest(processUsecase)
	// - domain sourcing
	sourcingRepo := SourcingRepository.NewRepository(mysqlSess)
	sourcingUsecase := SourcingUsecase.NewUsecase(sourcingRepo)
	sourcingRest := SourcingRest.NewRest(sourcingUsecase)
	// - domain mould_maker
	mouldMakerRepository := MouldMakerRepository.NewRepository(mysqlSess)
	mouldMakerUsecase := MouldMakerUsecase.NewUsecase(mouldMakerRepository)
	mouldMakerRest := MouldMakerRest.NewRest(mouldMakerUsecase)
	// - domain partner_type
	partnerTypeRepository := PartnerTypeRepository.NewRepository(mysqlSess)
	partnerTypeUsecase := PartnerTypeUsecase.NewUsecase(partnerTypeRepository)
	partnerTypeRest := PartnerTypeRest.NewRest(partnerTypeUsecase)
	// - domain partner
	partnerRepository := PartnerRepository.NewRepository(mysqlSess)
	partnerUsecase := PartnerUsecase.NewUsecase(partnerRepository, partnerTypeRepository)
	partnerRest := PartnerRest.NewRest(partnerUsecase)
	// #endregion

	// #region delivery endpoint implementation
	// - domain project
	ProcessRest.AddProcessEndpoints(e, processRest)
	// - domain process
	ProjectRest.AddProjectEndpoints(e, projectRest)
	// - domain sourcing
	SourcingRest.AddSourcingEndpoints(e, sourcingRest)
	// - domain mould_maker
	MouldMakerRest.AddMouldMakerEndpoints(e, mouldMakerRest)
	// - domain partner_type
	PartnerTypeRest.AddPartnerTypeEndpoints(e, partnerTypeRest)
	// - domain partner
	PartnerRest.AddPartnerEndpoints(e, partnerRest)
	// #endregion

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
