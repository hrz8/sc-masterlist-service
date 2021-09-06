package rest

import (
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	ProjectErrorInterface interface {
		ProjectErrorCreate(*utils.CustomContext, *string, *uint16) error
		ProjectErrorGetAll(*utils.CustomContext, *string, *uint16) error
		ProjectErrorGet(*utils.CustomContext, *string, *uint16) error
		ProjectErrorDelete(*utils.CustomContext, *string, *uint16) error
	}

	projectErrorImpl struct {
		prefix string
	}
)

func (i *projectErrorImpl) ProjectErrorCreate(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-001", nil)
}

func (i *projectErrorImpl) ProjectErrorGetAll(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-002", nil)
}

func (i *projectErrorImpl) ProjectErrorGet(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-003", nil)
}

func (i *projectErrorImpl) ProjectErrorDelete(ctx *utils.CustomContext, message *string, status *uint16) error {
	return ctx.ErrorResponse(nil, *message, *status, i.prefix+"-004", nil)
}

func NewProjectError() ProjectErrorInterface {
	return &projectErrorImpl{
		prefix: "SCM-PROJECT",
	}
}
