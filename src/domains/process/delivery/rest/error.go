package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	ProcessErrorInterface interface {
		ProcessErrorCreate(*utils.CustomContext, *string, *uint16) error
		ProcessErrorGetAll(*utils.CustomContext, *string, *uint16) error
		ProcessErrorGet(*utils.CustomContext, *string, *uint16) error
		ProcessErrorDelete(*utils.CustomContext, *string, *uint16) error
	}

	processErrorImpl struct {
		prefix string
	}
)

func (i *processErrorImpl) ProcessErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func (i *processErrorImpl) ProcessErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-002", nil)
}

func (i *processErrorImpl) ProcessErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-003", nil)
}

func (i *processErrorImpl) ProcessErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-004", nil)
}

func NewProcessError() ProcessErrorInterface {
	return &processErrorImpl{
		prefix: "SCM-PROCESS",
	}
}
