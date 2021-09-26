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

	// #region domain material
	PartRest "github.com/hrz8/sc-masterlist-service/src/domains/part/delivery/rest"
	PartRepository "github.com/hrz8/sc-masterlist-service/src/domains/part/repository"
	PartUsecase "github.com/hrz8/sc-masterlist-service/src/domains/part/usecase"

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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
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
	partnerTypeRepo := PartnerTypeRepository.NewRepository(mysqlSess)
	partnerTypeUsecase := PartnerTypeUsecase.NewUsecase(partnerTypeRepo)
	partnerTypeRest := PartnerTypeRest.NewRest(partnerTypeUsecase)
	// - domain partner
	partnerRepo := PartnerRepository.NewRepository(mysqlSess)
	partnerUsecase := PartnerUsecase.NewUsecase(partnerRepo, partnerTypeRepo)
	partnerRest := PartnerRest.NewRest(partnerUsecase)
	// - domain grain_type
	grainTypeRepo := GrainTypeRepository.NewRepository(mysqlSess)
	grainTypeUsecase := GrainTypeUsecase.NewUsecase(grainTypeRepo)
	grainTypeRest := GrainTypeRest.NewRest(grainTypeUsecase)
	// - domain mould_cav
	mouldCavRepo := MouldCavRepository.NewRepository(mysqlSess)
	mouldCavUsecase := MouldCavUsecase.NewUsecase(mouldCavRepo)
	mouldCavRest := MouldCavRest.NewRest(mouldCavUsecase)
	// - domain mould_ton
	mouldTonRepo := MouldTonRepository.NewRepository(mysqlSess)
	mouldTonUsecase := MouldTonUsecase.NewUsecase(mouldTonRepo)
	mouldTonRest := MouldTonRest.NewRest(mouldTonUsecase)
	// - domain material_grade
	materialGradeRepo := MaterialGradeRepository.NewRepository(mysqlSess)
	materialGradeUsecase := MaterialGradeUsecase.NewUsecase(materialGradeRepo)
	materialGradeRest := MaterialGradeRest.NewRest(materialGradeUsecase)
	// - domain color
	colorRepo := ColorRepository.NewRepository(mysqlSess)
	colorUsecase := ColorUsecase.NewUsecase(colorRepo)
	colorRest := ColorRest.NewRest(colorUsecase)
	// - domain material
	materialRepo := MaterialRepository.NewRepository(mysqlSess)
	materialUsecase := MaterialUsecase.NewUsecase(
		materialRepo,
		materialGradeRepo,
		partnerRepo,
	)
	materialRest := MaterialRest.NewRest(materialUsecase)
	// - domain part
	partRepo := PartRepository.NewRepository(mysqlSess)
	partUsecase := PartUsecase.NewUsecase(
		partRepo,
		projectRepo,
		materialRepo,
		grainTypeRepo,
		mouldTonRepo,
		mouldCavRepo,
		processRepo,
	)
	partRest := PartRest.NewRest(partUsecase)
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
	// - domain part
	PartRest.AddPartEndpoints(e, partRest)
	// #endregion

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
