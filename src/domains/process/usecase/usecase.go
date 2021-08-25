package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/process/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.ProcessCreatePayload) (*models.Process, error)
	}

	impl struct {
		repo repository.RepositoryInterface
	}
)

func (i *impl) Create(process *models.ProcessCreatePayload) (*models.Process, error) {
	id, _ := uuid.NewV4()
	payload := &models.Process{
		ID:          id,
		Name:        process.Name,
		Description: process.Description,
	}
	result, _ := i.repo.Create(payload)
	return result, nil
}

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repo: repo,
	}
}
