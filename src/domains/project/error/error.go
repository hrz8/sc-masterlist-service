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
		Err:    errors.New("failed to store project"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list project"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get project"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove project"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update project"),
	}
)
