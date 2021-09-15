package usecase

import (
	"fmt"

	"github.com/gofrs/uuid"
	PartnerError "github.com/hrz8/sc-masterlist-service/src/domains/partner/error"
	"github.com/hrz8/sc-masterlist-service/src/domains/partner/repository"
	PartnerTypeError "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/error"
	PartnerTypeRepository "github.com/hrz8/sc-masterlist-service/src/domains/partner_type/repository"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, partner *models.PartnerPayloadCreate) (*models.Partner, error)
		GetAll(ctx *utils.CustomContext, conditions *models.PartnerPayloadGetAll) (*[]models.Partner, *int64, error)
		GetById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Partner, error)
		DeleteById(ctx *utils.CustomContext, id *uuid.UUID) (*models.Partner, error)
		UpdateById(
			ctx *utils.CustomContext,
			id *uuid.UUID,
			payload *models.PartnerPayloadUpdateById,
		) (*models.Partner, error)
		AddPartnerType(
			ctx *utils.CustomContext,
			id *uuid.UUID,
			partnerTypeID *uuid.UUID,
		) (*models.Partner, error)
		DeletePartnerType(
			ctx *utils.CustomContext,
			id *uuid.UUID,
			partnerTypeID *uuid.UUID,
		) (*models.Partner, error)
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
		Address:     partner.Address,
		Contact:     partner.Contact,
		Description: partner.Description,
	}
	partnerCreated, err := i.repository.Create(trx, payload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	// check if partner_types payload not empty
	if len(partner.PartnerTypes) <= 0 {
		return nil, PartnerError.CreateWithEmptyPartnerTypes.Err
	}

	// get each partnerTypes by its uuid to check if its available
	partnerTypes := make([]*models.PartnerType, len(partner.PartnerTypes))
	for index, partnerTypeID := range partner.PartnerTypes {
		partnerType, err := i.partnerTypeRepository.GetById(trx, &partnerTypeID)
		if err != nil {
			trx.Rollback()
			return nil, PartnerTypeError.GetById.Err
		}
		partnerTypes[index] = partnerType
	}

	// add each partner_type into created partner
	if err := trx.Debug().Model(&partnerCreated).Association("PartnerTypes").Append(partnerTypes); err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return partnerCreated, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.PartnerPayloadGetAll) (*[]models.Partner, *int64, error) {
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

func (i *impl) GetById(_ *utils.CustomContext, id *uuid.UUID) (*models.Partner, error) {
	result, err := i.repository.GetById(nil, id)
	return result, err
}

func (i *impl) DeleteById(_ *utils.CustomContext, id *uuid.UUID) (*models.Partner, error) {
	instance, err := i.repository.GetById(nil, id)
	if err != nil {
		return nil, err
	}
	if err := i.repository.DeleteById(nil, id); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) UpdateById(
	ctx *utils.CustomContext,
	id *uuid.UUID,
	payload *models.PartnerPayloadUpdateById,
) (*models.Partner, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetById(trx, id)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	// update agnostic columns
	result, err := i.repository.Update(trx, instance, payload)

	// associating column
	if len(payload.PartnerTypes) > 0 {

		// get partnerType to be added
		var partnerTypesToBeAdd []*models.PartnerType
		for _, partnerTypeID := range payload.PartnerTypes {
			if helpers.SliceOfStructContainsFieldValue(instance.PartnerTypes, "ID", partnerTypeID) {
				continue
			}
			partnerType, err := i.partnerTypeRepository.GetById(trx, &partnerTypeID)
			if err != nil {
				trx.Rollback()
				return nil, PartnerTypeError.GetById.Err
			}

			// exclude partners asscoiate
			partnerType.Partners = nil
			partnerTypesToBeAdd = append(partnerTypesToBeAdd, partnerType)
		}

		// appending un-related-yet partnerType
		if err := trx.Model(&instance).Association("PartnerTypes").
			Append(partnerTypesToBeAdd); err != nil {
			trx.Rollback()
			return nil, err
		}

		for _, partnerType := range partnerTypesToBeAdd {
			partnerIDStr := fmt.Sprintf("%v", *id)
			partnerTypeID := fmt.Sprintf("%v", partnerType.ID)
			trx.Debug().Unscoped().Model(&models.PartnersPartnerTypes{}).
				Where("partner_id = ?", partnerIDStr).
				Where("partner_type_id = ?", partnerTypeID).
				Update("deleted_at", nil)
		}

		// get partnerType to be remove
		var partnerTypesToBeRemove []*models.PartnerType
		for _, instancePartnerType := range instance.PartnerTypes {
			if _, exists := helpers.SliceContains(payload.PartnerTypes, instancePartnerType.ID); exists {
				continue
			}

			// exclude partners asscoiate
			instancePartnerType.Partners = nil
			partnerTypesToBeRemove = append(partnerTypesToBeRemove, instancePartnerType)
		}

		// deleting related-yet partnerType
		if err := trx.Model(&instance).Association("PartnerTypes").
			Delete(partnerTypesToBeRemove); err != nil {
			trx.Rollback()
			return nil, err
		}
	}

	trx.Commit()
	return result, err
}

func (i *impl) AddPartnerType(
	ctx *utils.CustomContext,
	id *uuid.UUID,
	partnerTypeID *uuid.UUID,
) (*models.Partner, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetById(trx, id)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	partnerTypeInstance, err := i.partnerTypeRepository.GetById(trx, partnerTypeID)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	result, err := i.repository.AddPartnerType(trx, instance, partnerTypeInstance)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return result, nil
}

func (i *impl) DeletePartnerType(
	ctx *utils.CustomContext,
	id *uuid.UUID,
	partnerTypeID *uuid.UUID,
) (*models.Partner, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetById(trx, id)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	partnerTypeInstance, err := i.partnerTypeRepository.GetById(trx, partnerTypeID)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	result, err := i.repository.DeletePartnerType(trx, instance, partnerTypeInstance)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return result, nil
}

func NewUsecase(repo repository.RepositoryInterface, ptRepo PartnerTypeRepository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository:            repo,
		partnerTypeRepository: ptRepo,
	}
}
