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
		Err:    errors.New("failed to store mould cav"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list mould cav"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get mould cav"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove mould cav"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update mould cav"),
	}
)
