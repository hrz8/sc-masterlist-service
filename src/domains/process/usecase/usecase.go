package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.ProcessPayloadCreate) (*models.Process, error)
		GetAll(*models.ProcessPayloadGetAll) (*[]models.Process, error)
		Get(*string) (*models.Process, error)
		Delete(*string) (*models.Process, error)
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
	result, err := i.repository.Create(payload)
	return result, err
}

func (i *impl) GetAll(c *models.ProcessPayloadGetAll) (*[]models.Process, error) {
	result, err := i.repository.GetAll(c)
	return result, err
}

func (i *impl) Get(id *string) (*models.Process, error) {
	result, err := i.repository.Get(id)
	return result, err
}

func (i *impl) Delete(id *string) (*models.Process, error) {
	result, err := i.repository.Delete(id)
	return result, err
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: r,
	}
}
