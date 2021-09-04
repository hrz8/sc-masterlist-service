package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(*models.Process) (*models.Process, error)
		GetAll(*models.ProcessPayloadGetAll) (*[]models.Process, error)
		GetById(*uuid.UUID) (*models.Process, error)
		DeleteById(*uuid.UUID) error
		Update(*models.Process, *models.ProcessPayloadUpdateById) (*models.Process, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(p *models.Process) (*models.Process, error) {
	if err := i.db.Debug().Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (i *impl) GetAll(c *models.ProcessPayloadGetAll) (*[]models.Process, error) {
	result := []models.Process{}
	executor := i.db.
		Where("name LIKE ?", "%"+c.Name.Like+"%").
		Where("description LIKE ?", "%"+c.Description.Like+"%")

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
		executor = executor.Offset(helpers.GetOffset(c.Pagination.Limit.(int), c.Pagination.Limit.(int)))
	}

	if err := executor.Debug().Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) GetById(id *uuid.UUID) (*models.Process, error) {
	result := models.Process{}
	if err := i.db.Debug().First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) DeleteById(id *uuid.UUID) error {
	result := models.Process{}
	if err := i.db.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(ip *models.Process, p *models.ProcessPayloadUpdateById) (*models.Process, error) {
	if err := i.db.Debug().Model(ip).Updates(models.Process{
		Name:        (*p).Name,
		Description: (*p).Description,
	}).Error; err != nil {
		return nil, err
	}
	return ip, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Process{})
	return &impl{
		db: db,
	}
}
