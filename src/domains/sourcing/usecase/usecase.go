package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/sourcing/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.SourcingPayloadCreate) (*models.Sourcing, error)
		GetAll(*models.SourcingPayloadGetAll) (*[]models.Sourcing, *int64, error)
		GetById(*uuid.UUID) (*models.Sourcing, error)
		DeleteById(*uuid.UUID) (*models.Sourcing, error)
		UpdateById(*uuid.UUID, *models.SourcingPayloadUpdateById) (*models.Sourcing, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(sourcing *models.SourcingPayloadCreate) (*models.Sourcing, error) {
	id, _ := uuid.NewV4()
	payload := &models.Sourcing{
		ID:          id,
		Name:        sourcing.Name,
		Description: sourcing.Description,
	}
	result, err := i.repository.Create(payload)
	return result, err
}

func (i *impl) GetAll(conditions *models.SourcingPayloadGetAll) (*[]models.Sourcing, *int64, error) {
	result, total, err := i.repository.GetAll(conditions)
	return result, total, err
}

func (i *impl) GetById(id *uuid.UUID) (*models.Sourcing, error) {
	result, err := i.repository.GetById(id)
	return result, err
}

func (i *impl) DeleteById(id *uuid.UUID) (*models.Sourcing, error) {
	instance, err := i.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(id *uuid.UUID, payload *models.SourcingPayloadUpdateById) (*models.Sourcing, error) {
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
