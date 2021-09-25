package usecase

import (
	"github.com/gofrs/uuid"
	ColorError "github.com/hrz8/sc-masterlist-service/src/domains/color/error"
	ColorRepository "github.com/hrz8/sc-masterlist-service/src/domains/color/repository"
	GrainTypeError "github.com/hrz8/sc-masterlist-service/src/domains/grain_type/error"
	GrainTypeRepository "github.com/hrz8/sc-masterlist-service/src/domains/grain_type/repository"
	MaterialError "github.com/hrz8/sc-masterlist-service/src/domains/material/error"
	MaterialRepository "github.com/hrz8/sc-masterlist-service/src/domains/material/repository"
	MouldCavError "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/error"
	MouldCavRepository "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/repository"
	MouldTonError "github.com/hrz8/sc-masterlist-service/src/domains/mould_ton/error"
	MouldTonRepository "github.com/hrz8/sc-masterlist-service/src/domains/mould_ton/repository"
	PartError "github.com/hrz8/sc-masterlist-service/src/domains/part/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/part/repository"
	PartnerError "github.com/hrz8/sc-masterlist-service/src/domains/partner/error"
	PartnerRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	PartnerTypeError "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/error"
	ProcessRepository "github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	ProjectError "github.com/hrz8/sc-masterlist-service/src/domains/project/error"
	ProjectRepository "github.com/hrz8/sc-masterlist-service/src/domains/project/repository"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, part *models.PartPayloadCreate) (*models.Part, error)
	}

	impl struct {
		repository          repository.RepositoryInterface
		projectRepository   ProjectRepository.RepositoryInterface
		materialRepository  MaterialRepository.RepositoryInterface
		grainTypeRepository GrainTypeRepository.RepositoryInterface
		mouldTonRepository  MouldTonRepository.RepositoryInterface
		mouldCavRepository  MouldCavRepository.RepositoryInterface
		processRepository   ProcessRepository.RepositoryInterface
		colorRepository     ColorRepository.RepositoryInterface
		partnerRepository   PartnerRepository.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, part *models.PartPayloadCreate) (*models.Part, error) {
	trx := ctx.MysqlSess.Begin()
	id, _ := uuid.NewV4()

	// required relation
	projectInstance, err := i.projectRepository.GetById(trx, &part.Project)
	if err != nil {
		trx.Rollback()
		return nil, ProjectError.GetById.Err
	}

	// payload
	payload := &models.Part{
		ID:               id,
		Number:           part.Number,
		Name:             part.Name,
		Image:            part.Image,
		QtyPerUnit:       part.QtyPerUnit,
		QtyPerMonth:      part.QtyPerMonth,
		DwgWeight:        part.DwgWeight,
		ActualWeightPart: part.ActualWeightPart,
		ActualWeightRun:  part.ActualWeightRun,
		PaintColor:       part.PaintColor,
		PaintCode:        part.PaintCode,
		Remarks:          part.Remarks,
		SourcingRemarks:  part.SourcingRemarks,
		ProcessRouting:   part.ProcessRouting,
		// has one - required
		ProjectID: projectInstance.ID,
	}

	// optional foreign key definition
	var parentInstance *models.Part
	var materialInstance *models.Material
	var grainTypeInstance *models.GrainType
	var mouldTonInstance *models.MouldTon
	var mouldCavInstance *models.MouldCav

	// get Parent Part by its uuid to check if its available
	if !helpers.IsEmptyUUID(&part.Parent) {
		parentInstance, err = i.repository.GetById(trx, &part.Parent)
		if err != nil {
			trx.Rollback()
			return nil, PartError.GetById.Err
		}
		payload.ParentID = &parentInstance.ID
	}

	// get Material by its uuid to check if its available
	if !helpers.IsEmptyUUID(&part.Material) {
		materialInstance, err = i.materialRepository.GetById(trx, &part.Material)
		if err != nil {
			trx.Rollback()
			return nil, MaterialError.GetById.Err
		}
		payload.MaterialID = &materialInstance.ID
	}

	// get GrainType by its uuid to check if its available
	if !helpers.IsEmptyUUID(&part.GrainType) {
		grainTypeInstance, err = i.grainTypeRepository.GetById(trx, &part.GrainType)
		if err != nil {
			trx.Rollback()
			return nil, GrainTypeError.GetById.Err
		}
		payload.GrainTypeID = &grainTypeInstance.ID
	}

	// get MouldTon by its uuid to check if its available
	if !helpers.IsEmptyUUID(&part.MouldTon) {
		mouldTonInstance, err = i.mouldTonRepository.GetById(trx, &part.MouldTon)
		if err != nil {
			trx.Rollback()
			return nil, MouldTonError.GetById.Err
		}
		payload.MouldTonID = &mouldTonInstance.ID
	}

	// get MouldCav by its uuid to check if its available
	if !helpers.IsEmptyUUID(&part.MouldCav) {
		mouldCavInstance, err = i.mouldCavRepository.GetById(trx, &part.MouldCav)
		if err != nil {
			trx.Rollback()
			return nil, MouldCavError.GetById.Err
		}
		payload.MouldCavID = &mouldCavInstance.ID
	}

	partCreated, err := i.repository.Create(trx, payload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	// check if processes payload not empty
	if len(part.Processes) > 0 {
		// get each processes by its uuid to check if its available
		processes := make([]*models.Process, len(part.Processes))
		for index, partnerTypeID := range part.Processes {
			process, err := i.processRepository.GetById(trx, &partnerTypeID)
			if err != nil {
				trx.Rollback()
				return nil, PartnerTypeError.GetById.Err
			}
			processes[index] = process
		}

		// add all process into created part
		if err := trx.Debug().Model(&partCreated).
			Association("Process").Append(processes); err != nil {
			trx.Rollback()
			return nil, err
		}
	}

	// check if colors payload not empty
	if len(part.Colors) > 0 {
		// get each colors by its uuid to check if its available
		colors := make([]*models.Color, len(part.Colors))
		for index, colorID := range part.Colors {
			color, err := i.colorRepository.GetById(trx, &colorID)
			if err != nil {
				trx.Rollback()
				return nil, ColorError.GetById.Err
			}
			colors[index] = color
		}

		// add all process into created part
		if err := trx.Debug().Model(&partCreated).
			Association("Colors").Append(colors); err != nil {
			trx.Rollback()
			return nil, err
		}
	}

	// check if sourcings payload not empty
	if len(part.Sourcings) > 0 {
		// get each sourcings by its uuid to check if its available
		sourcings := make([]*models.Partner, len(part.Sourcings))
		for index, sourcingID := range part.Sourcings {
			sourcing, err := i.partnerRepository.GetById(trx, &sourcingID)
			if err != nil {
				trx.Rollback()
				return nil, PartnerError.GetById.Err
			}
			sourcings[index] = sourcing
		}

		// add all sourcings into created part
		if err := trx.Debug().Model(&partCreated).
			Association("Sourcings").Append(sourcings); err != nil {
			trx.Rollback()
			return nil, err
		}
	}

	// check if mould_makers payload not empty
	if len(part.MouldMakers) > 0 {
		// get each mould_makers by its uuid to check if its available
		mouldMakers := make([]*models.Partner, len(part.MouldMakers))
		for index, mouldMakerID := range part.MouldMakers {
			mouldMaker, err := i.partnerRepository.GetById(trx, &mouldMakerID)
			if err != nil {
				trx.Rollback()
				return nil, PartnerError.GetById.Err
			}
			mouldMakers[index] = mouldMaker
		}

		// add all mould_makers into created part
		if err := trx.Debug().Model(&partCreated).
			Association("MouldMakers").Append(mouldMakers); err != nil {
			trx.Rollback()
			return nil, err
		}
	}

	trx.Commit()
	// required
	partCreated.Project = *projectInstance
	// optional
	partCreated.Parent = parentInstance
	partCreated.Material = materialInstance
	partCreated.GrainType = grainTypeInstance
	partCreated.MouldTon = mouldTonInstance
	partCreated.MouldCav = mouldCavInstance
	return partCreated, err
}

func NewUsecase(
	repo repository.RepositoryInterface,
	projectRepo ProjectRepository.RepositoryInterface,
	materialRepo MaterialRepository.RepositoryInterface,
	grainTypeRepo GrainTypeRepository.RepositoryInterface,
	mouldTonRepo MouldTonRepository.RepositoryInterface,
	mouldCavRepo MouldCavRepository.RepositoryInterface,
	processRepo ProcessRepository.RepositoryInterface,
) UsecaseInterface {
	return &impl{
		repository:          repo,
		projectRepository:   projectRepo,
		materialRepository:  materialRepo,
		grainTypeRepository: grainTypeRepo,
		mouldTonRepository:  mouldTonRepo,
		mouldCavRepository:  mouldCavRepo,
		processRepository:   processRepo,
	}
}
