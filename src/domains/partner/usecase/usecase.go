package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	PartnerTypeRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, partner *models.PartnerPayloadCreate) (*models.Partner, error)
	}

	impl struct {
		repository            repository.RepositoryInterface
		partnerTypeRepository PartnerTypeRepository.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, partner *models.PartnerPayloadCreate) (*models.Partner, error) {
	trx := ctx.MysqlSess.Begin()

	// create partner
	id, _ := uuid.NewV4()
	payload := &models.Partner{
		ID:          id,
		Name:        partner.Name,
		Adress:      partner.Address,
		Contact:     partner.Address,
		Description: partner.Description,
	}
	result, err := i.repository.Create(trx, payload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	// add each partner_type into created partner
	resultTypes, err := i.partnerTypeRepository.AddTypeBatch(trx, &partner.Types)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	result.Types = resultTypes

	trx.Commit()
	return result, err
}

func NewUsecase(repo repository.RepositoryInterface, ptRepo PartnerTypeRepository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository:            repo,
		partnerTypeRepository: ptRepo,
	}
}
