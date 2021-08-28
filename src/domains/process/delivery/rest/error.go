package rest

import (
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	ProcessErrorInterface interface {
		ProcessErrorCreate(*utils.CustomContext, string) error
	}

	processErrorImpl struct {
		prefix string
	}
)

func (i *processErrorImpl) ProcessErrorCreate(ctx *utils.CustomContext, message string) error {
	return ctx.ErrorResponse(nil, message, http.StatusBadRequest, i.prefix+"-001", nil)
}

func NewProcessError() ProcessErrorInterface {
	return &processErrorImpl{
		prefix: "SCM-PROCESS",
	}
}
