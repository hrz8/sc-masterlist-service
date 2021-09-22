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
		Err:    errors.New("failed to store partner"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to list partner"),
	}
	GetById = errorMap{
		Status: 404,
		Err:    errors.New("failed to get partner"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to remove partner"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update partner"),
	}
	CreateWithEmptyPartnerTypes = errorMap{
		Status: 400,
		Err:    errors.New("cannot assign a partner with empty partner types"),
	}
	AddPartnerType = errorMap{
		Status: 400,
		Err:    errors.New("failed to add partner type into partner"),
	}
	DeletePartnerType = errorMap{
		Status: 400,
		Err:    errors.New("failed to delete partner type from partner"),
	}
	DeleteByIdHasMaterial = errorMap{
		Status: 400,
		Err:    errors.New("there is still material associate with this partner"),
	}
)
