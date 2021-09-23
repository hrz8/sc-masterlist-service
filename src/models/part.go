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
		ParentID    uuid.UUID `gorm:"size:40"`
		Parent      *Part     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
		ProjectID   uuid.UUID `gorm:"size:40"`
		Project     Project   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"project"`
		MaterialID  uuid.UUID `gorm:"size:40"`
		Material    Material  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"material"`
		GrainTypeID uuid.UUID `gorm:"size:40"`
		GrainType   GrainType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"grainType"`
		MouldTonID  uuid.UUID `gorm:"size:40"`
		MouldTon    MouldTon  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"mouldTon"`
		MouldCavID  uuid.UUID `gorm:"size:40"`
		MouldCav    MouldCav  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"mouldCav"`
		// has many
		Processes []Process `gorm:"many2many:parts_processes" json:"processes,omitempty"`
		Colors    []Color   `gorm:"many2many:parts_colors" json:"colors,omitempty"`
		// has many back-to-back
		Childs      []*Part    `gorm:"many2many:parts_childs;foreignKey:ID;joinForeignKey:ParentID;References:ID;JoinReferences:ChildID" json:"childs,omitempty"`
		Sourcings   []*Partner `gorm:"many2many:parts_sourcings;foreignKey:ID;joinForeignKey:PartID;References:ID;JoinReferences:SourcingID" json:"sourcings,omitempty"`
		MouldMakers []*Partner `gorm:"many2many:parts_mould_makers;foreignKey:ID;joinForeignKey:PartID;References:ID;JoinReferences:MouldMakerID" json:"mouldMakers,omitempty"`
		// timestamp
		CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// PartPayloadCreate represents payload to create part
	PartPayloadCreate struct {
		Number           string `json:"number" validate:"required,max=100"`
		Name             string `json:"name" validate:"required,max=50"`
		Image            string `json:"image"`
		QtyPerUnit       uint   `json:"qtyPerUnit"`
		QtyPerMonth      uint   `json:"qtyPerMonth"`
		DwgWeight        uint   `json:"dwgWeight"`
		ActualWeightPart uint   `json:"actualWeightPart"`
		ActualWeightRun  uint   `json:"actualWeightRun"`
		PaintColor       string `json:"paint_color"`
		PaintCode        string `json:"paint_code"`
		Remarks          string `json:"remarks"`
		SourcingRemarks  string `json:"sourcing_remarks"`
		ProcessRouting   string `json:"process_routing"`
		// 1to1 relation
		Parent    uuid.UUID `json:"parent"`
		Project   uuid.UUID `json:"project" validate:"required"`
		Material  uuid.UUID `json:"material"`
		GrainType uuid.UUID `json:"grainType"`
		MouldTon  uuid.UUID `json:"mouldTon"`
		MouldCav  uuid.UUID `json:"mouldCav"`
		// many2many relation
		Processes   []uuid.UUID `json:"processes" validate:"required"`
		Colors      []uuid.UUID `json:"colors" validate:"required"`
		Childs      []uuid.UUID `json:"partnerTypes" validate:"required"`
		Sourcings   []uuid.UUID `json:"sourcings" validate:"required"`
		MouldMakers []uuid.UUID `json:"mouldMakers" validate:"required"`
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

	// PartsProcesses represents join table schema for part -> process
	PartsProcesses struct {
		PartID    uuid.UUID
		ProcessID uuid.UUID
		CreatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	// PartsColors represents join table schema for part -> color
	PartsColors struct {
		PartID    uuid.UUID
		ColorID   uuid.UUID
		CreatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	// PartsChilds represents join table schema for part -> parts
	PartsChilds struct {
		ParentID  uuid.UUID
		ChildID   uuid.UUID
		CreatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	// PartsSourcings represents join table schema for part -> partner -> partner_type (sourcing)
	PartsSourcings struct {
		PartID     uuid.UUID
		SourcingID uuid.UUID
		CreatedAt  time.Time
		DeletedAt  gorm.DeletedAt `gorm:"index"`
	}

	// PartsMouldMakers represents join table schema for part -> partner -> partner_type (mould_maker)
	PartsMouldMakers struct {
		PartID       uuid.UUID
		MouldMakerID uuid.UUID
		CreatedAt    time.Time
		DeletedAt    gorm.DeletedAt `gorm:"index"`
	}
)

func (partProcess *PartsProcesses) BeforeCreate(tx *gorm.DB) error {
	partProcess.CreatedAt = time.Now()
	return nil
}

func (partColor *PartsColors) BeforeCreate(tx *gorm.DB) error {
	partColor.CreatedAt = time.Now()
	return nil
}

func (partSourcing *PartsChilds) BeforeCreate(tx *gorm.DB) error {
	partSourcing.CreatedAt = time.Now()
	return nil
}

func (partSourcing *PartsSourcings) BeforeCreate(tx *gorm.DB) error {
	partSourcing.CreatedAt = time.Now()
	return nil
}

func (partMouldMaker *PartsMouldMakers) BeforeCreate(tx *gorm.DB) error {
	partMouldMaker.CreatedAt = time.Now()
	return nil
}
