package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	RepositoryInterface interface {
		CountAll(trx *gorm.DB) (*int64, error)
		Create(trx *gorm.DB, partner *models.Partner) (*models.Partner, error)
		GetAll(trx *gorm.DB, conditions *models.PartnerPayloadGetAll) (*[]models.Partner, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.Partner, error)
		DeleteById(trx *gorm.DB, id *uuid.UUID) error
		Update(
			trx *gorm.DB,
			partnerInstance *models.Partner,
			payload *models.PartnerPayloadUpdateById,
		) (*models.Partner, error)
		AddPartnerType(
			trx *gorm.DB,
			partner *models.Partner,
			partnerTypes *models.PartnerType,
		) (*models.Partner, error)
		DeletePartnerType(
			trx *gorm.DB,
			partner *models.Partner,
			partnerTypes *models.PartnerType,
		) (*models.Partner, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) CountAll(trx *gorm.DB) (*int64, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	var total int64 = 0
	if err := trx.Model(&models.Partner{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, partner *models.Partner) (*models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&partner).Error; err != nil {
		return nil, err
	}
	return partner, nil
}

func (i *impl) GetAll(trx *gorm.DB, conditions *models.PartnerPayloadGetAll) (*[]models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Partner{}
	executor := trx.
		Where("partners.name LIKE ?", "%"+conditions.Name.Like+"%").
		Where("partners.address LIKE ?", "%"+conditions.Address.Like+"%").
		Where("partners.contact LIKE ?", "%"+conditions.Contact.Like+"%").
		Where("partners.description LIKE ?", "%"+conditions.Description.Like+"%")

	// scoping conditions
	if conditions.Deleted.Only {
		executor = executor.Unscoped().Where("partners.deleted_at IS NOT NULL")
	}
	if conditions.Deleted.Include {
		executor = executor.Unscoped()
	}

	// column conditions
	if conditions.Name.Eq != "" {
		executor = executor.Where("partners.name = ?", conditions.Name.Eq)
	}
	if conditions.Address.Eq != "" {
		executor = executor.Where("partners.address = ?", conditions.Address.Eq)
	}
	if conditions.Contact.Eq != "" {
		executor = executor.Where("partners.address = ?", conditions.Contact.Eq)
	}
	if conditions.Description.Eq != "" {
		executor = executor.Where("partners.description = ?", conditions.Description.Eq)
	}

	// assoc column condition
	if len(conditions.PartnerTypes.In) > 0 {
		executor = executor.
			Joins("JOIN partners_partner_types ON partners_partner_types.partner_id = partners.id").
			Joins("JOIN partner_types ON partners_partner_types.partner_type_id = partner_types.id").
			Where("partner_types.id IN ?", conditions.PartnerTypes.In)
	}

	// timestamp conditions
	if conditions.CreatedAt.Gte != nil && conditions.CreatedAt.Lte != nil {
		executor = executor.Where("partners.created_at BETWEEN ? AND ?", conditions.CreatedAt.Gte, conditions.CreatedAt.Lte)
	}
	if conditions.UpdatedAt.Gte != nil && conditions.UpdatedAt.Lte != nil {
		executor = executor.Where("partners.updated_at BETWEEN ? AND ?", conditions.UpdatedAt.Gte, conditions.UpdatedAt.Lte)
	}

	// sort and paging condition
	if conditions.Sort.By != "" && conditions.Sort.Mode != "" {
		executor = executor.Order(conditions.Sort.By + " " + conditions.Sort.Mode)
	}
	if conditions.Pagination.Limit != nil {
		executor = executor.Limit(conditions.Pagination.Limit.(int))
	}
	if conditions.Pagination.Limit != nil && conditions.Pagination.Page != nil {
		executor = executor.Offset(helpers.GetOffset(conditions.Pagination.Page.(int), conditions.Pagination.Limit.(int)))
	}

	// select executor
	if err := executor.Debug().Preload(clause.Associations).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Partner{}
	if err := trx.Debug().Preload("PartnerTypes").First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) DeleteById(trx *gorm.DB, id *uuid.UUID) error {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Partner{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(
	trx *gorm.DB,
	partnerInstance *models.Partner,
	payload *models.PartnerPayloadUpdateById,
) (*models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(partnerInstance).Updates(models.Partner{
		Name:        (*payload).Name,
		Address:     (*payload).Address,
		Contact:     (*payload).Contact,
		Description: (*payload).Description,
	}).Error; err != nil {
		return nil, err
	}
	return partnerInstance, nil
}

func (i *impl) AddPartnerType(
	trx *gorm.DB,
	partner *models.Partner,
	partnerTypes *models.PartnerType,
) (*models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Model(&partner).Association("PartnerTypes").
		Append(partnerTypes); err != nil {
		return nil, err
	}

	return partner, nil
}

func (i *impl) DeletePartnerType(
	trx *gorm.DB,
	partner *models.Partner,
	partnerTypes *models.PartnerType,
) (*models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Model(&partner).Association("PartnerTypes").
		Delete(partnerTypes); err != nil {
		return nil, err
	}

	return partner, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Partner{})
	db.AutoMigrate(&models.PartnersPartnerTypes{})
	db.SetupJoinTable(&models.Partner{}, "PartnerTypes", &models.PartnersPartnerTypes{})
	return &impl{
		db: db,
	}
}
