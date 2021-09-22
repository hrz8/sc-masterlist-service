package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Part represents Part object for DB
type (
	Part struct {
		ID               uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Number           string    `gorm:"column:number;index:idx_number;unique;not null" json:"number" validate:"required"`
		Name             string    `gorm:"column:name;not null" json:"name" validate:"required"`
		Image            string    `gorm:"column:image" json:"image"`
		QtyPerUnit       uint      `gorm:"column:qty_per_unit" json:"qtyPerUnit"`
		QtyPerMonth      uint      `gorm:"column:qty_per_month" json:"qtyPerMonth"`
		DwgWeight        uint      `gorm:"column:dwg_weight" json:"dwgWeight"`
		ActualWeightPart uint      `gorm:"column:actual_weight_part" json:"actualWeightPart"`
		ActualWeightRun  uint      `gorm:"column:actual_weight_run" json:"actualWeightRun"`
		PaintColor       string    `gorm:"column:paint_color" json:"paint_color"`
		PaintCode        string    `gorm:"column:paint_code" json:"paint_code"`
		Remarks          string    `gorm:"column:remarks" json:"remarks"`
		SourcingRemarks  string    `gorm:"column:sourcing_remarks" json:"sourcing_remarks"`
		ProcessRouting   string    `gorm:"column:process_routing" json:"process_routing"`
		// has one
		ProjectID   uuid.UUID `gorm:"size:40"`
		Project     Project   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"project"`
		MaterialID  uuid.UUID `gorm:"size:40"`
		Material    Material  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"material"`
		GrainTypeID uuid.UUID `gorm:"size:40"`
		GrainType   GrainType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"grainType"`
		MouldTonID  uuid.UUID `gorm:"size:40"`
		MouldTon    MouldTon  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"mouldTon"`
		MouldCavID  uuid.UUID `gorm:"size:40"`
		MouldCav    GrainType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"mouldCav"`
		// has many
		Processes []Process `gorm:"many2many:parts_processes" json:"processes,omitempty"`
		Colors    []Color   `gorm:"many2many:parts_colors" json:"colors,omitempty"`
		// has many back-to-back
		Sourcings   []*Partner `gorm:"many2many:parts_sourcings" json:"sourcings,omitempty"`
		MouldMakers []*Partner `gorm:"many2many:parts_mould_makers" json:"mouldMakers,omitempty"`
		// timestamp
		CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// PartPayloadCreate represents payload to create part
	PartPayloadCreate struct {
		Name        string `json:"name" validate:"required,max=50"`
		Description string `json:"description" validate:"max=140"`
	}

	// PartPayloadGetAll represents payload to fetch all parts
	PartPayloadGetAll struct {
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

	// PartPayloadGet represents payload to get part by identifier
	PartPayloadGet struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// PartPayloadUpdateById represents payload to update part by identifier
	PartPayloadUpdateById struct {
		ID          uuid.UUID `json:"-" param:"id" validate:"required"`
		Name        string    `json:"name" validate:"max=50"`
		Description string    `json:"description" validate:"max=140"`
	}

	// PartPayloadDeleteById represents payload to delete part by identifier
	PartPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}
)
