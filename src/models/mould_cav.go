package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// MouldCav represents MouldCav object for DB
type (
	MouldCav struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Value       string         `gorm:"column:value;index:idx_value;unique;not null" json:"value" validate:"required"`
		Description string         `gorm:"column:description" json:"description"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// MouldCavPayloadCreate represents payload to create mould cav
	MouldCavPayloadCreate struct {
		Value       string `json:"value" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// MouldCavPayloadGetAll represents payload to fetch all mould cavs
	MouldCavPayloadGetAll struct {
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

	// MouldCavPayloadGet represents payload to get mould cav by identifier
	MouldCavPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// MouldCavPayloadUpdateById represents payload to update mould cav by identifier
	MouldCavPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Value       string    `json:"value" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// MouldCavPayloadDeleteById represents payload to delete mould cav by identifier
	MouldCavPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
