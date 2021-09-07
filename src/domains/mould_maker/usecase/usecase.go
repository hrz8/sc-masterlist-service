package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/mould_maker/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.MouldMakerPayloadCreate) (*models.MouldMaker, error)
		GetAll(*models.MouldMakerPayloadGetAll) (*[]models.MouldMaker, *int64, error)
		GetById(*uuid.UUID) (*models.MouldMaker, error)
		DeleteById(*uuid.UUID) (*models.MouldMaker, error)
		UpdateById(*uuid.UUID, *models.MouldMakerPayloadUpdateById) (*models.MouldMaker, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(mouldMaker *models.MouldMakerPayloadCreate) (*models.MouldMaker, error) {
	id, _ := uuid.NewV4()
	payload := &models.MouldMaker{
		ID:          id,
		Name:        mouldMaker.Name,
		Description: mouldMaker.Description,
	}
	result, err := i.repository.Create(payload)
	return result, err
}

func (i *impl) GetAll(conditions *models.MouldMakerPayloadGetAll) (*[]models.MouldMaker, *int64, error) {
	result, total, err := i.repository.GetAll(conditions)
	return result, total, err
}

func (i *impl) GetById(id *uuid.UUID) (*models.MouldMaker, error) {
	result, err := i.repository.GetById(id)
	return result, err
}

func (i *impl) DeleteById(id *uuid.UUID) (*models.MouldMaker, error) {
	instance, err := i.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(id *uuid.UUID, payload *models.MouldMakerPayloadUpdateById) (*models.MouldMaker, error) {
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
