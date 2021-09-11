package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Project represents Project object for DB
type (
	Project struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Name        string         `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string         `gorm:"column:description" json:"description"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// ProjectPayloadCreate represents payload to create project
	ProjectPayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// ProjectPayloadGetAll represents payload to fetch all projectes
	ProjectPayloadGetAll struct {
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

	// ProjectPayloadGet represents payload to get project by identifier
	ProjectPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// ProjectPayloadUpdateById represents payload to update project by identifier
	ProjectPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Name        string    `json:"name" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// ProjectPayloadDeleteById represents payload to delete project by identifier
	ProjectPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
