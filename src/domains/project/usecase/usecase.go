package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/project/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.ProjectPayloadCreate) (*models.Project, error)
		GetAll(*models.ProjectPayloadGetAll) (*[]models.Project, error)
		GetById(*uuid.UUID) (*models.Project, error)
		DeleteById(*uuid.UUID) (*models.Project, error)
		UpdateById(*uuid.UUID, *models.ProjectPayloadUpdateById) (*models.Project, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(project *models.ProjectPayloadCreate) (*models.Project, error) {
	id, _ := uuid.NewV4()
	payload := &models.Project{
		ID:          id,
		Name:        project.Name,
		Description: project.Description,
	}
	result, err := i.repository.Create(payload)
	return result, err
}

func (i *impl) GetAll(conditions *models.ProjectPayloadGetAll) (*[]models.Project, error) {
	result, err := i.repository.GetAll(conditions)
	return result, err
}

func (i *impl) GetById(id *uuid.UUID) (*models.Project, error) {
	result, err := i.repository.GetById(id)
	return result, err
}

func (i *impl) DeleteById(id *uuid.UUID) (*models.Project, error) {
	instance, err := i.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(id *uuid.UUID, payload *models.ProjectPayloadUpdateById) (*models.Project, error) {
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
