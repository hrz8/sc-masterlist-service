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
		Status int
		Err    error
	}
)

var (
	PartnerTypeErrorCreate = partnerTypeErrorMap{
		Status: 400,
		Err:    errors.New("failed to store partner type"),
	}
	PartnerTypeErrorGetAll = partnerTypeErrorMap{
		Status: 400,
		Err:    errors.New("failed to list partner type"),
	}
	PartnerTypeErrorGetById = partnerTypeErrorMap{
		Status: 400,
		Err:    errors.New("failed to get partner type"),
	}
	PartnerTypeErrorDeleteById = partnerTypeErrorMap{
		Status: 400,
		Err:    errors.New("failed to remove partner type"),
	}
	PartnerTypeErrorUpdateById = partnerTypeErrorMap{
		Status: 400,
		Err:    errors.New("failed to update partner type"),
	}
)

func (i *partnerTypeErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, PartnerTypeErrorCreate.Err) {
		status := uint16(PartnerTypeErrorCreate.Status)
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
	if errors.Is(domainErr, PartnerTypeErrorGetAll.Err) {
		status := uint16(PartnerTypeErrorGetAll.Status)
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
	if errors.Is(domainErr, PartnerTypeErrorGetById.Err) {
		errStatus := uint16(PartnerTypeErrorGetById.Status)
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
	if errors.Is(domainErr, PartnerTypeErrorDeleteById.Err) {
		errStatus := uint16(PartnerTypeErrorDeleteById.Status)
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
	if errors.Is(domainErr, PartnerTypeErrorUpdateById.Err) {
		errStatus := uint16(PartnerTypeErrorUpdateById.Status)
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
