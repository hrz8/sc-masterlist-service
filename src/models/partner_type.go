package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// PartnerType represents PartnerType object for DB
	PartnerType struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	// PartnerTypePayloadCreate represents payload to create partner type
	PartnerTypePayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Description string `json:"description"`
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
		Name        string    `json:"name" validate:"required,max=50"`
		Description string    `json:"description"`
	}

	// PartnerTypePayloadDeleteById represents payload to delete partner type by identifier
	PartnerTypePayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
