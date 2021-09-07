package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	MouldMakerErrorInterface interface {
		MouldMakerErrorCreate(*utils.CustomContext, *string, *uint16) error
		MouldMakerErrorGetAll(*utils.CustomContext, *string, *uint16) error
		MouldMakerErrorGet(*utils.CustomContext, *string, *uint16) error
		MouldMakerErrorDelete(*utils.CustomContext, *string, *uint16) error
	}

	sourcingErrorImpl struct {
		prefix string
	}
)

func (i *sourcingErrorImpl) MouldMakerErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func (i *sourcingErrorImpl) MouldMakerErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-002", nil)
}

func (i *sourcingErrorImpl) MouldMakerErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-003", nil)
}

func (i *sourcingErrorImpl) MouldMakerErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-004", nil)
}

func NewMouldMakerError() MouldMakerErrorInterface {
	return &sourcingErrorImpl{
		prefix: "SCM-MOOULDMAKER",
	}
}
