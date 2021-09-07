package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// Sourcing represents Sourcing object for DB
	Sourcing struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	// SourcingPayloadCreate represents payload to create sourcing
	SourcingPayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Description string `json:"description"`
	}

	// SourcingPayloadGetAll represents payload to fetch all sourcings
	SourcingPayloadGetAll struct {
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

	// SourcingPayloadGet represents payload to get sourcing by identifier
	SourcingPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required,uuid4"`
	}

	// SourcingPayloadUpdateById represents payload to update sourcing by identifier
	SourcingPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required,uuid4"`
		Name        string    `json:"name" validate:"required,max=50"`
		Description string    `json:"description"`
	}

	// SourcingPayloadDeleteById represents payload to delete sourcing by identifier
	SourcingPayloadDeleteById struct {
		ID uuid.UUID `param:"id" vvalidate:"required,uuid4"`
	}
)
