package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	ValidatorMiddlewareInterface interface {
		Handler(echo.HandlerFunc) echo.HandlerFunc
	}

	CustomValidator struct {
		validator *validator.Validate
	}

	ValidatorMiddleware struct {
		models interface{}
	}
)

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewValidator() echo.Validator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (v *ValidatorMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*CustomContext)
		payload := v.models
		if err := ctx.Bind(payload); err != nil {
			return ctx.ErrorResponse(nil, "Internal Server Error", http.StatusInternalServerError, "SCM-VALIDATOR-001", nil)
		}
		if err := ctx.Validate(payload); err != nil {
			return ctx.ErrorResponse(nil, err.Error(), http.StatusBadRequest, "SCM-VALIDATOR-002", nil)
		}
		ctx.Payload = payload
		return next(ctx)
	}
}

func NewValidatorMiddleware(i interface{}) ValidatorMiddlewareInterface {
	return &ValidatorMiddleware{
		models: i,
	}
}
