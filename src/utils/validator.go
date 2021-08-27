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
		payload := v.models
		if err := c.Bind(payload); err != nil {
			return err
		}
		if err := c.Validate(payload); err != nil {
			return err
		}
		cc := &CustomContext{
			Context: c,
			Payload: payload,
		}
		return next(cc)
	}
}

func NewValidatorMiddleware(i interface{}) ValidatorMiddlewareInterface {
	return &ValidatorMiddleware{
		models: i,
	}
}
