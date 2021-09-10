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
		status int
		err    error
	}
)

var (
	PartnerErrorCreate = partnerErrorMap{
		status: 400,
		err:    errors.New("failed to store partner"),
	}
	PartnerErrorGetAll = partnerErrorMap{
		status: 400,
		err:    errors.New("failed to list partner"),
	}
	PartnerErrorGetById = partnerErrorMap{
		status: 400,
		err:    errors.New("failed to get partner"),
	}
	PartnerErrorDeleteById = partnerErrorMap{
		status: 400,
		err:    errors.New("failed to remove partner"),
	}
	PartnerErrorUpdateById = partnerErrorMap{
		status: 400,
		err:    errors.New("failed to update partner"),
	}
)

func (i *partnerErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, PartnerErrorCreate.err) {
		status := uint16(PartnerErrorCreate.status)
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
	if errors.Is(domainErr, PartnerErrorGetAll.err) {
		status := uint16(PartnerErrorGetAll.status)
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
	if errors.Is(domainErr, PartnerErrorGetById.err) {
		errStatus := uint16(PartnerErrorGetById.status)
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
	if errors.Is(domainErr, PartnerErrorDeleteById.err) {
		errStatus := uint16(PartnerErrorDeleteById.status)
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
	if errors.Is(domainErr, PartnerErrorUpdateById.err) {
		errStatus := uint16(PartnerErrorUpdateById.status)
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
