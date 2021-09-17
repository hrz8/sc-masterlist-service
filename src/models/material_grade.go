package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// MaterialGrade represents MaterialGrade object for DB
type (
	MaterialGrade struct {
		ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Code        string         `gorm:"column:code;index:idx_code;unique;not null" json:"code" validate:"required"`
		Description string         `gorm:"column:description" json:"description"`
		CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// MaterialGradePayloadCreate represents payload to create material grade
	MaterialGradePayloadCreate struct {
		Code        string `json:"code" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// MaterialGradePayloadGetAll represents payload to fetch all material grades
	MaterialGradePayloadGetAll struct {
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

	// MaterialGradePayloadGet represents payload to get material grade by identifier
	MaterialGradePayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// MaterialGradePayloadUpdateById represents payload to update material grade by identifier
	MaterialGradePayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Code        string    `json:"code" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// MaterialGradePayloadDeleteById represents payload to delete material grade by identifier
	MaterialGradePayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
