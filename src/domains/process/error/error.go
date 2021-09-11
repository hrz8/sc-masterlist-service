package error

import (
	"errors"
)

type (
	errorMap struct {
		Status int
		Err    error
	}
)

var (
	Create = errorMap{
		Status: 400,
		Err:    errors.New("failed to store process"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list process"),
	}
	GetById = errorMap{
		Status: 400,
		Err:    errors.New("failed to get process"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove process"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update process"),
	}
)
