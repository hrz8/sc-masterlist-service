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

	// #region domain graintype
	GrainTypeRest "github.com/hrz8/sc-masterlist-service/src/domains/grain_type/delivery/rest"
	GrainTypeRepository "github.com/hrz8/sc-masterlist-service/src/domains/grain_type/repository"
	GrainTypeUsecase "github.com/hrz8/sc-masterlist-service/src/domains/grain_type/usecase"

	// #endregion

	// #region domain mould_cav
	MouldCavRest "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/delivery/rest"
	MouldCavRepository "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/repository"
	MouldCavUsecase "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/usecase"

	// #endregion

	// #region domain mould_ton
	MouldTonRest "github.com/hrz8/sc-masterlist-service/src/domains/mould_ton/delivery/rest"
	MouldTonRepository "github.com/hrz8/sc-masterlist-service/src/domains/mould_ton/repository"
	MouldTonUsecase "github.com/hrz8/sc-masterlist-service/src/domains/mould_ton/usecase"

	// #endregion

	// #region domain material_grade
	MaterialGradeRest "github.com/hrz8/sc-masterlist-service/src/domains/material_grade/delivery/rest"
	MaterialGradeRepository "github.com/hrz8/sc-masterlist-service/src/domains/material_grade/repository"
	MaterialGradeUsecase "github.com/hrz8/sc-masterlist-service/src/domains/material_grade/usecase"

	// #endregion

	// #region domain color
	ColorRest "github.com/hrz8/sc-masterlist-service/src/domains/color/delivery/rest"
	ColorRepository "github.com/hrz8/sc-masterlist-service/src/domains/color/repository"
	ColorUsecase "github.com/hrz8/sc-masterlist-service/src/domains/color/usecase"

	// #endregion

	// #region domain material
	MaterialRest "github.com/hrz8/sc-masterlist-service/src/domains/material/delivery/rest"
	MaterialRepository "github.com/hrz8/sc-masterlist-service/src/domains/material/repository"
	MaterialUsecase "github.com/hrz8/sc-masterlist-service/src/domains/material/usecase"

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
	// - domain partner_type
	partnerTypeRepository := PartnerTypeRepository.NewRepository(mysqlSess)
	partnerTypeUsecase := PartnerTypeUsecase.NewUsecase(partnerTypeRepository)
	partnerTypeRest := PartnerTypeRest.NewRest(partnerTypeUsecase)
	// - domain partner
	partnerRepository := PartnerRepository.NewRepository(mysqlSess)
	partnerUsecase := PartnerUsecase.NewUsecase(partnerRepository, partnerTypeRepository)
	partnerRest := PartnerRest.NewRest(partnerUsecase)
	// - domain grain_type
	grainTypeRepository := GrainTypeRepository.NewRepository(mysqlSess)
	grainTypeUsecase := GrainTypeUsecase.NewUsecase(grainTypeRepository)
	grainTypeRest := GrainTypeRest.NewRest(grainTypeUsecase)
	// - domain mould_cav
	mouldCavRepository := MouldCavRepository.NewRepository(mysqlSess)
	mouldCavUsecase := MouldCavUsecase.NewUsecase(mouldCavRepository)
	mouldCavRest := MouldCavRest.NewRest(mouldCavUsecase)
	// - domain mould_ton
	mouldTonRepository := MouldTonRepository.NewRepository(mysqlSess)
	mouldTonUsecase := MouldTonUsecase.NewUsecase(mouldTonRepository)
	mouldTonRest := MouldTonRest.NewRest(mouldTonUsecase)
	// - domain material_grade
	materialGradeRepository := MaterialGradeRepository.NewRepository(mysqlSess)
	materialGradeUsecase := MaterialGradeUsecase.NewUsecase(materialGradeRepository)
	materialGradeRest := MaterialGradeRest.NewRest(materialGradeUsecase)
	// - domain color
	colorRepository := ColorRepository.NewRepository(mysqlSess)
	colorUsecase := ColorUsecase.NewUsecase(colorRepository)
	colorRest := ColorRest.NewRest(colorUsecase)
	//- domain material
	materialRepository := MaterialRepository.NewRepository(mysqlSess)
	materialUsecase := MaterialUsecase.NewUsecase(materialRepository)
	materialRest := MaterialRest.NewRest(materialUsecase)
	// #endregion

	// #region delivery endpoint implementation
	// - domain project
	ProcessRest.AddProcessEndpoints(e, processRest)
	// - domain process
	ProjectRest.AddProjectEndpoints(e, projectRest)
	// - domain partner_type
	PartnerTypeRest.AddPartnerTypeEndpoints(e, partnerTypeRest)
	// - domain partner
	PartnerRest.AddPartnerEndpoints(e, partnerRest)
	// - domain grain_type
	GrainTypeRest.AddGrainTypeEndpoints(e, grainTypeRest)
	// - domain mould_cav
	MouldCavRest.AddMouldCavEndpoints(e, mouldCavRest)
	// - domain mould_ton
	MouldTonRest.AddMouldTonEndpoints(e, mouldTonRest)
	// - domain material_grade
	MaterialGradeRest.AddMaterialGradeEndpoints(e, materialGradeRest)
	// - domain color
	ColorRest.AddColorEndpoints(e, colorRest)
	// - domain color
	MaterialRest.AddMaterialEndpoints(e, materialRest)
	// #endregion

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
