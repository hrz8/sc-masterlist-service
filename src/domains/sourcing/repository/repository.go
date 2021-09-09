package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(*gorm.DB, *models.Sourcing) (*models.Sourcing, error)
		GetAll(*gorm.DB, *models.SourcingPayloadGetAll) (*[]models.Sourcing, *int64, error)
		GetById(*gorm.DB, *uuid.UUID) (*models.Sourcing, error)
		DeleteById(*gorm.DB, *uuid.UUID) error
		Update(*gorm.DB, *models.Sourcing, *models.SourcingPayloadUpdateById) (*models.Sourcing, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(trx *gorm.DB, p *models.Sourcing) (*models.Sourcing, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (i *impl) GetAll(trx *gorm.DB, c *models.SourcingPayloadGetAll) (*[]models.Sourcing, *int64, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Sourcing{}
	executor := trx.
		Where("name LIKE ?", "%"+c.Name.Like+"%").
		Where("description LIKE ?", "%"+c.Description.Like+"%")

	if c.Deleted.Only {
		executor = executor.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if c.Deleted.Include {
		executor = executor.Unscoped()
	}
	if c.Name.Eq != "" {
		executor = executor.Where("name = ?", c.Name.Eq)
	}
	if c.Description.Eq != "" {
		executor = executor.Where("description = ?", c.Description.Eq)
	}
	if c.CreatedAt.Gte != nil && c.CreatedAt.Lte != nil {
		executor = executor.Where("created_at BETWEEN ? AND ?", c.CreatedAt.Gte, c.CreatedAt.Lte)
	}
	if c.UpdatedAt.Gte != nil && c.UpdatedAt.Lte != nil {
		executor = executor.Where("updated_at BETWEEN ? AND ?", c.UpdatedAt.Gte, c.UpdatedAt.Lte)
	}
	if c.Sort.By != "" && c.Sort.Mode != "" {
		executor = executor.Order(c.Sort.By + " " + c.Sort.Mode)
	}
	if c.Pagination.Limit != nil {
		executor = executor.Limit(c.Pagination.Limit.(int))
	}
	if c.Pagination.Limit != nil && c.Pagination.Page != nil {
		executor = executor.Offset(helpers.GetOffset(c.Pagination.Page.(int), c.Pagination.Limit.(int)))
	}

	if err := executor.Debug().Find(&result).Error; err != nil {
		return nil, nil, err
	}

	// get count from all rows
	var total int64 = 0
	if err := trx.Model(&models.Sourcing{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	return &result, &total, nil
}

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.Sourcing, error) {
	result := models.Sourcing{}
	if err := i.db.Debug().First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) DeleteById(trx *gorm.DB, id *uuid.UUID) error {
	result := models.Sourcing{}
	if err := i.db.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(trx *gorm.DB, ip *models.Sourcing, p *models.SourcingPayloadUpdateById) (*models.Sourcing, error) {
	if err := i.db.Debug().Model(ip).Updates(models.Sourcing{
		Name:        (*p).Name,
		Description: (*p).Description,
	}).Error; err != nil {
		return nil, err
	}
	return ip, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Sourcing{})
	return &impl{
		db: db,
	}
}
