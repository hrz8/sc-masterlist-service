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
		Err:    errors.New("failed to store grain type"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list grain type"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get grain type"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove grain type"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update grain type"),
	}
)
