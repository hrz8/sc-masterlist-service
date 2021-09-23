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
		Err:    errors.New("failed to store part"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list part"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get part"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove part"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update part"),
	}
)
