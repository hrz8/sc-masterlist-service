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
		Create(trx *gorm.DB, material *models.Material) (*models.Material, error)
		GetAll(trx *gorm.DB, conditions *models.MaterialPayloadGetAll) (*[]models.Material, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.Material, error)
		DeleteById(trx *gorm.DB, id *uuid.UUID) error
		Update(
			trx *gorm.DB,
			materialInstance *models.Material,
			payload *models.MaterialPayloadUpdateById,
		) (*models.Material, error)
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
	if err := trx.Model(&models.Material{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, material *models.Material) (*models.Material, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&material).Error; err != nil {
		return nil, err
	}
	return material, nil
}

func (i *impl) GetAll(trx *gorm.DB, conditions *models.MaterialPayloadGetAll) (*[]models.Material, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Material{}
	executor := trx.
		Where("tsm LIKE ?", "%"+conditions.Tsm.Like+"%").
		Where("description LIKE ?", "%"+conditions.Description.Like+"%")

	if conditions.Deleted.Only {
		executor = executor.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if conditions.Deleted.Include {
		executor = executor.Unscoped()
	}
	if conditions.Tsm.Eq != "" {
		executor = executor.Where("tsm = ?", conditions.Tsm.Eq)
	}
	if conditions.Description.Eq != "" {
		executor = executor.Where("description = ?", conditions.Description.Eq)
	}
	if conditions.MaterialGrade.Eq != "" {
		executor = executor.Where("material_grade_id = ?", conditions.MaterialGrade.Eq)
	}
	if conditions.Maker.Eq != "" {
		executor = executor.Where("maker_id = ?", conditions.Maker.Eq)
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

	if err := executor.Debug().Preload(clause.Associations).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.Material, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Material{}
	if err := trx.Debug().
		Preload("MaterialGrade").
		Preload("Maker").
		First(&result, id).Error; err != nil {
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
	result := models.Material{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(
	trx *gorm.DB,
	materialInstance *models.Material,
	payload *models.MaterialPayloadUpdateById,
) (*models.Material, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(materialInstance).Updates(models.Material{
		Tsm:         (*payload).Tsm,
		Description: (*payload).Description,
	}).Error; err != nil {
		return nil, err
	}
	return materialInstance, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Material{})
	return &impl{
		db: db,
	}
}
