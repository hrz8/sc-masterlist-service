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

	mouldMakerErrorImpl struct {
		prefix string
	}
)

func (i *mouldMakerErrorImpl) MouldMakerErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func (i *mouldMakerErrorImpl) MouldMakerErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-002", nil)
}

func (i *mouldMakerErrorImpl) MouldMakerErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-003", nil)
}

func (i *mouldMakerErrorImpl) MouldMakerErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-004", nil)
}

func NewMouldMakerError() MouldMakerErrorInterface {
	return &mouldMakerErrorImpl{
		prefix: "SCM-MOOULDMAKER",
	}
}
