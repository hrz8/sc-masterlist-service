package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.ProcessPayloadCreate) (*models.Process, error)
		GetAll(*models.ProcessPayloadGetAll) (*[]models.Process, *int64, error)
		GetById(*uuid.UUID) (*models.Process, error)
		DeleteById(*uuid.UUID) (*models.Process, error)
		UpdateById(*uuid.UUID, *models.ProcessPayloadUpdateById) (*models.Process, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(process *models.ProcessPayloadCreate) (*models.Process, error) {
	id, _ := uuid.NewV4()
	payload := &models.Process{
		ID:          id,
		Name:        process.Name,
		Description: process.Description,
	}
	result, err := i.repository.Create(nil, payload)
	return result, err
}

func (i *impl) GetAll(conditions *models.ProcessPayloadGetAll) (*[]models.Process, *int64, error) {
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

func (i *impl) GetById(id *uuid.UUID) (*models.Process, error) {
	result, err := i.repository.GetById(nil, id)
	return result, err
}

func (i *impl) DeleteById(id *uuid.UUID) (*models.Process, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(nil, id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(id *uuid.UUID, payload *models.ProcessPayloadUpdateById) (*models.Process, error) {
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
