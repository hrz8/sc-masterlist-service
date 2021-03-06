package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// PartnerType represents PartnerType object for DB
	PartnerType struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Name        string         `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string         `gorm:"column:description" json:"description"`
		Partners    []*Partner     `gorm:"many2many:partners_partner_types" json:"partners,omitempty"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// PartnerTypePayloadCreate represents payload to create partner type
	PartnerTypePayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// PartnerTypePayloadGetAll represents payload to fetch all partner types
	PartnerTypePayloadGetAll struct {
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

	// PartnerTypePayloadGet represents payload to get partner type by identifier
	PartnerTypePayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// PartnerTypePayloadUpdateById represents payload to update partner type by identifier
	PartnerTypePayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Name        string    `json:"name" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// PartnerTypePayloadDeleteById represents payload to delete partner type by identifier
	PartnerTypePayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
