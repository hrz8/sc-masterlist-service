package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// Partner represents Partner object for DB
	Partner struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Adress      string    `gorm:"column:address" json:"address"`
		Contact     string    `gorm:"column:cotact" json:"cotact"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	// PartnerPayloadCreate represents payload to create partner
	PartnerPayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Address     string `json:"address" validate:"max=140"`
		Contact     string `json:"contact" validate:"max=140"`
		Description string `json:"description" validate:"max=140"`
	}

	// PartnerPayloadGetAll represents payload to fetch all partners
	PartnerPayloadGetAll struct {
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

	// PartnerPayloadGet represents payload to get partner by identifier
	PartnerPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// PartnerPayloadUpdateById represents payload to update partner by identifier
	PartnerPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Name        string    `json:"name" validate:"required,max=50"`
		Address     string    `json:"address" validate:"max=140"`
		Contact     string    `json:"contact" validate:"max=140"`
		Description string    `json:"description" validate:"max=140"`
	}

	// PartnerPayloadDeleteById represents payload to delete partner by identifier
	PartnerPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
