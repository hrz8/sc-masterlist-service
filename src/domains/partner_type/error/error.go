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
		Err:    errors.New("failed to store partner type"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list partner type"),
	}
	GetById = errorMap{
		Status: 400,
		Err:    errors.New("failed to get partner type"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove partner type"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update partner type"),
	}
	DeleteByIdHasPartner = errorMap{
		Status: 400,
		Err:    errors.New("there is still partner associate with this partner type"),
	}
)
