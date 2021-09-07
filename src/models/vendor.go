package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// Vendor represents Vendor object for DB
	Vendor struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Adress      string    `gorm:"column:address" json:"address"`
		Contact     string    `gorm:"column:cotact" json:"cotact"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	// VendorPayloadCreate represents payload to create sourcing
	VendorPayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Address     string `json:"address" validate:"max=140"`
		Contact     string `json:"contact" validate:"max=140"`
		Description string `json:"description" validate:"max=140"`
	}

	// VendorPayloadGetAll represents payload to fetch all sourcinges
	VendorPayloadGetAll struct {
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

	// VendorPayloadGet represents payload to get sourcing by identifier
	VendorPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// VendorPayloadUpdateById represents payload to update sourcing by identifier
	VendorPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id"`
		Name        string    `json:"name" validate:"required,max=50"`
		Address     string    `json:"address" validate:"max=140"`
		Contact     string    `json:"contact" validate:"max=140"`
		Description string    `json:"description" validate:"max=140"`
	}

	// VendorPayloadDeleteById represents payload to delete sourcing by identifier
	VendorPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
