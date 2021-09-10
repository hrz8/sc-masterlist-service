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
		Create(trx *gorm.DB, project *models.Project) (*models.Project, error)
		GetAll(trx *gorm.DB, conditions *models.ProjectPayloadGetAll) (*[]models.Project, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.Project, error)
		DeleteById(trx *gorm.DB, id *uuid.UUID) error
		Update(
			trx *gorm.DB,
			projectInstance *models.Project,
			payload *models.ProjectPayloadUpdateById,
		) (*models.Project, error)
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
	if err := trx.Model(&models.Project{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, project *models.Project) (*models.Project, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (i *impl) GetAll(trx *gorm.DB, conditions *models.ProjectPayloadGetAll) (*[]models.Project, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Project{}
	executor := trx.
		Where("name LIKE ?", "%"+conditions.Name.Like+"%").
		Where("description LIKE ?", "%"+conditions.Description.Like+"%")

	if conditions.Deleted.Only {
		executor = executor.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if conditions.Deleted.Include {
		executor = executor.Unscoped()
	}
	if conditions.Name.Eq != "" {
		executor = executor.Where("name = ?", conditions.Name.Eq)
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

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.Project, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Project{}
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
	result := models.Project{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(
	trx *gorm.DB,
	projectInstance *models.Project,
	payload *models.ProjectPayloadUpdateById,
) (*models.Project, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(projectInstance).Updates(models.Project{
		Name:        (*payload).Name,
		Description: (*payload).Description,
	}).Error; err != nil {
		return nil, err
	}
	return projectInstance, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Project{})
	return &impl{
		db: db,
	}
}
