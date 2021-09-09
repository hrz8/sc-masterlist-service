package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	PartnerTypeDomain "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(*utils.CustomContext, *models.PartnerPayloadCreate) (*models.Partner, error)
	}

	impl struct {
		repository            repository.RepositoryInterface
		partnerTypeRepository PartnerTypeDomain.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, partner *models.PartnerPayloadCreate) (*models.Partner, error) {
	trx := ctx.MysqlSess.Begin()

	id, _ := uuid.NewV4()
	types := make([]*models.PartnerType, len(partner.Types))
	for i, item := range partner.Types {
		types[i] = &models.PartnerType{
			ID: item,
		}
	}
	payload := &models.Partner{
		ID:          id,
		Name:        partner.Name,
		Adress:      partner.Address,
		Contact:     partner.Address,
		Description: partner.Description,
		Types:       types,
	}
	result, err := i.repository.Create(trx, payload)

	trx.Commit()
	return result, err
}

func NewUsecase(repo repository.RepositoryInterface, ptRepo PartnerTypeDomain.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository:            repo,
		partnerTypeRepository: ptRepo,
	}
}
