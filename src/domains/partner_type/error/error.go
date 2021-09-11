package error

import (
	"errors"
)

type (
	restErrorMap struct {
		Status int
		Err    error
	}
)

var (
	Create = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to store partner type"),
	}
	GetAll = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to list partner type"),
	}
	GetById = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to get partner type"),
	}
	DeleteById = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to remove partner type"),
	}
	UpdateById = restErrorMap{
		Status: 400,
		Err:    errors.New("failed to update partner type"),
	}
)
