package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(*models.MouldMaker) (*models.MouldMaker, error)
		GetAll(*models.MouldMakerPayloadGetAll) (*[]models.MouldMaker, *int64, error)
		GetById(*uuid.UUID) (*models.MouldMaker, error)
		DeleteById(*uuid.UUID) error
		Update(*models.MouldMaker, *models.MouldMakerPayloadUpdateById) (*models.MouldMaker, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(p *models.MouldMaker) (*models.MouldMaker, error) {
	if err := i.db.Debug().Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (i *impl) GetAll(c *models.MouldMakerPayloadGetAll) (*[]models.MouldMaker, *int64, error) {
	result := []models.MouldMaker{}
	executor := i.db.
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
	var total int64
	if err := i.db.Model(&models.MouldMaker{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	return &result, &total, nil
}

func (i *impl) GetById(id *uuid.UUID) (*models.MouldMaker, error) {
	result := models.MouldMaker{}
	if err := i.db.Debug().First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) DeleteById(id *uuid.UUID) error {
	result := models.MouldMaker{}
	if err := i.db.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) Update(ip *models.MouldMaker, p *models.MouldMakerPayloadUpdateById) (*models.MouldMaker, error) {
	if err := i.db.Debug().Model(ip).Updates(models.MouldMaker{
		Name:        (*p).Name,
		Description: (*p).Description,
	}).Error; err != nil {
		return nil, err
	}
	return ip, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.MouldMaker{})
	return &impl{
		db: db,
	}
}
