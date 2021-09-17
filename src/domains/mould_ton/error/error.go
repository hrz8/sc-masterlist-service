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
		Err:    errors.New("failed to store mould ton"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list mould ton"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get mould ton"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove mould ton"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update mould ton"),
	}
)
