package rest

import (
	"errors"
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	PartnerErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	partnerErrorImpl struct {
		prefix string
	}

	partnerErrorMap struct {
		Status int
		Err    error
	}
)

var (
	PartnerErrorCreate = partnerErrorMap{
		Status: 400,
		Err:    errors.New("failed to store partner"),
	}
	PartnerErrorGetAll = partnerErrorMap{
		Status: 400,
		Err:    errors.New("failed to list partner"),
	}
	PartnerErrorGetById = partnerErrorMap{
		Status: 400,
		Err:    errors.New("failed to get partner"),
	}
	PartnerErrorDeleteById = partnerErrorMap{
		Status: 400,
		Err:    errors.New("failed to remove partner"),
	}
	PartnerErrorUpdateById = partnerErrorMap{
		Status: 400,
		Err:    errors.New("failed to update partner"),
	}
)

func (i *partnerErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, PartnerErrorCreate.Err) {
		status := uint16(PartnerErrorCreate.Status)
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
	if errors.Is(domainErr, PartnerErrorGetAll.Err) {
		status := uint16(PartnerErrorGetAll.Status)
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
	if errors.Is(domainErr, PartnerErrorGetById.Err) {
		errStatus := uint16(PartnerErrorGetById.Status)
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
	if errors.Is(domainErr, PartnerErrorDeleteById.Err) {
		errStatus := uint16(PartnerErrorDeleteById.Status)
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
	if errors.Is(domainErr, PartnerErrorUpdateById.Err) {
		errStatus := uint16(PartnerErrorUpdateById.Status)
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

func NewPartnerError() PartnerErrorInterface {
	return &partnerErrorImpl{
		prefix: "SCM-PARTNER",
	}
}
