package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/color/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, color *models.ColorPayloadCreate) (*models.Color, error)
		GetAll(ctx *utils.CustomContext, conditions *models.ColorPayloadGetAll) (*[]models.Color, *int64, error)
		GetById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Color, error)
		DeleteById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Color, error)
		UpdateById(ctx *utils.CustomContext, id *uuid.UUID, payload *models.ColorPayloadUpdateById) (*models.Color, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(_ *utils.CustomContext, color *models.ColorPayloadCreate) (*models.Color, error) {
	id, _ := uuid.NewV4()
	payload := &models.Color{
		ID:          id,
		Code:        color.Code,
		Description: color.Description,
	}
	result, err := i.repository.Create(nil, payload)
	return result, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.ColorPayloadGetAll) (*[]models.Color, *int64, error) {
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

func (i *impl) GetById(_ *utils.CustomContext, id *uuid.UUID) (*models.Color, error) {
	result, err := i.repository.GetById(nil, id)
	return result, err
}

func (i *impl) DeleteById(_ *utils.CustomContext, id *uuid.UUID) (*models.Color, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(nil, id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(_ *utils.CustomContext, id *uuid.UUID, payload *models.ColorPayloadUpdateById) (*models.Color, error) {
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
