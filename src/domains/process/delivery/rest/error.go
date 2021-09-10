package rest

import (
	"errors"
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	ProcessErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	processErrorImpl struct {
		prefix string
	}

	processErrorMap struct {
		Status int
		Err    error
	}
)

var (
	ProcessErrorCreate = processErrorMap{
		Status: 400,
		Err:    errors.New("failed to store process"),
	}
	ProcessErrorGetAll = processErrorMap{
		Status: 400,
		Err:    errors.New("failed to list process"),
	}
	ProcessErrorGetById = processErrorMap{
		Status: 400,
		Err:    errors.New("failed to get process"),
	}
	ProcessErrorDeleteById = processErrorMap{
		Status: 400,
		Err:    errors.New("failed to remove process"),
	}
	ProcessErrorUpdateById = processErrorMap{
		Status: 400,
		Err:    errors.New("failed to update process"),
	}
)

func (i *processErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, ProcessErrorCreate.Err) {
		status := uint16(ProcessErrorCreate.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-001",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorGetAll.Err) {
		status := uint16(ProcessErrorGetAll.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-002",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorGetById.Err) {
		errStatus := uint16(ProcessErrorGetById.Status)
		status := helpers.ParseStatusResponse(dataErr, errStatus)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-003",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorDeleteById.Err) {
		errStatus := uint16(ProcessErrorDeleteById.Status)
		status := helpers.ParseStatusResponse(dataErr, errStatus)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-004",
			nil,
		)
	}
	if errors.Is(domainErr, ProcessErrorUpdateById.Err) {
		errStatus := uint16(ProcessErrorUpdateById.Status)
		status := helpers.ParseStatusResponse(dataErr, errStatus)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
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
