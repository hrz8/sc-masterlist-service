package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Project represents Process object for DB
type (
	Project struct {
		gorm.Model
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name,unique" json:"name"`
		Description string    `gorm:"column:description" json:"description"`
	}
)
