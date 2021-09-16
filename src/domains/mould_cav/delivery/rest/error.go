package rest

import (
	"errors"
	"net/http"

	MouldCavError "github.com/hrz8/sc-masterlist-service/src/domains/mould_cav/error"
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
)

func (i *restErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, MouldCavError.Create.Err) {
		status := uint16(MouldCavError.Create.Status)
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
	if errors.Is(domainErr, MouldCavError.GetAll.Err) {
		status := uint16(MouldCavError.GetAll.Status)
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
	if errors.Is(domainErr, MouldCavError.GetById.Err) {
		errStatus := uint16(MouldCavError.GetById.Status)
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
	if errors.Is(domainErr, MouldCavError.DeleteById.Err) {
		errStatus := uint16(MouldCavError.DeleteById.Status)
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
	if errors.Is(domainErr, MouldCavError.UpdateById.Err) {
		errStatus := uint16(MouldCavError.UpdateById.Status)
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

func NewMouldCavError() RestErrorInterface {
	return &restErrorImpl{
		prefix: "SCM-MOULDCAV",
	}
}
