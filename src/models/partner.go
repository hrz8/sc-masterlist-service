package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	// Partner represents Partner object for DB
	Partner struct {
		ID              uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
		Name            string         `gorm:"column:name;index:idx_name;unique;not null" json:"name" validate:"required"`
		Address         string         `gorm:"column:address" json:"address"`
		Contact         string         `gorm:"column:contact" json:"contact"`
		Description     string         `gorm:"column:description" json:"description"`
		PartnerTypes    []*PartnerType `gorm:"many2many:partners_partner_types" json:"partnerTypes,omitempty"`
		SourcedParts    []*Part        `gorm:"many2many:parts_sourcings;foreignKey:ID;joinForeignKey:SourcingID;References:ID;JoinReferences:PartID" json:"sourcedParts,omitempty"`
		MouldMakedParts []*Part        `gorm:"many2many:parts_mould_makers;foreignKey:ID;joinForeignKey:MouldMakerID;References:ID;JoinReferences:PartID" json:"mouldMakedParts,omitempty"`
		// additional assoc
		Materials []Material `gorm:"foreignKey:MakerID" json:"materials,omitempty"`
		// timestamp
		CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	// PartnerPayloadCreate represents payload to create partner
	PartnerPayloadCreate struct {
		Name         string      `json:"name" validate:"required,max=50"`
		Address      string      `json:"address" validate:"max=140"`
		Contact      string      `json:"contact" validate:"max=140"`
		Description  string      `json:"description" validate:"max=140"`
		PartnerTypes []uuid.UUID `json:"partnerTypes" validate:"required"`
	}

	// PartnerPayloadGetAll represents payload to fetch all partners
	PartnerPayloadGetAll struct {
		// column
		Name        FilteringQueryParams `query:"name"`
		Address     FilteringQueryParams `query:"address"`
		Contact     FilteringQueryParams `query:"contact"`
		Description FilteringQueryParams `query:"description"`
		// relation
		PartnerTypes FilteringQueryParams `query:"partnerTypes"`
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
		ID           uuid.UUID   `json:"-" param:"id" validate:"required"`
		Name         string      `json:"name" validate:"max=50"`
		Address      string      `json:"address" validate:"max=140"`
		Contact      string      `json:"contact" validate:"max=140"`
		Description  string      `json:"description" validate:"max=140"`
		PartnerTypes []uuid.UUID `json:"partnerTypes"`
	}

	// PartnerPayloadDeleteById represents payload to delete partner by identifier
	PartnerPayloadDeleteById struct {
		ID uuid.UUID `param:"id" validate:"required"`
	}

	// PartnerAddPartnerTypePayload represents payload for adding type to particular partner
	PartnerAddPartnerTypePayload struct {
		ID            uuid.UUID `param:"id" validate:"required"`
		PartnerTypeID uuid.UUID `param:"partnerTypeId" validate:"required"`
	}

	// PartnerDeletePartnerTypePayload represents payload for adding type to particular partner
	PartnerDeletePartnerTypePayload struct {
		ID            uuid.UUID `param:"id" validate:"required"`
		PartnerTypeID uuid.UUID `param:"partnerTypeId" validate:"required"`
	}

	// PartnersPartnerTypes represents join table schema for partner -> partner_type
	PartnersPartnerTypes struct {
		PartnerID     uuid.UUID
		PartnerTypeID uuid.UUID
		CreatedAt     time.Time
		DeletedAt     gorm.DeletedAt `gorm:"index"`
	}
)

func (partnerPartnerType *PartnersPartnerTypes) BeforeCreate(tx *gorm.DB) error {
	partnerPartnerType.CreatedAt = time.Now()
	return nil
}
