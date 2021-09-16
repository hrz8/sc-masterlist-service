package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// MouldTon represents MouldTon object for DB
type (
	MouldTon struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Value       string         `gorm:"column:value;index:idx_value;unique;not null" json:"value" validate:"required"`
		Description string         `gorm:"column:description" json:"description"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// MouldTonPayloadCreate represents payload to create mould ton
	MouldTonPayloadCreate struct {
		Value       string `json:"value" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// MouldTonPayloadGetAll represents payload to fetch all mould tons
	MouldTonPayloadGetAll struct {
		// column
		Value       FilteringQueryParams `query:"value"`
		Description FilteringQueryParams `query:"description"`
		// date props
		CreatedAt FilteringQueryParams `query:"createdAt"`
		UpdatedAt FilteringQueryParams `query:"updatedAt"`
		// built-in
		Pagination PagingQueryParams `query:"_pagination"`
		Sort       SortQueryParams   `query:"_sort"`
		Deleted    DeleteQueryParams `query:"_deleted"`
	}

	// MouldTonPayloadGet represents payload to get mould ton by identifier
	MouldTonPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// MouldTonPayloadUpdateById represents payload to update mould ton by identifier
	MouldTonPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Value       string    `json:"value" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// MouldTonPayloadDeleteById represents payload to delete mould ton by identifier
	MouldTonPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
