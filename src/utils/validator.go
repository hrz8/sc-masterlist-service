package utils

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hrz8/sc-masterlist-service/src/helpers"
	"github.com/hrz8/sc-masterlist-service/src/models"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
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

func strToBool(val string) bool {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		boolVal = false
	}
	return boolVal
}

// FIXME: look for better way
func QueryParamsBind(destination interface{}, c echo.Context) (err error) {
	queryParams := c.QueryParams()
	if destination == nil || len(queryParams) == 0 {
		return nil
	}
	typ := reflect.TypeOf(destination).Elem()
	val := reflect.ValueOf(destination).Elem()

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		// structField := val.Field(i)
		queryTag := typeField.Tag.Get("query")

		found := false
		switch queryTag {
		case "_pagination":
			{
				page, pageExist := queryParams["pagination[page]"]
				limit, limitExists := queryParams["pagination[limit]"]

				if !pageExist || !limitExists {
					continue
				}

				val.Field(i).Set(reflect.ValueOf(models.PagingQueryParams{
					Page:  helpers.ParseStringToInt(page[0]),
					Limit: helpers.ParseStringToInt(limit[0]),
				}))
			}
		case "_sort":
			{
				sortBy, sortByExists := queryParams["sort[by]"]
				sortMode, sortModeExists := queryParams["sort[mode]"]

				if !sortByExists || !sortModeExists {
					continue
				}

				_, modeValid := helpers.SliceStringContains([]string{"asc", "desc"}, sortMode[0])
				if !modeValid {
					continue
				}

				val.Field(i).Set(reflect.ValueOf(models.SortQueryParams{
					By:   sortBy[0],
					Mode: sortMode[0],
				}))
			}
		case "_deleted":
			{
				deletedInclude, deletedIncludeExists := queryParams["deleted[include]"]
				deletedOnly, deletedOnlyExists := queryParams["deleted[only]"]

				include := false
				only := false

				if deletedIncludeExists {
					include = strToBool(deletedInclude[0])
				}

				if deletedOnlyExists {
					only = strToBool(deletedOnly[0])
				}

				val.Field(i).Set(reflect.ValueOf(models.DeleteQueryParams{
					Include: include,
					Only:    only,
				}))
			}
		default:
			{
				field := models.FilteringQueryParams{}
				for key, v := range queryParams {
					preservedKeys := strings.HasPrefix(key, "pagination") ||
						strings.HasPrefix(key, "sort")
					val := v[0]
					if !preservedKeys {
						switch key {
						case queryTag + "[eq]":
							{
								field.Eq = val
								break
							}
						case queryTag + "[like]":
							{
								field.Like = val
								break
							}
						case queryTag + "[gte]":
							{
								field.Gte = val
								break
							}
						case queryTag + "[lte]":
							{
								field.Lte = val
								break
							}
						default:
							found = false
						}
					}
				}
				val.Field(i).Set(reflect.ValueOf(field))
			}
		}

		if !found {
			continue
		}
	}
	return nil
}

func BinderError(c *CustomContext) error {
	return c.ErrorResponse(
		nil,
		"Internal Server Error",
		http.StatusInternalServerError,
		"SCM-VALIDATOR-001",
		nil,
	)
}

func ValidatorMiddleware(models reflect.Type, queryParamsBinder bool) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*CustomContext)
			payload := reflect.New(models).Interface()
			if queryParamsBinder {
				if err := QueryParamsBind(payload, ctx); err != nil {
					return BinderError(ctx)
				}
			} else {
				if err := ctx.Bind(payload); err != nil {
					if strings.Contains(err.Error(), "uuid") {
						return ctx.ErrorResponse(
							nil,
							"invalid id",
							http.StatusNotFound,
							"SCM-VALIDATOR-002",
							nil,
						)
					}
					return BinderError(ctx)
				}
			}
			if err := ctx.Validate(payload); err != nil {
				return ctx.ErrorResponse(
					nil,
					err.Error(),
					http.StatusBadRequest,
					"SCM-VALIDATOR-002",
					nil,
				)
			}
			ctx.Payload = payload
			return next(ctx)
		}
	}
}
