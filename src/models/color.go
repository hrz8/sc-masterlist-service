package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Color represents Color object for DB
type (
	Color struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Code        string         `gorm:"column:code;index:idx_code;unique;not null" json:"code" validate:"required"`
		Name        string         `gorm:"column:name;not null" json:"name" validate:"required"`
		Sfx         string         `gorm:"column:sfx" json:"sfx"`
		Description string         `gorm:"column:description" json:"description"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// ColorPayloadCreate represents payload to create color
	ColorPayloadCreate struct {
		Code        string `json:"code" validate:"required,max=50"`
		Name        string `json:"name" validate:"required,max=50"`
		Sfx         string `json:"sfx" validate:"max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// ColorPayloadGetAll represents payload to fetch all colors
	ColorPayloadGetAll struct {
		// column
		Code        FilteringQueryParams `query:"code"`
		Name        FilteringQueryParams `query:"name"`
		Sfx         FilteringQueryParams `query:"sfx"`
		Description FilteringQueryParams `query:"description"`
		// date props
		CreatedAt FilteringQueryParams `query:"createdAt"`
		UpdatedAt FilteringQueryParams `query:"updatedAt"`
		// built-in
		Pagination PagingQueryParams `query:"_pagination"`
		Sort       SortQueryParams   `query:"_sort"`
		Deleted    DeleteQueryParams `query:"_deleted"`
	}

	// ColorPayloadGet represents payload to get color by identifier
	ColorPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// ColorPayloadUpdateById represents payload to update color by identifier
	ColorPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Code        string    `json:"code" validate:"max=50"`
		Name        string    `json:"name" validate:"max=50"`
		Sfx         string    `json:"sfx" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// ColorPayloadDeleteById represents payload to delete color by identifier
	ColorPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
