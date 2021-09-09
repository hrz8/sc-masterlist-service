package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	SourcingErrorInterface interface {
		SourcingErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error
		SourcingErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error
		SourcingErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error
		SourcingErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error
	}

	sourcingErrorImpl struct {
		prefix string
	}
)

func (i *sourcingErrorImpl) SourcingErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func (i *sourcingErrorImpl) SourcingErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-002", nil)
}

func (i *sourcingErrorImpl) SourcingErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-003", nil)
}

func (i *sourcingErrorImpl) SourcingErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-004", nil)
}

func NewSourcingError() SourcingErrorInterface {
	return &sourcingErrorImpl{
		prefix: "SCM-SOURCING",
	}
}
