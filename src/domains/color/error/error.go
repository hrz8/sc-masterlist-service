package error

import "errors"

type (
	errorMap struct {
		Status int
		Err    error
	}
)

var (
	Create = errorMap{
		Status: 400,
		Err:    errors.New("failed to store color"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list color"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get color"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove color"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update color"),
	}
)
