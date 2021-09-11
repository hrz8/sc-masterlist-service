package repository

import (
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(trx *gorm.DB, partner *models.Partner) (*models.Partner, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(trx *gorm.DB, partner *models.Partner) (*models.Partner, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&partner).Error; err != nil {
		return nil, err
	}
	return partner, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Partner{})
	db.AutoMigrate(&models.PartnersPartnerTypes{})
	db.SetupJoinTable(&models.Partner{}, "PartnerTypes", &models.PartnersPartnerTypes{})
	return &impl{
		db: db,
	}
}
