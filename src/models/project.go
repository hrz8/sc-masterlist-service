package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Project represents Process object for DB
type (
	Project struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}
)
