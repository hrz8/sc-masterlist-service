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
		Create(trx *gorm.DB, color *models.Color) (*models.Color, error)
		GetAll(trx *gorm.DB, conditions *models.ColorPayloadGetAll) (*[]models.Color, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.Color, error)
		DeleteById(trx *gorm.DB, id *uuid.UUID) error
		Update(
			trx *gorm.DB,
			colorInstance *models.Color,
			payload *models.ColorPayloadUpdateById,
		) (*models.Color, error)
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
	if err := trx.Model(&models.Color{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, color *models.Color) (*models.Color, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&color).Error; err != nil {
		return nil, err
	}
	return color, nil
}

func (i *impl) GetAll(trx *gorm.DB, conditions *models.ColorPayloadGetAll) (*[]models.Color, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Color{}
	executor := trx.
		Where("code LIKE ?", "%"+conditions.Code.Like+"%").
		Where("name LIKE ?", "%"+conditions.Name.Like+"%").
		Where("sfx LIKE ?", "%"+conditions.Sfx.Like+"%").
		Where("description LIKE ?", "%"+conditions.Description.Like+"%")

	if conditions.Deleted.Only {
		executor = executor.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if conditions.Deleted.Include {
		executor = executor.Unscoped()
	}
	if conditions.Code.Eq != "" {
		executor = executor.Where("code = ?", conditions.Code.Eq)
	}
	if conditions.Name.Eq != "" {
		executor = executor.Where("name = ?", conditions.Name.Eq)
	}
	if conditions.Sfx.Eq != "" {
		executor = executor.Where("sfx = ?", conditions.Sfx.Eq)
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

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.Color, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Color{}
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
	result := models.Color{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(
	trx *gorm.DB,
	colorInstance *models.Color,
	payload *models.ColorPayloadUpdateById,
) (*models.Color, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(colorInstance).Updates(models.Color{
		Code:        (*payload).Code,
		Description: (*payload).Description,
	}).Error; err != nil {
		return nil, err
	}
	return colorInstance, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Color{})
	return &impl{
		db: db,
	}
}
