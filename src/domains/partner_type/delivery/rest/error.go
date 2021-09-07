package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	PartnerTypeErrorInterface interface {
		PartnerTypeErrorCreate(*utils.CustomContext, *string, *uint16) error
		PartnerTypeErrorGetAll(*utils.CustomContext, *string, *uint16) error
		PartnerTypeErrorGet(*utils.CustomContext, *string, *uint16) error
		PartnerTypeErrorDelete(*utils.CustomContext, *string, *uint16) error
	}

	partnerTypeErrorImpl struct {
		prefix string
	}
)

func (i *partnerTypeErrorImpl) PartnerTypeErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func (i *partnerTypeErrorImpl) PartnerTypeErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-002", nil)
}

func (i *partnerTypeErrorImpl) PartnerTypeErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-003", nil)
}

func (i *partnerTypeErrorImpl) PartnerTypeErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-004", nil)
}

func NewPartnerTypeError() PartnerTypeErrorInterface {
	return &partnerTypeErrorImpl{
		prefix: "SCM-PARTNERTYPE",
	}
}
