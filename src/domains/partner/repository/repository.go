package repository

import (
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(*models.Partner) (*models.Partner, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(p *models.Partner) (*models.Partner, error) {
	if err := i.db.Debug().Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Partner{})
	db.AutoMigrate(&models.PartnersPartnerTypes{})
	return &impl{
		db: db,
	}
}
