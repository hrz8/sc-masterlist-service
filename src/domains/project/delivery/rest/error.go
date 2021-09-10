package rest

import (
	"errors"
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	ProjectErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	projectErrorImpl struct {
		prefix string
	}

	projectErrorMap struct {
		status int
		err    error
	}
)

var (
	ProjectErrorCreate = projectErrorMap{
		status: 400,
		err:    errors.New("failed to store project"),
	}
	ProjectErrorGetAll = projectErrorMap{
		status: 400,
		err:    errors.New("failed to list project"),
	}
	ProjectErrorGetById = projectErrorMap{
		status: 400,
		err:    errors.New("failed to get project"),
	}
	ProjectErrorDeleteById = projectErrorMap{
		status: 400,
		err:    errors.New("failed to remove project"),
	}
	ProjectErrorUpdateById = projectErrorMap{
		status: 400,
		err:    errors.New("failed to update project"),
	}
)

func (i *projectErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, ProjectErrorCreate.err) {
		status := uint16(ProjectErrorCreate.status)
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
	if errors.Is(domainErr, ProjectErrorGetAll.err) {
		status := uint16(ProjectErrorGetAll.status)
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
	if errors.Is(domainErr, ProjectErrorGetById.err) {
		errStatus := uint16(ProjectErrorGetById.status)
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
	if errors.Is(domainErr, ProjectErrorDeleteById.err) {
		errStatus := uint16(ProjectErrorDeleteById.status)
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
	if errors.Is(domainErr, ProjectErrorUpdateById.err) {
		errStatus := uint16(ProjectErrorUpdateById.status)
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

func NewProjectError() ProjectErrorInterface {
	return &projectErrorImpl{
		prefix: "SCM-PROJECT",
	}
}
