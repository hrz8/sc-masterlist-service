package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		CountAll(trx *gorm.DB) (*int64, error)
		Create(trx *gorm.DB, part *models.Part) (*models.Part, error)
		GetById(trx *gorm.DB, id *uuid.UUID) (*models.Part, error)
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
	if err := trx.Model(&models.Part{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, part *models.Part) (*models.Part, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&part).Error; err != nil {
		return nil, err
	}
	return part, nil
}

func (i *impl) GetById(trx *gorm.DB, id *uuid.UUID) (*models.Part, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Part{}
	if err := trx.Debug().First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Part{})
	db.AutoMigrate(&models.PartsProcesses{})
	db.AutoMigrate(&models.PartsColors{})
	db.AutoMigrate(&models.PartsSourcings{})
	db.AutoMigrate(&models.PartsMouldMakers{})
	db.SetupJoinTable(&models.Part{}, "Process", &models.PartsProcesses{})
	db.SetupJoinTable(&models.Part{}, "Colors", &models.PartsColors{})
	db.SetupJoinTable(&models.Part{}, "Sourcings", &models.PartsSourcings{})
	db.SetupJoinTable(&models.Part{}, "MouldMaker", &models.PartsMouldMakers{})
	return &impl{
		db: db,
	}
}
