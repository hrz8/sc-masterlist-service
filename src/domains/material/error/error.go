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
		Err:    errors.New("failed to store material"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list material"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get material"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove material"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update material"),
	}
)
