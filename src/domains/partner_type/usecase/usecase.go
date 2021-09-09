package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(*utils.CustomContext, *models.PartnerTypePayloadCreate) (*models.PartnerType, error)
		GetAll(*utils.CustomContext, *models.PartnerTypePayloadGetAll) (*[]models.PartnerType, *int64, error)
		GetById(*utils.CustomContext, *uuid.UUID) (*models.PartnerType, error)
		DeleteById(*utils.CustomContext, *uuid.UUID) (*models.PartnerType, error)
		UpdateById(*utils.CustomContext, *uuid.UUID, *models.PartnerTypePayloadUpdateById) (*models.PartnerType, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(_ *utils.CustomContext, partnerType *models.PartnerTypePayloadCreate) (*models.PartnerType, error) {
	id, _ := uuid.NewV4()
	payload := &models.PartnerType{
		ID:          id,
		Name:        partnerType.Name,
		Description: partnerType.Description,
	}
	result, err := i.repository.Create(nil, payload)
	return result, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.PartnerTypePayloadGetAll) (*[]models.PartnerType, *int64, error) {
	result, total, err := i.repository.GetAll(nil, conditions)
	return result, total, err
}

func (i *impl) GetById(_ *utils.CustomContext, id *uuid.UUID) (*models.PartnerType, error) {
	result, err := i.repository.GetById(nil, id)
	return result, err
}

func (i *impl) DeleteById(_ *utils.CustomContext, id *uuid.UUID) (*models.PartnerType, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(nil, id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(_ *utils.CustomContext, id *uuid.UUID, payload *models.PartnerTypePayloadUpdateById) (*models.PartnerType, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	result, err := i.repository.Update(nil, instance, payload)
	return result, err
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: r,
	}
}
