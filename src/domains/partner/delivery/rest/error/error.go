package rest

import (
	"errors"
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	RestErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	restErrorImpl struct {
		prefix string
	}

	restErrorMap struct {
		Status int
		Err    error
	}
)

var (
	Create = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to store partner"),
	}
	GetAll = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to list partner"),
	}
	GetById = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to get partner"),
	}
	DeleteById = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to remove partner"),
	}
	UpdateById = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to update partner"),
	}
)

func (i *restErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, Create.Err) {
		status := uint16(Create.Status)
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
	if errors.Is(domainErr, GetAll.Err) {
		status := uint16(GetAll.Status)
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
	if errors.Is(domainErr, GetById.Err) {
		errStatus := uint16(GetById.Status)
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
	if errors.Is(domainErr, DeleteById.Err) {
		errStatus := uint16(DeleteById.Status)
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
	if errors.Is(domainErr, UpdateById.Err) {
		errStatus := uint16(UpdateById.Status)
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

func NewPartnerError() RestErrorInterface {
	return &restErrorImpl{
		prefix: "SCM-PARTNER",
	}
}
