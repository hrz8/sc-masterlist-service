package repository

import (
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(*models.Process) (*models.Process, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(process *models.Process) (*models.Process, error) {
	if err := i.db.Create(&process).Error; err != nil {
		return nil, err
	}
	return process, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Process{})
	return &impl{
		db: db,
	}
}
