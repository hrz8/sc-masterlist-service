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
		Status int
		Err    error
	}
)

var (
	ProjectErrorCreate = projectErrorMap{
		Status: 400,
		Err:    errors.New("failed to store project"),
	}
	ProjectErrorGetAll = projectErrorMap{
		Status: 400,
		Err:    errors.New("failed to list project"),
	}
	ProjectErrorGetById = projectErrorMap{
		Status: 400,
		Err:    errors.New("failed to get project"),
	}
	ProjectErrorDeleteById = projectErrorMap{
		Status: 400,
		Err:    errors.New("failed to remove project"),
	}
	ProjectErrorUpdateById = projectErrorMap{
		Status: 400,
		Err:    errors.New("failed to update project"),
	}
)

func (i *projectErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, ProjectErrorCreate.Err) {
		status := uint16(ProjectErrorCreate.Status)
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
	if errors.Is(domainErr, ProjectErrorGetAll.Err) {
		status := uint16(ProjectErrorGetAll.Status)
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
	if errors.Is(domainErr, ProjectErrorGetById.Err) {
		errStatus := uint16(ProjectErrorGetById.Status)
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
	if errors.Is(domainErr, ProjectErrorDeleteById.Err) {
		errStatus := uint16(ProjectErrorDeleteById.Status)
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
	if errors.Is(domainErr, ProjectErrorUpdateById.Err) {
		errStatus := uint16(ProjectErrorUpdateById.Status)
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
