package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	PartnerTypeError "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/error"
	PartnerTypeRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, partner *models.PartnerPayloadCreate) (*models.Partner, error)
		GetAll(ctx *utils.CustomContext, conditions *models.PartnerPayloadGetAll) (*[]models.Partner, *int64, error)
		GetById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Partner, error)
		DeleteById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Partner, error)
	}

	impl struct {
		repository            repository.RepositoryInterface
		partnerTypeRepository PartnerTypeRepository.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, partner *models.PartnerPayloadCreate) (*models.Partner, error) {
	trx := ctx.MysqlSess.Begin()

	// create partner
	id, _ := uuid.NewV4()
	payload := &models.Partner{
		ID:          id,
		Name:        partner.Name,
		Adress:      partner.Address,
		Contact:     partner.Contact,
		Description: partner.Description,
	}
	partnerCreated, err := i.repository.Create(trx, payload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	// get each partnerTypes by its uuid to check if its available
	partnerTypes := make([]*models.PartnerType, len(partner.PartnerTypes))
	for index, partnerTypeID := range partner.PartnerTypes {
		partnerType, err := i.partnerTypeRepository.GetById(trx, &partnerTypeID)
		if err != nil {
			trx.Rollback()
			return nil, PartnerTypeError.GetById.Err
		}
		partnerTypes[index] = partnerType
	}

	// add each partner_type into created partner
	if err := trx.Debug().Model(&partnerCreated).Association("PartnerTypes").Append(partnerTypes); err != nil {
		return nil, err
	}

	trx.Commit()
	return partnerCreated, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.PartnerPayloadGetAll) (*[]models.Partner, *int64, error) {
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

func (i *impl) GetById(_ *utils.CustomContext, id *uuid.UUID) (*models.Partner, error) {
	result, err := i.repository.GetById(nil, id)
	return result, err
}

func (i *impl) DeleteById(_ *utils.CustomContext, id *uuid.UUID) (*models.Partner, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(nil, id); err != nil {
		return nil, err
	}
	return instance, nil
}

func NewUsecase(repo repository.RepositoryInterface, ptRepo PartnerTypeRepository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository:            repo,
		partnerTypeRepository: ptRepo,
	}
}
