package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		CountAll(trx *gorm.DB) (*int64, error)
		Create(trx *gorm.DB, mouldTon *models.MouldTon) (*models.MouldTon, error)
		GetAll(trx *gorm.DB, conditions *models.MouldTonPayloadGetAll) (*[]models.MouldTon, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.MouldTon, error)
		DeleteById(trx *gorm.DB, id *uuid.UUID) error
		Update(
			trx *gorm.DB,
			mouldTonInstance *models.MouldTon,
			payload *models.MouldTonPayloadUpdateById,
		) (*models.MouldTon, error)
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
	if err := trx.Model(&models.MouldTon{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, mouldTon *models.MouldTon) (*models.MouldTon, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&mouldTon).Error; err != nil {
		return nil, err
	}
	return mouldTon, nil
}

func (i *impl) GetAll(trx *gorm.DB, conditions *models.MouldTonPayloadGetAll) (*[]models.MouldTon, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.MouldTon{}
	executor := trx.
		Where("value LIKE ?", "%"+conditions.Value.Like+"%").
		Where("description LIKE ?", "%"+conditions.Description.Like+"%")

	if conditions.Deleted.Only {
		executor = executor.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if conditions.Deleted.Include {
		executor = executor.Unscoped()
	}
	if conditions.Value.Eq != "" {
		executor = executor.Where("value = ?", conditions.Value.Eq)
	}
	if conditions.Description.Eq != "" {
		executor = executor.Where("description = ?", conditions.Description.Eq)
	}
	if conditions.CreatedAt.Gte != nil && conditions.CreatedAt.Lte != nil {
		executor = executor.Where("created_at BETWEEN ? AND ?", conditions.CreatedAt.Gte, conditions.CreatedAt.Lte)
	}
	if conditions.UpdatedAt.Gte != nil && conditions.UpdatedAt.Lte != nil {
		executor = executor.Where("updated_at BETWEEN ? AND ?", conditions.UpdatedAt.Gte, conditions.UpdatedAt.Lte)
	}
	if conditions.Sort.By != "" && conditions.Sort.Mode != "" {
		executor = executor.Order(conditions.Sort.By + " " + conditions.Sort.Mode)
	}
	if conditions.Pagination.Limit != nil {
		executor = executor.Limit(conditions.Pagination.Limit.(int))
	}
	if conditions.Pagination.Limit != nil && conditions.Pagination.Page != nil {
		executor = executor.Offset(helpers.GetOffset(conditions.Pagination.Page.(int), conditions.Pagination.Limit.(int)))
	}

	if err := executor.Debug().Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.MouldTon, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.MouldTon{}
	if err := trx.Debug().First(&result, id).Error; err != nil {
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
	result := models.MouldTon{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(
	trx *gorm.DB,
	mouldTonInstance *models.MouldTon,
	payload *models.MouldTonPayloadUpdateById,
) (*models.MouldTon, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(mouldTonInstance).Updates(models.MouldTon{
		Value:       (*payload).Value,
		Description: (*payload).Description,
	}).Error; err != nil {
		return nil, err
	}
	return mouldTonInstance, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.MouldTon{})
	return &impl{
		db: db,
	}
}
