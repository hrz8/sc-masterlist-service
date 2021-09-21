package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/material/repository"
	MaterialGradeError "github.com/hrz8/sc-masterlist-service/src/domains/material_grade/error"
	MaterialGradeRepository "github.com/hrz8/sc-masterlist-service/src/domains/material_grade/repository"
	MakerError "github.com/hrz8/sc-masterlist-service/src/domains/partner/error"
	MakerRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, material *models.MaterialPayloadCreate) (*models.Material, error)
		GetAll(ctx *utils.CustomContext, conditions *models.MaterialPayloadGetAll) (*[]models.Material, *int64, error)
		GetById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Material, error)
		DeleteById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Material, error)
		UpdateById(ctx *utils.CustomContext, id *uuid.UUID, payload *models.MaterialPayloadUpdateById) (*models.Material, error)
	}

	impl struct {
		repository              repository.RepositoryInterface
		materialGradeRepository MaterialGradeRepository.RepositoryInterface
		makerRepository         MakerRepository.RepositoryInterface
	}
)

// Create is a facade function to implemented model creation
func (i *impl) Create(ctx *utils.CustomContext, material *models.MaterialPayloadCreate) (*models.Material, error) {
	trx := ctx.MysqlSess.Begin()

	id, _ := uuid.NewV4()

	// get MaterialGrade by its uuid to check if its available
	materialGrade, err := i.materialGradeRepository.GetById(trx, &material.MaterialGrade)
	if err != nil {
		trx.Rollback()
		return nil, MaterialGradeError.GetById.Err
	}

	// get Maker by its uuid to check if its available
	maker, err := i.makerRepository.GetById(trx, &material.Maker)
	if err != nil {
		trx.Rollback()
		return nil, MakerError.GetById.Err
	}

	payload := &models.Material{
		ID:              id,
		Tsm:             material.Tsm,
		Description:     material.Description,
		MaterialGradeID: materialGrade.ID,
		MakerID:         maker.ID,
	}
	materialCreated, err := i.repository.Create(trx, payload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	materialCreated.MaterialGrade = *materialGrade
	materialCreated.Maker = *maker
	materialCreated.Maker.PartnerTypes = nil
	return materialCreated, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.MaterialPayloadGetAll) (*[]models.Material, *int64, error) {
	result, err := i.repository.GetAll(nil, conditions)
	if err != nil {
		return nil, nil, err
	}
	total, err := i.repository.CountAll(nil)
	if err != nil {
		return nil, nil, err
	}
	return result, total, err
}

func (i *impl) GetById(_ *utils.CustomContext, id *uuid.UUID) (*models.Material, error) {
	result, err := i.repository.GetById(nil, id)
	return result, err
}

func (i *impl) DeleteById(_ *utils.CustomContext, id *uuid.UUID) (*models.Material, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(nil, id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(_ *utils.CustomContext, id *uuid.UUID, payload *models.MaterialPayloadUpdateById) (*models.Material, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	result, err := i.repository.Update(nil, instance, payload)
	return result, err
}

func NewUsecase(
	r repository.RepositoryInterface,
	mgRepo MaterialGradeRepository.RepositoryInterface,
	mRepo MakerRepository.RepositoryInterface,
) UsecaseInterface {
	return &impl{
		repository:              r,
		materialGradeRepository: mgRepo,
		makerRepository:         mRepo,
	}
}
