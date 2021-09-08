package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	PartnerErrorInterface interface {
		PartnerErrorCreate(*utils.CustomContext, *string, *uint16) error
	}

	partnerErrorImpl struct {
		prefix string
	}
)

func (i *partnerErrorImpl) PartnerErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func NewPartnerError() PartnerErrorInterface {
	return &partnerErrorImpl{
		prefix: "SCM-PARTNER",
	}
}
