package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.PartnerTypePayloadCreate) (*models.PartnerType, error)
		GetAll(*models.PartnerTypePayloadGetAll) (*[]models.PartnerType, *int64, error)
		GetById(*uuid.UUID) (*models.PartnerType, error)
		DeleteById(*uuid.UUID) (*models.PartnerType, error)
		UpdateById(*uuid.UUID, *models.PartnerTypePayloadUpdateById) (*models.PartnerType, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(partnerType *models.PartnerTypePayloadCreate) (*models.PartnerType, error) {
	id, _ := uuid.NewV4()
	payload := &models.PartnerType{
		ID:          id,
		Name:        partnerType.Name,
		Description: partnerType.Description,
	}
	result, err := i.repository.Create(payload)
	return result, err
}

func (i *impl) GetAll(conditions *models.PartnerTypePayloadGetAll) (*[]models.PartnerType, *int64, error) {
	result, total, err := i.repository.GetAll(conditions)
	return result, total, err
}

func (i *impl) GetById(id *uuid.UUID) (*models.PartnerType, error) {
	result, err := i.repository.GetById(id)
	return result, err
}

func (i *impl) DeleteById(id *uuid.UUID) (*models.PartnerType, error) {
	instance, err := i.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(id *uuid.UUID, payload *models.PartnerTypePayloadUpdateById) (*models.PartnerType, error) {
	instance, err := i.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	result, err := i.repository.Update(instance, payload)
	return result, err
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: r,
	}
}
