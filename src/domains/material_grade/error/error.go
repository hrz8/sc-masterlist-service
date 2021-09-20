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
		Err:    errors.New("failed to store material grade"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list material grade"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get material grade"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove material grade"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update material grade"),
	}
	DeleteByIdHasMaterial = errorMap{
		Status: 400,
		Err:    errors.New("there is still material associate with this material grade"),
	}
)
