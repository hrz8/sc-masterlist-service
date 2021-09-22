package repository

import (
	"github.com/hrz8/sc-masterlist-service/src/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface{}

	impl struct {
		db *gorm.DB
	}
)

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
