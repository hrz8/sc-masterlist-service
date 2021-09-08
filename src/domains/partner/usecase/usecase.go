package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
)

type (
	UsecaseInterface interface {
		Create(*models.PartnerPayloadCreate) (*models.Partner, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(partner *models.PartnerPayloadCreate) (*models.Partner, error) {
	id, _ := uuid.NewV4()
	payload := &models.Partner{
		ID:          id,
		Name:        partner.Name,
		Adress:      partner.Address,
		Contact:     partner.Address,
		Description: partner.Description,
	}
	result, err := i.repository.Create(payload)
	return result, err
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: r,
	}
}
