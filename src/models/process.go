package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// Process represents Process object for DB
	Process struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	// ProcessPayloadCreate represents payload to create process
	ProcessPayloadCreate struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	// ProcessPayloadGetAll represents payload to fetch all processes
	ProcessPayloadGetAll struct {
		// column
		Name        FilteringQueryParams `query:"name"`
		Description FilteringQueryParams `query:"description"`
		// date props
		CreatedAt FilteringQueryParams `query:"createdAt"`
		UpdatedAt FilteringQueryParams `query:"updatedAt"`
		// built-in
		Pagination PagingQueryParams `query:"_pagination"`
		Sort       SortQueryParams   `query:"_sort"`
	}

	// ProcessPayloadGet represents payload to get process by identifier
	ProcessPayloadGet struct {
		// column
		ID string `param:"id"`
	}
)
