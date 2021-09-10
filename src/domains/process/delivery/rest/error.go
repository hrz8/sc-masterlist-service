package rest

import (
	"errors"
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	ProcessErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, err error) error
	}

	processErrorImpl struct {
		prefix string
	}

	processErrorMap struct {
		status int
		err    error
	}
)

var (
	ProcessErrorCreate = processErrorMap{
		status: 400,
		err:    errors.New("failed to store process"),
	}
	ProcessErrorGetAll = processErrorMap{
		status: 400,
		err:    errors.New("failed to list process"),
	}
	ProcessErrorGetById = processErrorMap{
		status: 400,
		err:    errors.New("failed to get process"),
	}
	ProcessErrorDeleteById = processErrorMap{
		status: 400,
		err:    errors.New("failed to remove process"),
	}
	ProcessErrorUpdateById = processErrorMap{
		status: 400,
		err:    errors.New("failed to update process"),
	}
)

func (i *processErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, err error) error {
	if errors.Is(domainErr, ProcessErrorCreate.err) {
		status := uint16(ProcessErrorCreate.status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": err.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-001",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorGetAll.err) {
		status := uint16(ProcessErrorGetAll.status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": err.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-002",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorGetById.err) {
		errStatus := uint16(ProcessErrorGetById.status)
		status := helpers.ParseStatusResponse(err, errStatus)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": err.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-003",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorDeleteById.err) {
		errStatus := uint16(ProcessErrorDeleteById.status)
		status := helpers.ParseStatusResponse(err, errStatus)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": err.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-004",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorUpdateById.err) {
		errStatus := uint16(ProcessErrorUpdateById.status)
		status := helpers.ParseStatusResponse(err, errStatus)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": err.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-005",
			nil,
		)
	}
	return ctx.ErrorResponse(
		nil,
		"Internal Server Error",
		http.StatusInternalServerError,
		i.prefix+"-REST-500",
		nil,
	)
}

func NewProcessError() ProcessErrorInterface {
	return &processErrorImpl{
		prefix: "SCM-PROCESS",
	}
}
