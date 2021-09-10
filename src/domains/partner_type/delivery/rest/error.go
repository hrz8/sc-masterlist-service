package rest

import (
	"errors"
	"net/http"

	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/utils"
)

type (
	PartnerTypeErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	partnerTypeErrorImpl struct {
		prefix string
	}

	partnerTypeErrorMap struct {
		status int
		err    error
	}
)

var (
	PartnerTypeErrorCreate = partnerTypeErrorMap{
		status: 400,
		err:    errors.New("failed to store partner type"),
	}
	PartnerTypeErrorGetAll = partnerTypeErrorMap{
		status: 400,
		err:    errors.New("failed to list partner type"),
	}
	PartnerTypeErrorGetById = partnerTypeErrorMap{
		status: 400,
		err:    errors.New("failed to get partner type"),
	}
	PartnerTypeErrorDeleteById = partnerTypeErrorMap{
		status: 400,
		err:    errors.New("failed to remove partner type"),
	}
	PartnerTypeErrorUpdateById = partnerTypeErrorMap{
		status: 400,
		err:    errors.New("failed to update partner type"),
	}
)

func (i *partnerTypeErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, PartnerTypeErrorCreate.err) {
		status := uint16(PartnerTypeErrorCreate.status)
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
	if errors.Is(domainErr, PartnerTypeErrorGetAll.err) {
		status := uint16(PartnerTypeErrorGetAll.status)
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
	if errors.Is(domainErr, PartnerTypeErrorGetById.err) {
		errStatus := uint16(PartnerTypeErrorGetById.status)
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
	if errors.Is(domainErr, PartnerTypeErrorDeleteById.err) {
		errStatus := uint16(PartnerTypeErrorDeleteById.status)
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
	if errors.Is(domainErr, PartnerTypeErrorUpdateById.err) {
		errStatus := uint16(PartnerTypeErrorUpdateById.status)
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

func NewPartnerTypeError() PartnerTypeErrorInterface {
	return &partnerTypeErrorImpl{
		prefix: "SCM-PARTNERTYPE",
	}
}
