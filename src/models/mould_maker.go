package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// MouldMaker represents MouldMaker object for DB
	MouldMaker struct {
		ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Description string    `gorm:"column:description" json:"description"`
		gorm.Model  `json:"-"`
	}

	// MouldMakerPayloadCreate represents payload to create mould maker
	MouldMakerPayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Description string `json:"description"`
	}

	// MouldMakerPayloadGetAll represents payload to fetch all mould makers
	MouldMakerPayloadGetAll struct {
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

	// MouldMakerPayloadGet represents payload to get mould maker by identifier
	MouldMakerPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// MouldMakerPayloadUpdateById represents payload to update mould maker by identifier
	MouldMakerPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Name        string    `json:"name" validate:"required,max=50"`
		Description string    `json:"description"`
	}

	// MouldMakerPayloadDeleteById represents payload to delete mould maker by identifier
	MouldMakerPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
