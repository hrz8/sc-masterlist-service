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
		Name        string `json:"name" validate:"required,max=50"`
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
		Deleted    DeleteQueryParams `query:"_deleted"`
	}

	// ProcessPayloadGet represents payload to get process by identifier
	ProcessPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required,uuid4"`
	}

	// ProcessPayloadUpdateById represents payload to update process by identifier
	ProcessPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required,uuid4"`
		Name        string    `json:"name" validate:"required,max=50"`
		Description string    `json:"description"`
	}

	// ProcessPayloadDeleteById represents payload to delete process by identifier
	ProcessPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required,uuid4"`
	}
)
