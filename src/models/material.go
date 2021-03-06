package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Material represents Material object for DB
type (
	Material struct {
		ID              uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Tsm             string         `gorm:"column:tsm;index:idx_tsm;unique;not null" json:"tsm" validate:"required"`
		Description     string         `gorm:"column:description" json:"description"`
		MaterialGradeID uuid.UUID      `gorm:"size:40;not null" json:"-"`
		MaterialGrade   MaterialGrade  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"materialGrade"`
		MakerID         uuid.UUID      `gorm:"size:40;not null" json:"-"`
		Maker           *Partner       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"maker,omitempty"`
		CreatedAt       time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// MaterialPayloadCreate represents payload to create material
	MaterialPayloadCreate struct {
		Tsm           string    `json:"tsm" validate:"required,max=50"`
		Description   string    `json:"description" validate:"max=140"`
		MaterialGrade uuid.UUID `json:"materialGrade" validate:"required"`
		Maker         uuid.UUID `json:"maker" validate:"required"`
	}

	// MaterialPayloadGetAll represents payload to fetch all materials
	MaterialPayloadGetAll struct {
		// column
		Tsm         FilteringQueryParams `query:"tsm"`
		Description FilteringQueryParams `query:"description"`
		// relation
		MaterialGrade FilteringQueryParams `query:"materialGrade"`
		Maker         FilteringQueryParams `query:"maker"`
		// date props
		CreatedAt FilteringQueryParams `query:"createdAt"`
		UpdatedAt FilteringQueryParams `query:"updatedAt"`
		// built-in
		Pagination PagingQueryParams `query:"_pagination"`
		Sort       SortQueryParams   `query:"_sort"`
		Deleted    DeleteQueryParams `query:"_deleted"`
	}

	// MaterialPayloadGet represents payload to get material by identifier
	MaterialPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// MaterialPayloadUpdateById represents payload to update material by identifier
	MaterialPayloadUpdateById struct {
		ID            uuid.UUID `json:"-" param:"id" validate:"required"`
		Tsm           string    `json:"tsm" validate:"max=50"`
		Description   string    `json:"description" validate:"max=140"`
		MaterialGrade uuid.UUID `json:"materialGrade"`
		Maker         uuid.UUID `json:"maker"`
	}

	// MaterialPayloadDeleteById represents payload to delete material by identifier
	MaterialPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
