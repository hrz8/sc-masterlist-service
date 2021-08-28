package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Process represents Process object for DB
type (
	Process struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	ProcessPayloadCreate struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	ProcessPayloadGetAll struct {
		// column
		Name        FilteringQueryParams `query:"name"`
		Description FilteringQueryParams `query:"description"`
		// built-in
		Pagination PagingQueryParams `query:"_pagination"`
		Sort       SortQueryParams   `query:"_sort"`
	}
)
