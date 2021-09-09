package repository

import (
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(*gorm.DB, *models.Partner) (*models.Partner, error)
	}

	impl struct{}
)

func (i *impl) Create(trx *gorm.DB, p *models.Partner) (*models.Partner, error) {
	if err := trx.Debug().Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Partner{})
	db.AutoMigrate(&models.PartnersPartnerTypes{})
	return &impl{}
}
