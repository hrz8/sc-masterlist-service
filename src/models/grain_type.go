package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// GrainType represents GrainType object for DB
type (
	GrainType struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Code        string         `gorm:"column:code;index:idx_code;unique;not null" json:"code" validate:"required"`
		Description string         `gorm:"column:description" json:"description"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// GrainTypePayloadCreate represents payload to create grain type
	GrainTypePayloadCreate struct {
		Code        string `json:"code" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// GrainTypePayloadGetAll represents payload to fetch all grain types
	GrainTypePayloadGetAll struct {
		// column
		Code        FilteringQueryParams `query:"code"`
		Description FilteringQueryParams `query:"description"`
		// date props
		CreatedAt FilteringQueryParams `query:"createdAt"`
		UpdatedAt FilteringQueryParams `query:"updatedAt"`
		// built-in
		Pagination PagingQueryParams `query:"_pagination"`
		Sort       SortQueryParams   `query:"_sort"`
		Deleted    DeleteQueryParams `query:"_deleted"`
	}

	// GrainTypePayloadGet represents payload to get grain type by identifier
	GrainTypePayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// GrainTypePayloadUpdateById represents payload to update grain type by identifier
	GrainTypePayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Code        string    `json:"code" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// GrainTypePayloadDeleteById represents payload to delete grain type by identifier
	GrainTypePayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
