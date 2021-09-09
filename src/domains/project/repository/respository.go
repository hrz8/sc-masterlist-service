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
		GetAll(trx *gorm.DB, condition *models.ProjectPayloadGetAll) (*[]models.Project, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.Project, error)
		DeleteById(trx *gorm.DB, id *uuid.UUID) error
		Update(trx *gorm.DB, instanceProject *models.Project, project *models.ProjectPayloadUpdateById) (*models.Project, error)
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

func (i *impl) GetAll(trx *gorm.DB, condition *models.ProjectPayloadGetAll) (*[]models.Project, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Project{}
	executor := trx.
		Where("name LIKE ?", "%"+condition.Name.Like+"%").
		Where("description LIKE ?", "%"+condition.Description.Like+"%")

	if condition.Deleted.Only {
		executor = executor.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if condition.Deleted.Include {
		executor = executor.Unscoped()
	}
	if condition.Name.Eq != "" {
		executor = executor.Where("name = ?", condition.Name.Eq)
	}
	if condition.Description.Eq != "" {
		executor = executor.Where("description = ?", condition.Description.Eq)
	}
	if condition.CreatedAt.Gte != nil && condition.CreatedAt.Lte != nil {
		executor = executor.Where("created_at BETWEEN ? AND ?", condition.CreatedAt.Gte, condition.CreatedAt.Lte)
	}
	if condition.UpdatedAt.Gte != nil && condition.UpdatedAt.Lte != nil {
		executor = executor.Where("updated_at BETWEEN ? AND ?", condition.UpdatedAt.Gte, condition.UpdatedAt.Lte)
	}
	if condition.Sort.By != "" && condition.Sort.Mode != "" {
		executor = executor.Order(condition.Sort.By + " " + condition.Sort.Mode)
	}
	if condition.Pagination.Limit != nil {
		executor = executor.Limit(condition.Pagination.Limit.(int))
	}
	if condition.Pagination.Limit != nil && condition.Pagination.Page != nil {
		executor = executor.Offset(helpers.GetOffset(condition.Pagination.Page.(int), condition.Pagination.Limit.(int)))
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

func (i *impl) Update(trx *gorm.DB, instanceProject *models.Project, project *models.ProjectPayloadUpdateById) (*models.Project, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(instanceProject).Updates(models.Project{
		Name:        (*project).Name,
		Description: (*project).Description,
	}).Error; err != nil {
		return nil, err
	}
	return instanceProject, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Project{})
	return &impl{
		db: db,
	}
}
